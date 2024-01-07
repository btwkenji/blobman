package handlers

import (
	"github.com/btwkenji/blobman/internal/service/requests"
	"github.com/btwkenji/blobman/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func GetBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetBlobRequest(r)
	if err != nil {
		Log(r).WithError(err).Info("Invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	blob, err := BlobQ(r).GetBlobById(request.BlobID)
	if err != nil {
		Log(r).WithError(err).Error("Failed to get blob")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if blob == nil {
		Log(r).WithError(err).Error("Blob not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}
	response := resources.BlobResponse{
		Data: NewBlob(blob),
	}
	ape.Render(w, &response)
	w.WriteHeader(http.StatusOK)
}
