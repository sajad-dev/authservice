package sqldb

type SqlDB[M any] interface {
	Create(params M) error
	FindWithField(params M) ([]M, error)
	Save(params M) error
	GetByID(id int) (M, error)
	Delete(id int) error
}
