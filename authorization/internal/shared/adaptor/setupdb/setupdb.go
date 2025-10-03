package setupdb

type SetupDB[M any] interface {
	Connection() (M, error)
}
