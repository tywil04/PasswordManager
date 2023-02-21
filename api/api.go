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
	authNotRequired.POST("/api/v1/auth/signup", middleware.ProcessParams(signup.PostInput{}), signup.Post)
	authNotRequired.POST("/api/v1/auth/signin", middleware.ProcessParams(signin.PostInput{}), signin.Post)
	authNotRequired.GET("/api/v1/email/challenge", middleware.ProcessParams(email.GetChallengeInput{}), email.GetChallenge)
	authNotRequired.POST("/api/v1/email/challenge", middleware.ProcessParams(email.PostChallengeInput{}), email.PostChallenge)
	authNotRequired.GET("/api/v1/webauthn/challenge", middleware.ProcessParams(webauthnEndpoints.GetChallengeInput{}), webauthnEndpoints.GetChallenge)
	authNotRequired.POST("/api/v1/webauthn/challenge", middleware.ProcessParams(webauthnEndpoints.PostChallengeInput{}), webauthnEndpoints.PostChallenge)
	authNotRequired.POST("/api/v1/totp/challenge", middleware.ProcessParams(totp.PostChallengeInput{}), totp.PostChallenge)

	authRequired := router.Group("/")
	authRequired.Use(middleware.AuthMiddleware())
	authRequired.DELETE("/api/v1/auth/signout", middleware.ProcessParams(signout.DeleteInput{}), signout.Delete)
	authRequired.GET("/api/v1/auth/test", middleware.ProcessParams(test.GetInput{}), test.Get)
	authRequired.GET("/api/v1/password", middleware.ProcessParams(password.GetInput{}), password.Get)
	authRequired.POST("/api/v1/password", middleware.ProcessParams(password.PostInput{}), password.Post)
	authRequired.DELETE("/api/v1/password", middleware.ProcessParams(password.DeleteInput{}), password.Delete)
	authRequired.GET("/api/v1/webauthn/register", middleware.ProcessParams(webauthnEndpoints.GetRegisterInput{}), webauthnEndpoints.GetRegister)
	authRequired.POST("/api/v1/webauthn/register", middleware.ProcessParams(webauthnEndpoints.PostRegisterInput{}), webauthnEndpoints.PostRegister)
	authRequired.GET("/api/v1/webauthn/credential", middleware.ProcessParams(webauthnEndpoints.GetCredentialInput{}), webauthnEndpoints.GetCredential)
	authRequired.DELETE("/api/v1/webauthn/credential", middleware.ProcessParams(webauthnEndpoints.DeleteCredentialInput{}), webauthnEndpoints.DeleteCredential)
	authRequired.GET("/api/v1/totp/register", middleware.ProcessParams(totp.GetRegisterInput{}), totp.GetRegister)
	authRequired.POST("/api/v1/totp/register", middleware.ProcessParams(totp.PostRegisterInput{}), totp.PostRegister)
}

func Stop() {
	db.Disconnect()
}
