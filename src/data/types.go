package data

import (
	"encoding/json"
	"errors"
)

type GeneralResponse[T any] struct {
	Status         string
	AdditionalInfo any
	Data           T
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
