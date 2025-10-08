package sqldb

type SqlDB[M any] interface {
	Create(params M) error
	Where(params M) ([]M, error)
	WhereField(column string, value string) (M, error)
	Save(params M) error
	GetByID(id int) (M, error)
	Delete(id int) error
	RemoveExpierd(column string, id int) error
}

type SqlDBGlobal interface {
	Exists(tablename string, fieldname, value string) bool
}
