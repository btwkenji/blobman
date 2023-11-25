package data

type Owner struct {
	ID string `db:"id"`
}

type Owners interface {
	New() Owners
	CreateOwner(id string) error
	Exists(id string) error
}
