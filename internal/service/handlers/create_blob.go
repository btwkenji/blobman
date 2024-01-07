package handlers

import (
	"github.com/btwkenji/blobman/internal/data"
	"github.com/btwkenji/blobman/internal/service/requests"
	"github.com/btwkenji/blobman/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func CreateBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateBlobRequest(r)
	if err != nil {
		Log(r).WithError(err).Info("Invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	blob, err := requests.Blob(request)
	if err != nil {
		Log(r).WithError(err).Error("Failed to create blob")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	ownerId := request.Relationships.Owner
	err = OwnerQ(r).Exists(ownerId)
	if err != nil {
		err = OwnerQ(r).CreateOwner(ownerId)
		if err != nil {
			Log(r).WithError(err).Error("Failed to create blob owner")
			ape.RenderErr(w, problems.Conflict())
			return
		}
	}
	err = BlobQ(r).CreateBlob(blob, data.Owner{ID: ownerId})
	if err != nil {
		Log(r).WithError(err).Error("Failed to create blob")
		ape.RenderErr(w, problems.Conflict())
		return
	}
	response := resources.BlobResponse{
		Data: NewBlob(blob),
	}
	ape.Render(w, response)
}

func NewBlob(blob *data.Blob) resources.Blob {
	b := resources.Blob{
		Key: resources.Key{
			ID: blob.ID,
		},
		Attributes: resources.BlobAttributes{
			Value: blob.Value,
		},
		Relationships: resources.BlobRelationships{
			Owner: blob.OwnerId,
		},
	}
	return b
}
