package smtp

import (
	"embed"
	"html/template"
	"io/fs"
)

//go:embed all:templates
var templatesFS embed.FS
var templatesDir, _ = fs.Sub(templatesFS, "templates")

var ChallengeTemplate *template.Template
var SigninNotificationTemplate *template.Template

type ChallengeTemplateData struct {
	Code string
}

type SigninNotificationTemplateData struct {
	Successful bool
}

func LoadTemplates() {
	ChallengeTemplate = template.Must(template.ParseFS(templatesDir, "challenge.html", "base.html"))
	SigninNotificationTemplate = template.Must(template.ParseFS(templatesDir, "signinNotification.html", "base.html"))
}
