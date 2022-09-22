package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dela-dels/go-expenses-tracker/database"
	"github.com/dela-dels/go-expenses-tracker/database/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRegistrationDetails struct {
	Firstname string `validate:"required"`
	Lastname  string `validate:"required"`
	Email     string `validate:"required,unique-email"`
	Password  string `validate:"required,gte=8"`
}

var db = database.Connect()

func ShowRegistrationForm(context *gin.Context) {
	context.HTML(http.StatusOK, "registration.html", gin.H{})
}

func Register(context *gin.Context) {
	err := db.AutoMigrate(models.User{})
	if err != nil {
		fmt.Printf("Unable to migrate tables %s", err)
	}

	password, _ := hashPassword(context.PostForm("password"))

	userRegistrationDetails := UserRegistrationDetails{
		context.PostForm("first_name"),
		context.PostForm("last_name"),
		context.PostForm("email"),
		password,
	}

	validationErrors := userRegistrationDetails.validate()

	if len(validationErrors) > 0 {
		context.HTML(http.StatusPermanentRedirect, "registration.html", gin.H{
			"validation_errors": validationErrors,
		})

	} else {
		db.Create(&models.User{
			Firstname: userRegistrationDetails.Firstname,
			Lastname:  userRegistrationDetails.Lastname,
			Email:     userRegistrationDetails.Email,
			Password:  userRegistrationDetails.Password,
		})
	}

	session := sessions.Default(context)
	randomSessionValue, _ := uuid.NewRandom()
	session.Set(os.Getenv("APP_NAME"), randomSessionValue.String())
	err = session.Save()
	if err != nil {
		fmt.Printf("session could not be established or saved. err: %s", err)
	}

	context.SetCookie(os.Getenv("APP_NAME"), randomSessionValue.String(), time.Now().Hour(), "/", os.Getenv("APP_URL"), true, true)
	context.Redirect(http.StatusFound, "home")
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPassword), err
}

func (u UserRegistrationDetails) validate() map[string]string {
	validate := validator.New()
	err := validate.RegisterValidation("unique-email", validateEmailToBeUnique)
	if err != nil {
		return nil
	}

	var userRegistrationValidationErrors = map[string]string{}
	err = validate.Struct(u)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			// Check to see of the provided email already exists in storage
			// and pass the error as part of the validation errors
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
