package request

type AuthorizeRequest struct {
	Id       string
	Headers  map[string]string
	Host     string
	Body     string
	Protocol string
	Method   string
	Path     string
	Size     string
	RawBody  []byte
}
