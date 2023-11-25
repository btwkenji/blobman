package data

import (
	"encoding/json"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Blob struct {
	ID      string          `db:"id"`
	OwnerId string          `db:"owner_id"`
	Value   json.RawMessage `db:"value"`
}

type Blobs interface {
	New() Blobs
	CreateBlob(blob *Blob, owner Owner) error
	DeleteBlob(id string) error
	GetBlobById(id string) (*Blob, error)
	GetPage(pageParams pgdb.OffsetPageParams) Blobs
	GetBlobsList() ([]Blob, error)
	GetTotalPages(limit uint64) (uint64, error)
}
