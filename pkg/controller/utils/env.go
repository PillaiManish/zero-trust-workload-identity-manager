package utils

import (
	"os"
)

const (
	createOnlyEnvName        = "CREATE_ONLY_MODE"
	CreateOnlyModeStatusType = "CreateOnlyMode"
	CreateOnlyModeEnabled    = "CreateOnlyModeEnabled"
	CreateOnlyModeDisabled   = "CreateOnlyModeDisabled"
)

func IsCreateOnlyEnvEnabled() bool {
	createOnlyEnvValue := os.Getenv(createOnlyEnvName)
	return createOnlyEnvValue == "true"
}

func IsInCreateOnlyMode(createOnlyFlag *bool) bool {
	currentCreateOnly := IsCreateOnlyEnvEnabled()
	if currentCreateOnly {
		*createOnlyFlag = true
		return true
	}
	if *createOnlyFlag {
		return true
	}
	return false
}
