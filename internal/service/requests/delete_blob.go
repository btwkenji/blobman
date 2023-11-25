package requests

import (
    "github.com/go-chi/chi"
    validation "github.com/go-ozzo/ozzo-validation/v4"
    "net/http"
)

type DeleteBlobRequest struct {
    BlobID string `json:"-"`
}

func (r DeleteBlobRequest) Validate() error {
    err := validation.Errors{
        "blob": validation.Validate(&r.BlobID, validation.Required),
    }
    return err.Filter()
}

func NewDeleteBlobRequest(r *http.Request) (DeleteBlobRequest, error) {
    req := DeleteBlobRequest{
        BlobID: chi.URLParam(r, "id"),
    }
    return req, req.Validate()
}
