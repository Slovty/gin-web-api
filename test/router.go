package test

import (
	"github.com/gin-gonic/gin"
)
func Hello() {
	r:=gin.Default()
	r.GET("/", GetProjects)
	r.POST("/people", CreatePerson)
	r.PUT("/people/:id", UpdatePerson)
	r.DELETE("/people/:id", DeletePerson)

	r.Run("127.0.0.1:8081")
}
g