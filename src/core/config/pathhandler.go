package config

import (
	"errors"
	"net/http"

	"mockserver.jiratviriyataranon.io/src/data"
)

type PathHandler struct {
	Store  *PathStore
	getEnv func(string) string
}

func (handler *PathHandler) HandleGetPath(w http.ResponseWriter, r *http.Request) {
	result, err := handler.Store.findAllPath(r.Context())
	if err != nil {
		data.Encode(w, http.StatusInternalServerError, data.ErrorResponse[any](err, nil, nil))
		return
	}

	response := make(getPathResponse)
	for _, path := range result {
		response[path.Path] = &pathInfo{
			DefaultHost: path.DefaultHost, Description: path.Description, IsActive: path.IsActive,
		}
	}

	data.Encode(w, http.StatusOK, data.SuccessResponse(nil, response))
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
			Path:        *path.Path,
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

func (handler *PathHandler) HandleDeletePath(w http.ResponseWriter, r *http.Request) {
	request, err := data.Decode[deletePathRequest](r)
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

	var paths []string
	for _, path := range request.Paths {
		if path.Path != nil {
			paths = append(paths, *path.Path)
		}
	}

	err = handler.Store.deleteManyPath(r.Context(), pathDeleteMany{Path: paths})
	if err != nil {
		data.Encode(w, http.StatusInternalServerError, data.ErrorResponse[any](err, nil, nil))
		return
	}

	data.Encode(w, http.StatusOK, data.SuccessResponse[any](nil, nil))
}
