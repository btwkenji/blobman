package postgres

import (
	"github.com/Masterminds/squirrel"
	"github.com/kenjitheman/blobman/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
	"math"
)

const blobsTableName = "blobs"

type BlobsDB struct {
	db  *pgdb.DB
	sql squirrel.SelectBuilder
}

func NewBlobs(db *pgdb.DB) data.Blobs {
	return &BlobsDB{
		db:  db.Clone(),
		sql: squirrel.Select("*").From(blobsTableName)}
}

func (b *BlobsDB) New() data.Blobs {
	return NewBlobs(b.db)
}

func (b *BlobsDB) GetBlobById(id string) (*data.Blob, error) {
    var result data.Blob
    statement := squirrel.Select("*").From(blobsTableName).Where("id = ?", id)
    err := b.db.Get(&result, statement)
    if err != nil {
        return nil, err
    }
    return &result, nil
}

func (b *BlobsDB) CreateBlob(blob *data.Blob, owner data.Owner) error {
	statement := squirrel.Insert(blobsTableName).SetMap(map[string]interface{}{
		"id":       blob.ID,
		"owner_id": owner.ID,
		"value":    blob.Value,
	})
	err := b.db.Exec(statement)
	return err
}

func (b *BlobsDB) DeleteBlob(id string) error {
	statement := squirrel.Delete(blobsTableName).Where(squirrel.Eq{"id": id})
	err := b.db.Exec(statement)
	return err
}

func (b *BlobsDB) GetTotalPages(limit uint64) (uint64, error) {
	res := make([]data.Blob, 0)
	b.sql = squirrel.Select("*").From(blobsTableName)
	err := b.db.Select(&res, b.sql)
	if err != nil {
		return 0, err
	}
	return (uint64)(math.Ceil((float64)(len(res)) / float64(limit))), nil
}


func (b *BlobsDB) GetPage(pageParams pgdb.OffsetPageParams) data.Blobs {
	b.sql = pageParams.ApplyTo(b.sql, "id")
	return b
}

func (b *BlobsDB) GetBlobsList() ([]data.Blob, error) {
	result := make([]data.Blob, 0)
	b.sql = squirrel.Select("*").FromSelect(b.sql, "blobs_by_page")
	err := b.db.Select(&result, b.sql)
	if err != nil {
		return nil, err
	}
	return result, nil
}
