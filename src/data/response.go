package data

func ErrorResponse(err error, additionalInfo any) GeneralResponse {
	return GeneralResponse{Status: err.Error(), AdditionalInfo: additionalInfo}
}

func SuccessResponse() GeneralResponse {
	return GeneralResponse{Status: "Success"}
}
