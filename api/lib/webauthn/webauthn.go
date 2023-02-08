package webauthn

import (
	"log"
	"os"
	"strings"

	"github.com/go-webauthn/webauthn/webauthn"
)

var Web *webauthn.WebAuthn

func Register() {
	tempWeb, webErr := webauthn.New(&webauthn.Config{
		RPDisplayName: os.Getenv("RP_DISPLAY_NAME"),
		RPID:          os.Getenv("RP_ID"),
		RPOrigins:     strings.Split(os.Getenv("RP_ORIGINS"), ","),
		RPIcon:        os.Getenv("RP_ICON"), // Optional
	})

	if webErr != nil {
		log.Fatal("Failed setting relying party information.")
	}

	Web = tempWeb
}
