package auth

import (
	"github.com/dela-dels/go-expenses-tracker/database/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)

type UserLoginDetails struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func ShowLoginForm(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", gin.H{})
}

func Login(context *gin.Context) {

	if err := context.Request.ParseForm(); err != nil {
		log.Fatal(err)
		return
	}

	loginDetails := UserLoginDetails{
		context.PostForm("email"),
		context.PostForm("password"),
	}

	var user models.User
	results := db.Where("email = ?", loginDetails.Email).First(&user)

	if results.RowsAffected == 1 {
		err := checkPasswordHash(loginDetails.Password, user.Password)
		if err != nil {
			context.HTML(http.StatusPermanentRedirect, "login.html", gin.H{
				"error": "Your login credentials are incorret.",
			})
		}
	}
	session := sessions.Default(context)
	sessionValue, _ := uuid.NewRandom()
	session.Set(os.Getenv("APP_NAME"), sessionValue.String())
	session.Save()

	context.SetCookie(os.Getenv("APP_NAME"), sessionValue.String(), time.Now().Hour(), "/", os.Getenv("APP_URL"), true, true)
	context.Redirect(http.StatusFound, "home")
}

func checkPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
