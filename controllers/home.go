package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowHomePage(context *gin.Context) {
	context.HTML(http.StatusOK, "home.html", gin.H{})
}
