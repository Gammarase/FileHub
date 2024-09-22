package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func setupRouter() *gin.Engine {
	r := gin.Default()

  r.LoadHTMLGlob("resources/views/*")

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

  r.GET("/", func(c *gin.Context) {
      c.HTML(200, "main.tmpl", nil)
  })

	return r
}

func main() {
	r := setupRouter()
	r.Run(":9000")
}
