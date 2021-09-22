package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserLoginDetails struct {
	User     string `form:"user"`
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

	log.Println(loginDetails)
}

// func checkPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }
