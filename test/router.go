package test

import (
	"github.com/gin-gonic/gin"
)
func Hello() {
	r:=gin.Default()
	_ = r.Run("127.0.0.1:8081")
}
