package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"log/slog"
	"net/http"
	"os"
	"proxy-server/internal/handler"
)

// @title HTTP Proxy Server API
// @version 1.0
// @description This is a server to proxy HTTP requests.

func main() {
	router := gin.Default()
	router.POST("/proxy", handler.HandleRequest)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	slog.Info("Starting server", slog.String("port", port))
	log.Fatal(http.ListenAndServe(":"+port, router))
}
