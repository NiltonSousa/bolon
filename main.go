package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	var r = gin.Default()

	err := r.Run("localhost:8000")

	if err != nil {
		panic(err)
	}
}
