package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func Decode[T any](r *http.Request) (T, error) {
	var v T

	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		return v, fmt.Errorf("Error decoding JSON: %w", err)
	}

	return v, nil
}

func Encode[T any](w http.ResponseWriter, status int, v T) error {
	fmt.Println(v)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		return fmt.Errorf("Error encoding JSON: %w", err)
	}

	return nil
}

func ToNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}

	return sql.NullString{String: s, Valid: true}
}

func ToString(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}
