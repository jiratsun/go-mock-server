package data

type GeneralResponse[T any] struct {
	Status         string
	AdditionalInfo any
	Data           T
}
