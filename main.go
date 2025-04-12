package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"placeholder-api/generator"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/:size", func(c *gin.Context) {
		generator.Generate(c)
		c.Status(http.StatusOK)
	})
	err := r.Run()
	if err != nil {
		_ = fmt.Errorf("cannot start Gin: %v", err)
	}
}
