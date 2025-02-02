package path

import (
	"errors"
	"net/http"

	"mockserver.jiratviriyataranon.io/src/data"
)

type PathHandler struct {
	Store  *PathStore
	getEnv func(string) string
}

func (handler *PathHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
}

func (handler *PathHandler) HandleRegisterPathToHost(w http.ResponseWriter, r *http.Request) {
	request, err := data.Decode[registerPathRequest](r)
	if err != nil {
		data.Encode(w, http.StatusBadRequest, data.ErrorResponse[any](err, nil, nil))
		return
	}

	problems := request.valid(r.Context())
	if len(problems) > 0 {
		err = errors.New("Invalid request body")
		data.Encode(w, http.StatusBadRequest, data.ErrorResponse[any](err, problems, nil))
		return
	}

	var dto []pathToHostUpsertMany
	for hostAlias, paths := range request {
		for _, path := range paths {
			dto = append(dto, pathToHostUpsertMany{path: path, hostAlias: hostAlias})
		}
	}

	err = handler.Store.upsertMany(r.Context(), dto)
	if err != nil {
		data.Encode(w, http.StatusInternalServerError, data.ErrorResponse[any](err, nil, nil))
		return
	}

	data.Encode(w, http.StatusOK, data.SuccessResponse[any](nil, nil))
}
