package data

func ErrorResponse[T any](err error, additionalInfo any, data T) GeneralResponse[T] {
	return GeneralResponse[T]{
		Status: err.Error(), AdditionalInfo: additionalInfo, Data: data,
	}
}

func SuccessResponse[T any](additionalInfo any, data T) GeneralResponse[T] {
	return GeneralResponse[T]{
		Status: "Success", Data: data,
	}
}
