package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"

	"PasswordManager/api"
	"PasswordManager/ui"
)

func main() {
	// Load .env
	godotenv.Load()

	// Set gin mode
	if os.Getenv("ENVIRONMENT") == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	if allowed := os.Getenv("ALLOWED_ORIGINS"); strings.TrimSpace(allowed) != "" {
		router.SetTrustedProxies(strings.Split(allowed, ","))
	}

	// Start Server
	log.Println("Starting server...")

	apiEnabled := !(os.Getenv("DISABLE_API") == "true")
	uiEnabled := !(os.Getenv("DISABLE_UI") == "true")

	if apiEnabled {
		api.Start(router)
		log.Println("API started")
	}

	if uiEnabled {
		ui.Start(router)
		log.Println("UI started")
	}

	addr := os.Getenv("SERVER_ADDRESS")
	if strings.TrimSpace(addr) == "" {
		addr = ":8000"
	}

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go router.Run(addr)

	log.Printf("Successfully started server on %s\n", addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("")
	log.Println("Stopping Server...")

	// Stop API/Frontend
	if apiEnabled {
		api.Stop()
		log.Println("API stopped")
	}

	if uiEnabled {
		ui.Stop()
		log.Println("UI stopped")
	}

	server.Shutdown(context.Background())

	log.Println("Successfully stopped server")
}
