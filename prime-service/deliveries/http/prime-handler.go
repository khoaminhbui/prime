package http

import (
	"net/http"

	"github.com/khoaminhbui/prime/prime-service/entities"
	"github.com/khoaminhbui/prime/prime-service/usecases"

	"github.com/labstack/echo"
)

// InitHandler set the http request handlers
func InitHandler(e *echo.Echo) {
	e.GET("/prime/:n", handlePrime)
}

func handlePrime(c echo.Context) error {
	n := c.Param("n")
	response := entities.Response{ErrorCode: 0, Message: "", Value: entities.ERROR_CODE_OK}

	// adapter validation
	if n == "" {
		response.ErrorCode = entities.ERROR_CODE_EMPTY_INPUT
		response.Message = "Empty number"
	} else {
		response = usecases.FindLargestPrimeUnder(n)
	}

	return c.JSON(http.StatusOK, response)
}
