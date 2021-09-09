package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func showLoginForm(context *gin.Context) {
	context.HTML(http.StatusOK, "app.html", gin.H{})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("public/css", "public/css")

	database, error := sql.Open("mysql", "root:@tcp(localhost:3306)/go_expenses")

	if error != nil {
		log.Fatal("Could not connect to database. Error", error)
	}

	println(database)

	// router.GET("/", func(context *gin.Context) {
	// 	context.HTML(http.StatusOK, "app.html", gin.H{})
	// })

	router.GET("/", showLoginForm)
	// router.GET("register", ShowLO)

	log.Fatal(router.Run(":8080"))
}
