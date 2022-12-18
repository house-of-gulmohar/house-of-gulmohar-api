package utils

import "github.com/gofrs/uuid"

func ValidateUUID(id string) error {
	_, err := uuid.FromString(id)
	return err
}
