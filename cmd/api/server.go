package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (app *Application) serve() error {
	gin.SetMode(gin.ReleaseMode)
	router := app.routes()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.Port),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Starting server on port %d", app.config.Port)
	return server.ListenAndServe()
}
