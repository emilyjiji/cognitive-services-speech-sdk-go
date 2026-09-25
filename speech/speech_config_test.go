// Copyright (c) Microsoft. All rights reserved.
// Licensed under the MIT license. See LICENSE.md file in the project root for full license information.

package speech

import (
	"testing"

	"github.com/Microsoft/cognitive-services-speech-sdk-go/audio"
	"github.com/Microsoft/cognitive-services-speech-sdk-go/common"
)

func TestFromSubscription(t *testing.T) {
	subscription := "test"
	region := "region"
	config, err := NewSpeechConfigFromSubscription(subscription, region)
	if err != nil {
		t.Error("Unexpected error")
	}
	if config.SubscriptionKey() != subscription {
		t.Error("Subscription not properly set")
	}
	if config.Region() != region {
		t.Error("Region not properly set")
	}
}

func TestFromAuthorizationToken(t *testing.T) {
	auth := "test"
	region := "region"
	config, err := NewSpeechConfigFromAuthorizationToken(auth, region)
	if err != nil {
		t.Error("Unexpected error")
	}
	if config.AuthorizationToken() != auth {
		t.Error("Authorization Token not properly set")
	}
	if config.Region() != region {
		t.Error("Region not properly set")
	}
}

func TestPropertiesByID(t *testing.T) {
	subscription := "test"
	region := "region"
	config, err := NewSpeechConfigFromSubscription(subscription, region)
	if err != nil {
		t.Error("Unexpected error")
	}
	value := "value1"
	err = config.SetProperty(common.SpeechServiceConnectionKey, value)
	if err != nil {
		t.Error("Unexpected error")
	}
	if config.GetProperty(common.SpeechServiceConnectionKey) != value {
		t.Error("Propery value not valid")
	}
}

func TestPropertiesByString(t *testing.T) {
	subscription := "test"
	region := "region"
	config, err := NewSpeechConfigFromSubscription(subscription, region)
	if err != nil {
		t.Error("Unexpected error")
	}
	value := "value1"
	err = config.SetPropertyByString("key1", value)
	if err != nil {
		t.Error("Unexpected error")
	}
	if config.GetPropertyByString("key1") != value {
		t.Error("Propery value not valid")
	}

}

func TestModel(t *testing.T) {
	config, err := NewSpeechConfigFromSubscription("test", "region")
	if err != nil {
		t.Fatalf("NewSpeechConfigFromSubscription returned an error: %v", err)
	}
	defer config.Close()

	if model := config.Model(); model != "" {
		t.Fatalf("Model() = %q, want empty string", model)
	}

	const model = "test-model"
	if err := config.SetModel(model); err != nil {
		t.Fatalf("SetModel(%q) returned an error: %v", model, err)
	}
	if got := config.Model(); got != model {
		t.Fatalf("Model() = %q, want %q", got, model)
	}
	if got := config.GetPropertyByString(speechModelNameProperty); got != model {
		t.Fatalf("model property = %q, want %q", got, model)
	}

	audioConfig, err := audio.NewAudioConfigFromWavFileInput("../test_files/turn_on_the_lamp.wav")
	if err != nil {
		t.Fatalf("NewAudioConfigFromWavFileInput returned an error: %v", err)
	}
	defer audioConfig.Close()

	recognizer, err := NewSpeechRecognizerFromConfig(config, audioConfig)
	if err != nil {
		t.Fatalf("NewSpeechRecognizerFromConfig returned an error: %v", err)
	}
	defer recognizer.Close()

	if got := recognizer.Properties.GetPropertyByString(speechModelNameProperty, ""); got != model {
		t.Fatalf("recognizer model property = %q, want %q", got, model)
	}

	if err := config.SetModel(""); err != nil {
		t.Fatalf("SetModel(\"\") returned an error: %v", err)
	}
	if model := config.Model(); model != "" {
		t.Fatalf("Model() after clear = %q, want empty string", model)
	}
	if model := config.properties.GetPropertyByString(speechModelNameProperty, "not-empty"); model != "" {
		t.Fatalf("raw model property after clear = %q, want empty string", model)
	}
}

func TestTranslationConfigModel(t *testing.T) {
	config, err := NewSpeechTranslationConfigFromSubscription("test", "region")
	if err != nil {
		t.Fatalf("NewSpeechTranslationConfigFromSubscription returned an error: %v", err)
	}
	defer config.Close()

	if model := config.Model(); model != "" {
		t.Fatalf("Model() = %q, want empty string", model)
	}

	const model = "test-model"
	if err := config.SetModel(model); err != nil {
		t.Fatalf("SetModel(%q) returned an error: %v", model, err)
	}
	if got := config.Model(); got != model {
		t.Fatalf("Model() = %q, want %q", got, model)
	}

	if err := config.SetSpeechRecognitionLanguage("en-US"); err != nil {
		t.Fatalf("SetSpeechRecognitionLanguage returned an error: %v", err)
	}
	if err := config.AddTargetLanguage("de"); err != nil {
		t.Fatalf("AddTargetLanguage returned an error: %v", err)
	}

	audioConfig, err := audio.NewAudioConfigFromWavFileInput("../test_files/turn_on_the_lamp.wav")
	if err != nil {
		t.Fatalf("NewAudioConfigFromWavFileInput returned an error: %v", err)
	}
	defer audioConfig.Close()

	recognizer, err := NewTranslationRecognizerFromConfig(config, audioConfig)
	if err != nil {
		t.Fatalf("NewTranslationRecognizerFromConfig returned an error: %v", err)
	}
	defer recognizer.Close()

	if got := recognizer.Properties.GetPropertyByString(speechModelNameProperty, ""); got != model {
		t.Fatalf("translation recognizer model property = %q, want %q", got, model)
	}

	if err := config.SetModel(""); err != nil {
		t.Fatalf("SetModel(\"\") returned an error: %v", err)
	}
	if model := config.Model(); model != "" {
		t.Fatalf("Model() after clear = %q, want empty string", model)
	}
	if model := config.properties.GetPropertyByString(speechModelNameProperty, "not-empty"); model != "" {
		t.Fatalf("raw model property after clear = %q, want empty string", model)
	}
}
