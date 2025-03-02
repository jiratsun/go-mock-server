package config

import (
	"errors"
	"net/http"

	"mockserver.jiratviriyataranon.io/src/data"
)

type PathHandler struct {
	Store  *ConfigStore
	getEnv func(string) string
}

func (handler *PathHandler) HandleRegisterPath(w http.ResponseWriter, r *http.Request) {
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

	var dto []pathUpsertMany
	for _, path := range request.Paths {
		dto = append(dto, pathUpsertMany{
			Path:        path.Path,
			DefaultHost: data.ToNullString(path.DefaultHost),
			Description: path.Description,
		})
	}

	err = handler.Store.upsertManyPath(r.Context(), dto)
	if err != nil {
		data.Encode(w, http.StatusInternalServerError, data.ErrorResponse[any](err, nil, nil))
		return
	}

	data.Encode(w, http.StatusOK, data.SuccessResponse[any](nil, nil))
}
