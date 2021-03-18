package loan

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLoanHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "get loan")
}

func PostLoanHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "post loan")
}
