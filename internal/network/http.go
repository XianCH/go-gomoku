package network

type HTTPserver struct {
	address string
}

func NewHTTPServer(add string) *HTTPserver {
	return &HTTPserver{address: add}
}

func (s *HTTPserver) Start() {
}
