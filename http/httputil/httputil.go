package httputil

import (
	"mime"
	"net/http"
	"strings"
)

// Content Type is not supported,
// then set response http status to 415, Unsupported Media Type
func HandleUnsupportedContentType(w http.ResponseWriter, r *http.Request, mimetype string) bool {

	contentType := r.Header.Get("content-type")

	for _, v := range strings.Split(contentType, ",") {
		t, _, err := mime.ParseMediaType(v)
		if err != nil {
			break
		}
		if t == mimetype {
			return false
		}
	}

	http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
	return true
}

// Accept request header is not supported,
// then set response http status to 406, Not Acceptable
func HandleNotAcceptable(w http.ResponseWriter, r *http.Request, mimetype string) bool {

	accept := r.Header.Get("accept")

	for _, v := range strings.Split(accept, ",") {
		t, _, err := mime.ParseMediaType(v)
		if err != nil {
			break
		}
		if t == mimetype || t == "*/*" {
			return false
		}
	}

	http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
	return true
}
