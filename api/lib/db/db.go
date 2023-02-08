package db

import (
	"context"
	"log"
	"os"

	"PasswordManager/ent"
)

var Client *ent.Client
var Context context.Context

func Connect() {
	client, clientErr := ent.Open("sqlite3", os.Getenv("DB_PATH")+"?cache=shared&_fk=1")
	if clientErr != nil {
		log.Fatalf("Failed opening connection to sqlite: %v.", clientErr)
	}

	migrationErr := client.Schema.Create(context.Background())
	if migrationErr != nil {
		log.Fatalf("Failed creating schema resources: %v.", migrationErr)
	}

	Client = client
	Context = context.Background()
}

func Disconnect() {
	Client.Close()
}
