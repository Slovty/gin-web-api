package main

import (
	"gin-web-api/test"
	"github.com/gin-gonic/gin"
)

func main(){
	test.Hello()

	r:=gin.Default()
	route := InitRouter()
	r.Run(route)
}


