package config

import (
	"errors"
	"net/http"

	"mockserver.jiratviriyataranon.io/src/data"
)

type HostHandler struct {
	Store  *ConfigStore
	getEnv func(string) string
}

func (handler *HostHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	result, err := handler.Store.findAllWithPath(r.Context())
	if err != nil {
		data.Encode(w, http.StatusInternalServerError, data.ErrorResponse[any](err, nil, nil))
		return
	}

	response := make(getHostResponse)
	for _, hostWithPath := range result {
		aliasInfo, exist := response[hostWithPath.alias]
		if !exist {
			aliasInfo = &hostInfo{Host: hostWithPath.host, IsActive: hostWithPath.isActive, Paths: []pathInfo{}}
			response[hostWithPath.alias] = aliasInfo
		}

		if hostWithPath.path.Valid && hostWithPath.pathIsActive.Valid {
			aliasInfo.Paths = append(
				aliasInfo.Paths,
				pathInfo{Path: hostWithPath.path.String, IsActive: hostWithPath.pathIsActive.Bool},
			)
		}
	}

	data.Encode(w, http.StatusOK, data.SuccessResponse(nil, response))
}

func (handler *HostHandler) HandleRegisterHost(w http.ResponseWriter, r *http.Request) {
	request, err := data.Decode[registerhostRequest](r)
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

	var dto []aliasToHostUpsertMany
	for k, v := range request {
		dto = append(dto, aliasToHostUpsertMany{alias: k, host: v})
	}

	err = handler.Store.upsertMany(r.Context(), dto)
	if err != nil {
		data.Encode(w, http.StatusInternalServerError, data.ErrorResponse[any](err, nil, nil))
		return
	}

	data.Encode(w, http.StatusOK, data.SuccessResponse[any](nil, nil))
}
