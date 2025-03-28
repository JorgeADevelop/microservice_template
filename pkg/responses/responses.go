package responses

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"www.marawa.com/microservice_service/pkg/translater"
	"www.marawa.com/microservice_service/pkg/validation"
)

type ResponseData struct {
	MessageID string `json:"message_id"`
	Data      any    `json:"data"`
	Error     error  `json:"error"`
	Code      int    `json:"status"`
}

type HTTPResponse struct {
	Message string       `json:"message"`
	Data    any          `json:"data"`
	Error   *string      `json:"error"`
	Errors  []HTTPErrors `json:"errors"`
	Status  string       `json:"status"`
	Code    int          `json:"code"`
}

type HTTPErrors struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GenerateResponse(ctx *gin.Context, responseData ResponseData) {
	lang := getLang(ctx)
	message := translater.TranslateMessage(lang, responseData.MessageID)

	switch responseData.Code {
	case 201:
		ctx.JSON(201, HTTPResponse{
			Message: message,
			Data:    responseData.Data,
			Error:   nil,
			Status:  "created",
			Code:    201,
		})
	case 204:
		ctx.JSON(204, HTTPResponse{
			Message: message,
			Data:    nil,
			Error:   nil,
			Status:  "no content",
			Code:    204,
		})
	case 400:
		var errors []HTTPErrors
		if validationErrors, ok := responseData.Error.(validator.ValidationErrors); ok {
			var translator ut.Translator
			switch lang {
			case "es":
				translator = validation.TransES
			default:
				translator = validation.TransEN
			}
			for _, fieldError := range validationErrors {
				errors = append(errors, HTTPErrors{
					Field:   fieldError.Field(),
					Message: fieldError.Translate(translator),
				})
			}
		}

		errString := responseData.Error.Error()

		ctx.JSON(400, HTTPResponse{
			Message: message,
			Data:    nil,
			Error:   &errString,
			Errors:  errors,
			Status:  "bad request",
			Code:    400,
		})
	case 401:
		errString := responseData.Error.Error()

		ctx.JSON(401, HTTPResponse{
			Message: message,
			Data:    nil,
			Error:   &errString,
			Status:  "unauthorized",
			Code:    401,
		})
	case 403:
		errString := responseData.Error.Error()

		ctx.JSON(403, HTTPResponse{
			Message: message,
			Data:    nil,
			Error:   &errString,
			Status:  "forbidden",
			Code:    403,
		})
	case 404:
		errString := responseData.Error.Error()

		ctx.JSON(404, HTTPResponse{
			Message: message,
			Data:    nil,
			Error:   &errString,
			Status:  "not found",
			Code:    404,
		})
	case 500:
		errString := responseData.Error.Error()

		ctx.JSON(500, HTTPResponse{
			Message: message,
			Data:    nil,
			Error:   &errString,
			Status:  "internal server error",
			Code:    500,
		})
	default:
		ctx.JSON(200, HTTPResponse{
			Message: message,
			Data:    responseData.Data,
			Error:   nil,
			Status:  "success",
			Code:    200,
		})
	}
}
func getLang(ctx *gin.Context) string {
	lang := ctx.GetHeader("Accept-Language")
	if lang == "" {
		return "en"
	}
	return lang
}
