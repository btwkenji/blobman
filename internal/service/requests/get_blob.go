package requests

import (
    "github.com/go-chi/chi"
    validation "github.com/go-ozzo/ozzo-validation/v4"
    "net/http"
)

type GetBlobRequest struct {
    BlobID string `json:"-"`
}

func NewGetBlobRequest(r *http.Request) (GetBlobRequest, error) {
    req := GetBlobRequest{
        BlobID: chi.URLParam(r, "id"),
    }
    return req, req.Validate()
}

func (r GetBlobRequest) Validate() error {
    err := validation.Errors{
        "blob": validation.Validate(&r.BlobID, validation.Required),
    }
    return err.Filter()
}
