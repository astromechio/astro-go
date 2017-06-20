package servlib

import "net/http"

type AResponse struct {
	Body       []byte `json:"body"`
	StatusCode int    `json:statuscode`
}

func (a AResponse) Write(b []byte) (int, error) {
	a.Body = b

	return len(b), nil
}

func (a AResponse) Header() http.Header {
	return http.Header{}
}

func (a AResponse) WriteHeader(code int) {
	a.StatusCode = code
}
