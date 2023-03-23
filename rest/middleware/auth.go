package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yonisaka/loan-service/grpc/client"
	"github.com/yonisaka/loan-service/rest/dto"
	pbAuth "github.com/yonisaka/protobank/auth"
)

func AuthB2B(client *client.GRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("authorization", c.GetHeader("Authorization"))
		payload := pbAuth.AuthB2BPayload{
			Token: c.GetHeader("Authorization"),
		}
		user, err := client.AuthB2B(c, &payload)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				*dto.NewResponse().WithCode(http.StatusUnauthorized).WithMessage(err.Error()),
			)
			return
		}
		c.Set("username", user.Username)
		c.Set("userId", user.Id)
		c.Next()
	}
}
