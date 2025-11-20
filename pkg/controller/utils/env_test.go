package utils

import (
	"testing"
)

func TestIsCreateOnlyEnvEnabled(t *testing.T) {
	tests := []struct {
		name        string
		envValue    string
		expected    bool
		description string
	}{
		{
			name:        "env var set to true",
			envValue:    "true",
			expected:    true,
			description: "should return true when CREATE_ONLY_MODE is 'true'",
		},
		{
			name:        "env var set to false",
			envValue:    "false",
			expected:    false,
			description: "should return false when CREATE_ONLY_MODE is 'false'",
		},
		{
			name:        "env var not set",
			envValue:    "",
			expected:    false,
			description: "should return false when CREATE_ONLY_MODE is not set",
		},
		{
			name:        "env var set to invalid value",
			envValue:    "invalid",
			expected:    false,
			description: "should return false when CREATE_ONLY_MODE has invalid value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				t.Setenv(createOnlyEnvName, tt.envValue)
			} else {
				t.Setenv(createOnlyEnvName, "")
			}
			result := IsCreateOnlyEnvEnabled()
			if result != tt.expected {
				t.Errorf("IsCreateOnlyEnvEnabled() = %v, expected %v - %s", result, tt.expected, tt.description)
			}
		})
	}
}

func TestIsInCreateOnlyMode(t *testing.T) {
	tests := []struct {
		name              string
		envValue          string
		createOnlyFlag    bool
		expectedResult    bool
		expectedFlagAfter bool
		description       string
	}{
		{
			name:              "flag false without env var",
			envValue:          "",
			createOnlyFlag:    false,
			expectedResult:    false,
			expectedFlagAfter: false,
			description:       "should return false when flag is false and env is not set",
		},
		{
			name:              "flag true without env var",
			envValue:          "",
			createOnlyFlag:    true,
			expectedResult:    true,
			expectedFlagAfter: true,
			description:       "should return true when flag is true",
		},
		{
			name:              "env true, flag false",
			envValue:          "true",
			createOnlyFlag:    false,
			expectedResult:    true,
			expectedFlagAfter: true,
			description:       "should return true and set flag to true when env is true",
		},
		{
			name:              "env true, flag true",
			envValue:          "true",
			createOnlyFlag:    true,
			expectedResult:    true,
			expectedFlagAfter: true,
			description:       "should return true and keep flag true when both are true",
		},
		{
			name:              "env false, flag true",
			envValue:          "false",
			createOnlyFlag:    true,
			expectedResult:    true,
			expectedFlagAfter: true,
			description:       "should return true when flag is true despite env being false",
		},
		{
			name:              "env false, flag false",
			envValue:          "false",
			createOnlyFlag:    false,
			expectedResult:    false,
			expectedFlagAfter: false,
			description:       "should return false when both are false",
		},
		{
			name:              "env invalid value, flag false",
			envValue:          "invalid",
			createOnlyFlag:    false,
			expectedResult:    false,
			expectedFlagAfter: false,
			description:       "should return false when env has invalid value and flag is false",
		},
		{
			name:              "env invalid value, flag true",
			envValue:          "invalid",
			createOnlyFlag:    true,
			expectedResult:    true,
			expectedFlagAfter: true,
			description:       "should return true when flag is true despite invalid env value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(createOnlyEnvName, tt.envValue)
			flagCopy := tt.createOnlyFlag
			result := IsInCreateOnlyMode(&flagCopy)

			if result != tt.expectedResult {
				t.Errorf("IsInCreateOnlyMode() = %v, expected %v - %s", result, tt.expectedResult, tt.description)
			}

			if flagCopy != tt.expectedFlagAfter {
				t.Errorf("createOnlyFlag after call = %v, expected %v - %s", flagCopy, tt.expectedFlagAfter, tt.description)
			}
		})
	}
}
