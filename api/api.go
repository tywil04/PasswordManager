package api

import (
	"github.com/gin-gonic/gin"

	"PasswordManager/api/endpoints/auth/signin"
	"PasswordManager/api/endpoints/auth/signout"
	"PasswordManager/api/endpoints/auth/signup"
	"PasswordManager/api/endpoints/auth/test"
	"PasswordManager/api/endpoints/email"
	"PasswordManager/api/endpoints/emoji"
	"PasswordManager/api/endpoints/password"
	webauthnEndpoints "PasswordManager/api/endpoints/webauthn"
	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/middleware"
	"PasswordManager/api/lib/smtp"
	internalWebauthn "PasswordManager/api/lib/webauthn"
)

func Start(router *gin.Engine) {
	// Config
	db.Connect()
	smtp.Connect()
	internalWebauthn.Register()

	// Endpoints
	authNotRequired := router.Group("/")
	authNotRequired.POST("/api/v1/auth/signup", signup.Post)
	authNotRequired.POST("/api/v1/auth/signin", signin.Post)
	authNotRequired.POST("/api/v1/email/signupChallenge", email.PostSignupEmailChallenge)
	authNotRequired.POST("/api/v1/email/signinChallenge", email.PostSigninEmailChallenge)
	authNotRequired.GET("/api/v1/webauthn/signinChallenge", webauthnEndpoints.GetSigninChallenge)
	authNotRequired.POST("/api/v1/webauthn/signinChallenge", webauthnEndpoints.PostSigninChallenge)
	authNotRequired.GET("/api/v1/emoji", emoji.Get)

	authRequired := router.Group("/")
	authRequired.Use(middleware.AuthMiddleware())
	authRequired.DELETE("/api/v1/auth/signout", signout.Delete)
	authRequired.GET("/api/v1/auth/test", test.Get)
	authRequired.GET("/api/v1/password", password.Get)
	authRequired.POST("/api/v1/password", password.Post)
	authRequired.DELETE("/api/v1/password", password.Delete)
	authRequired.GET("/api/v1/webauthn/register", webauthnEndpoints.GetRegister)
	authRequired.POST("/api/v1/webauthn/register", webauthnEndpoints.PostRegister)
}

func Stop() {
	db.Disconnect()
}
