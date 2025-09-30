package hashing

type Hashing interface {
	Sum(data []byte) (string, error)
}
