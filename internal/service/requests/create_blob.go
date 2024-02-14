package requests

import (
	"encoding/json"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/kenjitheman/blobman/internal/data"
	"github.com/kenjitheman/blobman/resources"
	"net/http"
)

func NewCreateBlobRequest(r *http.Request) (resources.Blob, error) {
	var req resources.BlobResponse
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return req.Data, errors.New("Failed to decode request body")
	}
	return req.Data, ValidateCreateBlobRequest(req.Data)
}

func ValidateCreateBlobRequest(r resources.Blob) error {
	return validation.Errors{
		"/data/attributes/value":    validation.Validate(&r.Attributes.Value, validation.Required),
		"/data/relationships/owner": validation.Validate(&r.Relationships.Owner, validation.Required),
	}.Filter()
}

func generateUUID() string {
	return uuid.New().String()
}

func Blob(r resources.Blob) (*data.Blob, error) {
	return &data.Blob{
		ID:      generateUUID(),
		Value:   r.Attributes.Value,
		OwnerId: r.Relationships.Owner,
	}, nil
}
