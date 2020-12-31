package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jiuzhou-zhao/elevation/handler"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())

	// routers
	r.POST("/exec_command", handler.ExecCommand)

	srv := &http.Server{
		Addr:    ":1981",
		Handler: r,
	}
	_ = srv.ListenAndServe()
}
