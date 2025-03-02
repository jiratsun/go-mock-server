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
	request, err := data.Decode[registerHostRequest](r)
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

func (handler *HostHandler) HandleDeleteHost(w http.ResponseWriter, r *http.Request) {
	request, err := data.Decode[deleteHostRequest](r)
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

	var both []data.Tuple2[string, string]
	var domainName []string
	var alias []string
	for _, host := range request.Hosts {
		switch {
		case host.DomainName != nil && host.Alias != nil:
			both = append(both, data.Pair(*host.DomainName, *host.Alias))
		case host.DomainName != nil:
			domainName = append(domainName, *host.DomainName)
		case host.Alias != nil:
			alias = append(alias, *host.Alias)
		}
	}

	err = handler.Store.deleteManyHost(r.Context(), hostDeleteMany{
		Both: both, DomainName: domainName, Alias: alias,
	})
	if err != nil {
		data.Encode(w, http.StatusInternalServerError, data.ErrorResponse[any](err, nil, nil))
		return
	}

	data.Encode(w, http.StatusOK, data.SuccessResponse[any](nil, nil))
}
