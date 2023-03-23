package route

import (
	"github.com/gin-gonic/gin"
	"github.com/yonisaka/loan-service/config"
	"github.com/yonisaka/loan-service/domain/service"
	"github.com/yonisaka/loan-service/grpc/client"
	"github.com/yonisaka/loan-service/rest/handler"
	"github.com/yonisaka/loan-service/rest/middleware"
)

// WithConfig is function
func WithConfig(config *config.Config) RouterOption {
	return func(r *Router) {
		r.config = config
	}
}

// WithRepository is function
func WithRepository(repo *service.Repositories) RouterOption {
	return func(r *Router) {
		r.repo = repo
	}
}

// WithGRPCClient is function
func WithGRPCClient(client *client.GRPCClient) RouterOption {
	return func(r *Router) {
		r.client = client
	}
}

// Init is a function
func (r *Router) Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	e := gin.Default()

	hand := handler.NewHandler(r.repo, r.client)

	loan := handler.NewLoanHandler(hand)

	e.Use(middleware.Logger())
	e.Use(middleware.SaveHttpLog(r.client))

	api := e.Group("/api")
	api.Use(middleware.AuthB2B(r.client))
	api.GET("/loans", loan.GetLoanList)
	api.POST("/loans", loan.CreateLoan)

	return e
}
