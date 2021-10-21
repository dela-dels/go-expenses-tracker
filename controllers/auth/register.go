package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dela-dels/go-expenses-tracker/database"
	"github.com/dela-dels/go-expenses-tracker/database/models"
	"github.com/dela-dels/go-expenses-tracker/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserRegistrationDetails struct {
	Firstname string `form:"fist_name"`
	Lastname  string `form:"last_name"`
	Email     string `form:"email"`
	Password  string `form:"password"`
}

var registrationError string

func ShowRegistrationForm(context *gin.Context) {
	context.HTML(http.StatusOK, "registration.html", gin.H{})
}

func Register(context *gin.Context) {

	db, err := database.Connect()

	if err != nil {
		fmt.Printf("could not connect to the database. Error : %s", err)
	}

	db.AutoMigrate(models.User{})

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
	})

	if utils.ConvertGormErrorToStruct(results.Error).Code == 1062 {
		registrationError = "The email you provided has already been taken"
	}

	if results.Error != nil {
		context.HTML(http.StatusOK, "registration.html", gin.H{
			"errors": map[string]string{
				"email": registrationError,
			},
		})
	}
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPassword), err
}
