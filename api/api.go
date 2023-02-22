package api

import (
	"github.com/gin-gonic/gin"

	"PasswordManager/api/endpoints/auth"
	"PasswordManager/api/endpoints/email"
	"PasswordManager/api/endpoints/password"
	"PasswordManager/api/endpoints/totp"
	"PasswordManager/api/endpoints/webauthn"
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

	unauthGroup.POST("/auth/signup", middleware.ProcessParams(auth.PostSignupInput{}), auth.PostSignup)
	unauthGroup.POST("/auth/signin", middleware.ProcessParams(auth.PostSigninInput{}), auth.PostSignin)
	authGroup.DELETE("/auth/signout", middleware.ProcessParams(auth.DeleteSignoutInput{}), auth.DeleteSignout)
	authGroup.GET("/auth/test", middleware.ProcessParams(auth.GetTestInput{}), auth.GetTest)

	authGroup.GET("/email/challenge", middleware.ProcessParams(email.GetChallengeInput{}), email.GetChallenge)
	authGroup.POST("/email/challenge", middleware.ProcessParams(email.PostChallengeInput{}), email.PostChallenge)

	unauthGroup.GET("/webauthn/challenge", middleware.ProcessParams(webauthn.GetChallengeInput{}), webauthn.GetChallenge)
	unauthGroup.POST("/webauthn/challenge", middleware.ProcessParams(webauthn.PostChallengeInput{}), webauthn.PostChallenge)
	authGroup.GET("/webauthn/register", middleware.ProcessParams(webauthn.GetRegisterInput{}), webauthn.GetRegister)
	authGroup.POST("/webauthn/register", middleware.ProcessParams(webauthn.PostRegisterInput{}), webauthn.PostRegister)
	authGroup.GET("/webauthn/credential", middleware.ProcessParams(webauthn.GetCredentialInput{}), webauthn.GetCredential)
	authGroup.DELETE("/webauthn/credential", middleware.ProcessParams(webauthn.DeleteCredentialInput{}), webauthn.DeleteCredential)

	unauthGroup.POST("/totp/challenge", middleware.ProcessParams(totp.PostChallengeInput{}), totp.PostChallenge)
	authGroup.GET("/totp/register", middleware.ProcessParams(totp.GetRegisterInput{}), totp.GetRegister)
	authGroup.POST("/totp/register", middleware.ProcessParams(totp.PostRegisterInput{}), totp.PostRegister)

	authGroup.GET("/password", middleware.ProcessParams(password.GetInput{}), password.Get)
	authGroup.POST("/password", middleware.ProcessParams(password.PostInput{}), password.Post)
	authGroup.DELETE("/password", middleware.ProcessParams(password.DeleteInput{}), password.Delete)
}

func Stop() {
	db.Disconnect()
}
