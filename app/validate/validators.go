package app

import (
	"app/app/models"
	"errors"
)

func ValidateJsonData(payload *models.Payload) error {
	if payload.BuildID == "" || payload.Status == "" || payload.Timestamp == "" {
		return errors.New("Missing required fields")
	}
	return nil
}
