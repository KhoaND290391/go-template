package utilities

import "github.com/google/uuid"
import "regexp"

import "chat-service/pkg/logger"

import "fmt"

func GetUUID() string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		logger.WriteError("Utilities.uuidHelper.go", "GetUUID", "", fmt.Sprintf("%v", err))
	}
	return reg.ReplaceAllString(uuid.New().String(), "")
}
