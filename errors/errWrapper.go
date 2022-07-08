package errors

import (
	errors2 "errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

func IsValidationError(err error) bool {
	_, ok := err.(validator.ValidationErrors)
	return ok
}

func ErrWrapper(handler func(g *gin.Context) (interface{}, error)) func(*gin.Context) {
	return func(context *gin.Context) {
		data, err := handler(context)
		if err != nil {
			var customErr = new(customError)

			if errors2.As(err, &customErr) {
				ResultFail(http.StatusBadRequest, customErr.Code, context)
			} else if IsValidationError(err) {
				ParamBindErrorResult(err.Error(), context)
			} else {
				ResultFail(http.StatusInternalServerError, ServerError, context)
			}
			// 自行更改成特定的log
			log.Printf("url: %s ,erorr: %+v", context.Request.URL, err)

		} else {
			ResultOk(data, context)
		}

	}
}
