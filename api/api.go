package api

import (
	"github.com/gin-gonic/gin"

	"PasswordManager/api/endpoints/2fa/email"
	"PasswordManager/api/endpoints/2fa/totp"
	"PasswordManager/api/endpoints/2fa/webauthn"
	"PasswordManager/api/endpoints/auth"
	"PasswordManager/api/endpoints/password"
	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/middleware"
	"PasswordManager/api/lib/smtp"
	webauthnBackend "PasswordManager/api/lib/webauthn"
)

func Start(router *gin.Engine) {
	// Config
	db.Connect()
	smtp.Connect()
	webauthnBackend.Register()

	// Endpoints
	apiV1 := router.Group("/api/v1")

	authGroup := apiV1.Group("/")
	authGroup.Use(middleware.AuthMiddleware())

	unauthGroup := apiV1.Group("/")

	// Auth Endpoints
	unauthGroup.POST("/auth/signup", middleware.ProcessParams(auth.PostSignupInput{}), auth.PostSignup)
	unauthGroup.POST("/auth/signin", middleware.ProcessParams(auth.PostSigninInput{}), auth.PostSignin)
	authGroup.DELETE("/auth/signout", middleware.ProcessParams(auth.DeleteSignoutInput{}), auth.DeleteSignout)
	authGroup.GET("/auth/test", middleware.ProcessParams(auth.GetTestInput{}), auth.GetTest)

	// 2FA Endpoints
	unauthGroup.GET("/2fa/email/challenge", middleware.ProcessParams(email.GetChallengeInput{}), email.GetChallenge)
	unauthGroup.POST("/2fa/email/challenge", middleware.ProcessParams(email.PostChallengeInput{}), email.PostChallenge)

	unauthGroup.GET("/2fa/webauthn/challenge", middleware.ProcessParams(webauthn.GetChallengeInput{}), webauthn.GetChallenge)
	unauthGroup.POST("/2fa/webauthn/challenge", middleware.ProcessParams(webauthn.PostChallengeInput{}), webauthn.PostChallenge)
	authGroup.GET("/2fa/webauthn/register", middleware.ProcessParams(webauthn.GetRegisterInput{}), webauthn.GetRegister)
	authGroup.POST("/2fa/webauthn/register", middleware.ProcessParams(webauthn.PostRegisterInput{}), webauthn.PostRegister)
	authGroup.GET("/2fa/webauthn/credential", middleware.ProcessParams(webauthn.GetCredentialInput{}), webauthn.GetCredential)
	authGroup.DELETE("/2fa/webauthn/credential", middleware.ProcessParams(webauthn.DeleteCredentialInput{}), webauthn.DeleteCredential)

	unauthGroup.POST("/2fa/totp/challenge", middleware.ProcessParams(totp.PostChallengeInput{}), totp.PostChallenge)
	authGroup.GET("/2fa/totp/register", middleware.ProcessParams(totp.GetRegisterInput{}), totp.GetRegister)
	authGroup.POST("/2fa/totp/register", middleware.ProcessParams(totp.PostRegisterInput{}), totp.PostRegister)

	// Other Endpoints
	authGroup.GET("/password", middleware.ProcessParams(password.GetInput{}), password.Get)
	authGroup.POST("/password", middleware.ProcessParams(password.PostInput{}), password.Post)
	authGroup.DELETE("/password", middleware.ProcessParams(password.DeleteInput{}), password.Delete)
}

func Stop() {
	db.Disconnect()
}
