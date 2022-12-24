package utils

import (
	"encoding/json"
	"house-of-gulmohar/internal/model"
	"net/http"

	"github.com/gofrs/uuid"
)

func ValidateUUID(id string) error {
	_, err := uuid.FromString(id)
	return err
}

func SendResponse(w http.ResponseWriter, res *model.Response) {
	w.WriteHeader(res.Code)
	json.NewEncoder(w).Encode(res.Send())
}
