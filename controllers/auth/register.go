package auth

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dela-dels/go-expenses-tracker/database"
	"github.com/dela-dels/go-expenses-tracker/database/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserRegistrationDetails struct {
	Firstname string `form:"fist_name"`
	Lastname  string `form:"last_name"`
	Email     string `form:"email"`
	Password  string `form:"password"`
}

func ShowRegistrationForm(context *gin.Context) {
	context.HTML(http.StatusOK, "registration.html", gin.H{})
}

func Register(context *gin.Context) {

	db, err := database.Connect()

	if err != nil {
		fmt.Printf("could not connect to the database. Error : %s", err)
	}

	db.AutoMigrate(models.User{})

	if err := context.Request.ParseForm(); err != nil {
		log.Fatal(err)
	}

	password, err := hashPassword(context.PostForm("password"))

	if err != nil {
		log.Fatal("unable to hash password")
	}

	userRegistrationDetails := UserRegistrationDetails{
		context.PostForm("first_name"),
		context.PostForm("last_name"),
		context.PostForm("email"),
		password,
	}

	results := db.Create(&models.User{
		Firstname: userRegistrationDetails.Firstname,
		Lastname:  userRegistrationDetails.Lastname,
		Email:     userRegistrationDetails.Email,
		Password:  userRegistrationDetails.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	log.Println(results)
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPassword), err
}
