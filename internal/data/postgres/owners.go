package postgres

import (
	"github.com/Masterminds/squirrel"
	"github.com/kenjitheman/blobman/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const ownersTableName = "owners"

type OwnersDB struct {
	db  *pgdb.DB
	sql squirrel.SelectBuilder
}

func NewOwners(db *pgdb.DB) data.Owners {
	return &OwnersDB{
		db:  db.Clone(),
		sql: squirrel.Select("*").From(ownersTableName)}
}

func (b *OwnersDB) CreateOwner(ownerID string) error {
	statement := squirrel.Insert(ownersTableName).SetMap(map[string]interface{}{
		"id": ownerID,
	})
	err := b.db.Exec(statement)
	return err
}

func (b *OwnersDB) Exists(id string) error {
	statement := squirrel.Select("id").From(ownersTableName).Where(squirrel.Eq{"id": id})
	var ownerID string
	err := b.db.Get(&ownerID, statement)
	return err
}

func (b *OwnersDB) New() data.Owners {
	return NewOwners(b.db)
}
