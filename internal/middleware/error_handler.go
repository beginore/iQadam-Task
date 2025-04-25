package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ums/internal/dto"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		for _, err := range ctx.Errors {
			switch e := err.Err.(type) {
			case *dto.APIError:
				ctx.JSON(e.StatusCode, dto.ErrorResponse{Error: e.Message})
			default:
				ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Internal server error"})
			}
		}
	}
}
