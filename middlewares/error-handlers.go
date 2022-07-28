package middlewares

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {

	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		// Override status code if fiber.Error type
		code = e.Code
	}
	if code == 404 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"responseCode": 4004,
			"message":      "Can't found API",
		})
	}

	var errMessage string = err.Error() // Get error message from panic
	if errMessage != "" && code != 404 {
		var finalMessage string = "Request failed"
		var errorCode string = "ERR001"
		var finalStatusCode int = fiber.StatusInternalServerError
		splitErrorMessage := strings.Split(errMessage, "::") // Split errMessage from "388:message" <- Example

		if len(splitErrorMessage) == 2 {
			finalStatusCode, err = strconv.Atoi(splitErrorMessage[0]) // Get error code from splitErrorMessage[0] "400::message"
			if err != nil {
				finalStatusCode = fiber.StatusInternalServerError
			}
			finalMessage = splitErrorMessage[1]
			splitCodeError := strings.Split(splitErrorMessage[0], "-") // Split splitErrorMessage[0] from "400-4000:message" <- Example
			if len(splitCodeError) == 2 {
				errorCode = splitCodeError[1]
				finalStatusCode, err = strconv.Atoi(splitCodeError[0])
				if err != nil {
					finalStatusCode = fiber.StatusInternalServerError
				}
			}
		}

		return c.Status(finalStatusCode).JSON(fiber.Map{
			"responseCode": errorCode,
			"message":      finalMessage,
		})
	}
	return nil
}
