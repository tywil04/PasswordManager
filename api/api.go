package api

import (
	"github.com/gin-gonic/gin"

	"PasswordManager/api/endpoints/auth/signin"
	"PasswordManager/api/endpoints/auth/signout"
	"PasswordManager/api/endpoints/auth/signup"
	"PasswordManager/api/endpoints/auth/test"
	"PasswordManager/api/endpoints/email"
	"PasswordManager/api/endpoints/password"
	"PasswordManager/api/endpoints/totp"
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
	authNotRequired.GET("/api/v1/email/challenge", email.GetChallenge)
	authNotRequired.POST("/api/v1/email/challenge", email.PostChallenge)
	authNotRequired.GET("/api/v1/webauthn/challenge", webauthnEndpoints.GetChallenge)
	authNotRequired.POST("/api/v1/webauthn/challenge", webauthnEndpoints.PostChallenge)
	authNotRequired.POST("/api/v1/totp/challenge", totp.PostChallenge)

	authRequired := router.Group("/")
	authRequired.Use(middleware.AuthMiddleware())
	authRequired.DELETE("/api/v1/auth/signout", signout.Delete)
	authRequired.GET("/api/v1/auth/test", test.Get)
	authRequired.GET("/api/v1/password", password.Get)
	authRequired.POST("/api/v1/password", password.Post)
	authRequired.DELETE("/api/v1/password", password.Delete)
	authRequired.GET("/api/v1/webauthn/register", webauthnEndpoints.GetRegister)
	authRequired.POST("/api/v1/webauthn/register", webauthnEndpoints.PostRegister)
	authRequired.GET("/api/v1/webauthn/credential", webauthnEndpoints.GetCredential)
	authRequired.DELETE("/api/v1/webauthn/credential", webauthnEndpoints.DeleteCredential)
	authRequired.GET("/api/v1/totp/register", totp.GetRegister)
	authRequired.POST("/api/v1/totp/register", totp.PostRegister)
}

func Stop() {
	db.Disconnect()
}
