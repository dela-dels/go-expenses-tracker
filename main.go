package main

import (
	"log"
	"net/http"

	"github.com/dela-dels/go-expenses-tracker/controllers/auth"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*.html")
	router.Static("public/css", "public/css")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{})
	})
	router.GET("/login", auth.ShowLoginForm)
	router.POST("/login", auth.Login)
	router.GET("register", auth.ShowRegistrationForm)
	router.POST("register", auth.Register)

	log.Fatal(router.Run(":8080"))
}
