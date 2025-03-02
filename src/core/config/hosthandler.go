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

func (handler *HostHandler) HandleGetHost(w http.ResponseWriter, r *http.Request) {
	result, err := handler.Store.findAllHost(r.Context())
	if err != nil {
		data.Encode(w, http.StatusInternalServerError, data.ErrorResponse[any](err, nil, nil))
		return
	}

	response := make(getHostResponse)
	for _, host := range result {
		response[host.DomainName] = &hostInfo{
			Alias: host.Alias, Description: host.Description, IsActive: host.IsActive,
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

	var dto []hostUpsertMany
	for _, host := range request.Hosts {
		dto = append(dto, hostUpsertMany{
			DomainName:  host.DomainName,
			Alias:       host.Alias,
			Description: host.Description,
		})
	}

	err = handler.Store.upsertManyHost(r.Context(), dto)
	if err != nil {
		data.Encode(w, http.StatusInternalServerError, data.ErrorResponse[any](err, nil, nil))
		return
	}

	data.Encode(w, http.StatusOK, data.SuccessResponse[any](nil, nil))
}
