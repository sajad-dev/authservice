package request

type AuthorizeRequest struct {
	Headers  map[string]string
	Host     string
	Body     string
	Protocol string
	Method   string
	Path     string
	RawBody  []byte
}
