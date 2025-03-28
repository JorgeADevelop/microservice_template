package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"www.marawa.com/microservice_service/pkg/responses"
	"www.marawa.com/microservice_service/pkg/validation"
)

func BindAndValidate[V any](ctx *gin.Context, model *V) error {
	if err := ctx.ShouldBindJSON(model); err != nil {
		responses.GenerateResponse(
			ctx,
			responses.ResponseData{
				Code:      400,
				MessageID: "bind_error",
				Error:     err,
			},
		)

		return err
	}

	if err := validation.Validate.Struct(model); err != nil {
		responses.GenerateResponse(
			ctx,
			responses.ResponseData{
				Code:      400,
				MessageID: "validation_error",
				Error:     err,
			},
		)

		return err
	}

	return nil
}

func GetToken(ctx *gin.Context) (string, error) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		err := errors.New("Authorization header not found")
		responses.GenerateResponse(
			ctx,
			responses.ResponseData{
				Code:      401,
				MessageID: "empty_authorization_header",
				Error:     err,
			},
		)

		return "", err
	}

	return token, nil
}
