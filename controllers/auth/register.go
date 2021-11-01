package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dela-dels/go-expenses-tracker/database"
	"github.com/dela-dels/go-expenses-tracker/database/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserRegistrationDetails struct {
	Firstname string `validate:"required"`
	Lastname  string `validate:"required"`
	Email     string `validate:"required,unique-email"`
	Password  string `validate:"required,gte=8"`
}

var validate *validator.Validate

var db, connectionError = database.Connect()

func ShowRegistrationForm(context *gin.Context) {
	context.HTML(http.StatusOK, "registration.html", gin.H{})
}

func Register(context *gin.Context) {

	if connectionError != nil {
		log.Fatalf("could not connect to the database. Error : %s", connectionError)
	}

	db.AutoMigrate(models.User{})

	password, _ := hashPassword(context.PostForm("password"))

	userRegistrationDetails := UserRegistrationDetails{
		context.PostForm("first_name"),
		context.PostForm("last_name"),
		context.PostForm("email"),
		password,
	}

	validationErrors := validatUserRegistrationStruct(userRegistrationDetails)

	if len(validationErrors) > 0 {
		context.HTML(http.StatusOK, "registration.html", gin.H{
			"validation_errors": validationErrors,
		})

	} else {
		db.Create(&models.User{
			Firstname: userRegistrationDetails.Firstname,
			Lastname:  userRegistrationDetails.Lastname,
			Email:     userRegistrationDetails.Email,
			Password:  userRegistrationDetails.Password,
		})

		context.HTML(http.StatusOK, "dashboard.html", gin.H{})
	}
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPassword), err
}

func validatUserRegistrationStruct(userRegistrationDetails UserRegistrationDetails) map[string]string {
	validate = validator.New()
	validate.RegisterValidation("unique-email", validateEmailToBeUnique)

	var userRegistrationValidationErrors = map[string]string{}
	err := validate.Struct(userRegistrationDetails)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "unique-email" {
				userRegistrationValidationErrors[err.Field()] =
					fmt.Sprintf("The %v provided has already been taken", err.Field())
			} else {
				userRegistrationValidationErrors[err.Field()] =
					fmt.Sprintf("The %v field is %v", err.Field(), err.Tag())
			}
		}
	}

	return userRegistrationValidationErrors
}

func validateEmailToBeUnique(fl validator.FieldLevel) bool {
	results := db.Where("email = ?", fl.Field().String()).Find(&models.User{})
	return results.RowsAffected != 1
}
