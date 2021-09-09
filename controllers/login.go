package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowLoginForm(context *gin.Context) {
	context.HTML(http.StatusOK, "app.html", gin.H{})
}
