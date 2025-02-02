package host

import (
	"errors"
	"net/http"

	"mockserver.jiratviriyataranon.io/src/data"
)

type HostHandler struct {
	Store  *HostStore
	getEnv func(string) string
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
