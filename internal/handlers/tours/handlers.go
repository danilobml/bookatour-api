package tours

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTours(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{ "id": "1", "name": "Test tour" })
}