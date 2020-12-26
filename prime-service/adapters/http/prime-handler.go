package http

import (
	"net/http"

	"github.com/khoaminhbui/prime/prime-service/core/entities"
	"github.com/khoaminhbui/prime/prime-service/core/usecases"

	"github.com/labstack/echo"
)

// InitHandler set the http request handlers
func InitHandler(e *echo.Echo) {
	e.GET("/prime/:n", handlePrime)
}

func handlePrime(c echo.Context) error {
	ns := c.Param("n")
	response := entities.Response{ErrorCode: 0, Message: "", Value: entities.ERROR_CODE_OK}

	// adapter validation
	if ns == "" {
		response.ErrorCode = entities.ERROR_CODE_EMPTY_INPUT
		response.Message = "Empty number"
	} else {
		response = usecases.FindLargestPrimeUnder(ns)
	}

	return c.JSON(http.StatusOK, response)
}
