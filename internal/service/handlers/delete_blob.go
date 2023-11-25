package handlers

import (
	"github.com/kenjitheman/blobman/internal/data"
	"github.com/kenjitheman/blobman/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func DeleteBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteBlobRequest(r)
	if err != nil {
		Log(r).WithError(err).Info("Invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	var blob *data.Blob
	blob, err = BlobQ(r).GetBlobById(request.BlobID)
	if err != nil {
		Log(r).WithError(err).Info("Not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}
	err = BlobQ(r).DeleteBlob(blob.ID)
	if err != nil {
		Log(r).WithError(err).Error("Internal error")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
