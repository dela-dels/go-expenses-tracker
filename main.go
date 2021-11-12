package main

import (
	"log"
	"os"

	"github.com/dela-dels/go-expenses-tracker/controllers"
	"github.com/dela-dels/go-expenses-tracker/controllers/auth"
	"github.com/dela-dels/go-expenses-tracker/middlewares"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func setup() {
	// load the environment variables for the .env
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Unable to load .env file %v", err)
	}
}

func main() {
	setup()
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*.html")
	router.Static("public/css", "public/css")

	sessionStore := cookie.NewStore([]byte(os.Getenv("APP_KEY")))
	//sessionStore.Options(sessions.Options{MaxAge: 60 * 60 * 24}) // allow cookie to be sotred for 24 hours
	router.Use(sessions.Sessions(os.Getenv("APP_NAME"), sessionStore))

	router.GET("login", auth.ShowLoginForm)
	router.POST("login", auth.Login)
	router.GET("register", auth.ShowRegistrationForm)
	router.POST("register", auth.Register)

	//Routes behind the authentication middleware
	authRoutes := router.Use(middlewares.Authenticated())
	authRoutes.GET("home", controllers.ShowHomePage)

	log.Fatal(router.Run(os.Getenv("APP_PORT")))
}
