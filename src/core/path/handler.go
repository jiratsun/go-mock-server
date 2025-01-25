package path

import (
	"fmt"
	"net/http"

	"mockserver.jiratviriyataranon.io/src/data"
)

type PathHandler struct {
	Store  *PathStore
	getEnv func(string) string
}

func (handler *PathHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
}

func (handler *PathHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
}

func (handler *PathHandler) HandleRegisterPathToHost(w http.ResponseWriter, r *http.Request) {
	request, err := data.Decode[registerPathRequest](r)
	if err != nil {
		data.Encode(w, http.StatusBadRequest, err.Error())
		return
	}

	problems := request.valid(r.Context())
	if len(problems) > 0 {
		data.Encode(w, http.StatusBadRequest, problems)
		return
	}

	var dto []pathToHostUpsertMany
	for k, v := range request.PathToHost {
		dto = append(dto, pathToHostUpsertMany{path: k, host: v})
	}

	err = handler.Store.upsertMany(r.Context(), dto)
	if err != nil {
		err = fmt.Errorf("Error upserting SQL: %w", err)
		data.Encode(w, http.StatusInternalServerError, err.Error())
		return
	}

	data.Encode(w, http.StatusOK, "Success")
}
