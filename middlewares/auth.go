package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Authenticated this function is a middleware that runs to check if the user
//making the current request is authenticated or not (has a session)
func Authenticated() gin.HandlerFunc {
	return func(context *gin.Context) {
		_, cookieError := context.Cookie(os.Getenv("APP_NAME"))

		if cookieError != nil {
			context.Redirect(http.StatusPermanentRedirect, "/login")
			context.Abort()
		}

		context.Next()
	}
}
