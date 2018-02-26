package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
)

var (
	msgCh chan []byte
)

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "app.tmpl", gin.H{
		"content": "test",
	})
}

func healthCheckHandler(c *gin.Context) {
	fmt.Fprint(c.Writer, "ok")
}

func wsHandler(c *gin.Context) {
	wsWrite(c.Writer, c.Request)
}

func startServer(port int) error {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// statics
	router.Static("/css", "./static/css")
	router.Static("/img", "./static/img")
	router.Static("/js", "./static/js")

	// templates
	router.LoadHTMLGlob("templates/*")

	// handlers
	router.GET("/", indexHandler)
	router.GET("/_ah/health", healthCheckHandler)
	router.GET("/ws", wsHandler)

	// server
	httpserver := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      handlers.CombinedLoggingHandler(os.Stdout, router),
		ReadTimeout:  40 * time.Second,
		WriteTimeout: 40 * time.Second,
	}

	msgCh = make(chan []byte, 1)
	go subscribe(msgCh)

	// run server
	log.Fatal(httpserver.ListenAndServe())
	return nil
}
