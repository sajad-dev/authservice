package sqldb

type SqlDB[M any] interface {
	Create(params M) error
	Save(params M) error
	Read(id int) (M, error)
	Delete(id int) error
}
