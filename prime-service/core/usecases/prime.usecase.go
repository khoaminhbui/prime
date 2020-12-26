package usecases

import (
	"strconv"

	"github.com/khoaminhbui/prime/prime-service/core/entities"
)

// FindLargestPrimeUnder calculate the largest prime number under an integet
func FindLargestPrimeUnder(ns string) entities.Response {
	response := entities.Response{ErrorCode: 0, Message: "", Value: 0}

	// type conversion
	n, err := strconv.Atoi(ns)

	// usecase validation
	if err != nil {
		response.ErrorCode = entities.ERROR_CODE_INVALID_FORMAT
		response.Message = "Invalid input number"
	} else {
		response.Message = "Success"
		response.Value = entities.TrialDivision(n)
	}

	return response
}
