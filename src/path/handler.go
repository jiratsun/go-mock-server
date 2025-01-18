package path

import "net/http"

type PathHandler struct {
	Store *PathStore
}

func (handler *PathHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
}

func (handler *PathHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
}

func (handler *PathHandler) HandlePost(w http.ResponseWriter, r *http.Request) {
}
