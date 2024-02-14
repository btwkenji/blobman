package handlers

import (
	"github.com/kenjitheman/blobman/internal/data"
	"github.com/kenjitheman/blobman/internal/service/requests"
	"github.com/kenjitheman/blobman/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/kit/pgdb"
	"net/http"
	"strconv"
)

const (
	pageParamLimit  = "page[limit]"
	pageParamNumber = "page[number]"
	pageParamOrder  = "page[order]"
)

func GetBlobsList(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetBlobsListRequest(r)
	if err != nil {
		Log(r).WithError(err).Info("Invalid request")
		ape.RenderErr(w, problems.InternalError())
	}
	var blobs []data.Blob
	blobs, err = BlobQ(r).GetPage(request.OffsetPageParams).GetBlobsList()
	if err != nil {
		Log(r).WithError(err).Error("Internal error")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if blobs == nil {
		Log(r).WithError(err).Error("Not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}
	totalPages, er := BlobQ(r).GetTotalPages(request.Limit)
	if er != nil {
		return
	}
	response := resources.BlobListResponse{
		Data:  NewBlobsList(blobs),
		Links: GetOffsetLinks(r, request.OffsetPageParams, totalPages),
	}
	ape.Render(w, &response)
}

func NewBlobsList(blobsList []data.Blob) []resources.Blob {
	r := make([]resources.Blob, 0, len(blobsList))
	for _, blobs := range blobsList {
		r = append(r, resources.Blob{
			Key: resources.Key{
				ID: blobs.ID,
			},
			Attributes: resources.BlobAttributes{
				Value: blobs.Value,
			},
			Relationships: resources.BlobRelationships{
				Owner: blobs.OwnerId,
			},
		})
	}
	return r
}

func GetOffsetLinks(r *http.Request, p pgdb.OffsetPageParams, totalPages uint64) *resources.Links {
	result := resources.Links{
		Self:  getOffsetLink(r, p.PageNumber, p.Limit, p.Order),
		First: getOffsetLink(r, 0, p.Limit, p.Order),
		Last:  getOffsetLink(r, totalPages-1, p.Limit, p.Order),
	}
	if p.PageNumber > 0 && p.PageNumber <= totalPages-1 {
		result.Prev = getOffsetLink(r, p.PageNumber-1, p.Limit, p.Order)
	}
	if p.PageNumber < totalPages-1 && p.PageNumber >= 0 {
		result.Next = getOffsetLink(r, p.PageNumber+1, p.Limit, p.Order)
	}
	return &result
}

func getOffsetLink(r *http.Request, pageNumber, limit uint64, order string) string {
	query := r.URL.Query()
	query.Set(pageParamNumber, strconv.FormatUint(pageNumber, 10))
	query.Set(pageParamLimit, strconv.FormatUint(limit, 10))
	query.Set(pageParamOrder, order)
	r.URL.RawQuery = query.Encode()
	return r.URL.String()
}
