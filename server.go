package main

import "github.com/gin-gonic/gin"

func createServer() {
	r := gin.Default()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
