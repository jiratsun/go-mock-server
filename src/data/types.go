package data

import (
	"encoding/json"
	"errors"
)

type GeneralResponse[T any] struct {
	Status         string `json:"status,omitempty"`
	AdditionalInfo any    `json:"additionalInfo,omitempty"`
	Data           T      `json:"data,omitempty"`
}

type Tuple2[L, R any] struct {
	Left  L
	Right R
}

func Pair[L, R any](left L, right R) Tuple2[L, R] {
	return Tuple2[L, R]{left, right}
}

type StringOrSlice []string

func (s *StringOrSlice) UnmarshalJSON(jsonValue []byte) error {
	var str string
	err := json.Unmarshal(jsonValue, &str)
	if err == nil {
		*s = []string{str}
		return nil
	}

	var strSlice []string
	err = json.Unmarshal(jsonValue, &strSlice)
	if err == nil {
		*s = strSlice
		return nil
	}

	return errors.New("Error deserializing StringOrSlice")
}
