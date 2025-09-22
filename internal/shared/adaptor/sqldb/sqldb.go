package sqldb

type SqlDB[M any] interface {
	Create(params M) error
	Where(params M) ([]M, error)
	WhereField(column string, value string) (M, error)
	Save(params M) error
	GetByID(id int) (M, error)
	Delete(id int) error
}
