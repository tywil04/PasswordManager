package main

import (
	"context"
	"fmt"
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
	fmt.Println("Starting Server...")

	apiEnabled := !(os.Getenv("DISABLE_API") == "true")
	uiEnabled := !(os.Getenv("DISABLE_UI") == "true")

	if apiEnabled {
		api.Start(router)
		fmt.Println("Started API...")
	}

	if uiEnabled {
		ui.Start(router)
		fmt.Println("Started UI...")
	}

	addr := os.Getenv("SERVER_ADDRESS")
	if strings.TrimSpace(addr) == "" {
		addr = ":8080"
	}

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go router.Run(addr)

	fmt.Println("Server started.")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Stopping Server...")

	// Stop API/Frontend
	if apiEnabled {
		api.Stop()
		fmt.Println("Stopped API...")
	}

	if uiEnabled {
		ui.Stop()
		fmt.Println("Stopped UI...")
	}

	server.Shutdown(context.Background())

	fmt.Println("Server stopped.")
}
