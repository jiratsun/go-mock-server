package path

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	request, err := decode[pathRequest](r)
	if err != nil {
		fmt.Printf("%v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var dto []pathToHostUpsertMany
	for k, v := range request.PathToHost {
		dto = append(dto, pathToHostUpsertMany{path: k, host: v})
	}

	err = handler.Store.upsertMany(r.Context(), dto)
	if err != nil {
		err = fmt.Errorf("Error inserting to SQL: %w", err)
		fmt.Printf("%v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write([]byte("Success 1000"))
}

func decode[T any](r *http.Request) (T, error) {
	var v T
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		return v, fmt.Errorf("Error decoding JSON: %w", err)
	}
	return v, nil
}
