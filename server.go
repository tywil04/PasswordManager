package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"

	"PasswordManager/api"
	"PasswordManager/api/lib/db"
	"PasswordManager/ent/challenge"
	"PasswordManager/ent/session"
	"PasswordManager/ui"
)

func cleanup() {
	for {
		db.Client.Session.Delete().Where(session.ExpiryLT(time.Now())).Exec(db.Context)
		db.Client.Challenge.Delete().Where(challenge.ExpiryLT(time.Now())).Exec(db.Context)

		time.Sleep(time.Hour * 24)
	}
}

func main() {
	godotenv.Load()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Start API/Frontend
	api.Start(router)
	ui.Start(router)

	// Start cleanup
	go cleanup()

	// Graceful Start/Stop
	fmt.Println("Starting Server...")

	addr := os.Getenv("SERVER_ADDRESS")
	if addr == "" {
		addr = ":8080"
	}
	go router.Run(addr)

	fmt.Println("Server started.")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Stopping Server...")

	// Stop API/Frontend
	api.Stop()
	ui.Stop()
	server.Shutdown(context.Background())

	fmt.Println("Server stopped.")
}
