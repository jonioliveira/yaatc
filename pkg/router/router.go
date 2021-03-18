package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jonioliveira/yaatc/internal/loan"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

const (
	v1 = "/v1"
	v2 = "/v2"
)

func BuildRouter() *gin.Engine {
	router := gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	v1 := router.Group(v1)
	{
		loanGroup := v1.Group("/loan")
		{
			loanGroup.GET("/", loan.GetLoanHandler)
			loanGroup.POST("/", loan.PostLoanHandler)
		}
	}

	return router
}
