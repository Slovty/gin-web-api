package router

import (
	"github.com/gin-gonic/gin"
)
func InitRouter() *gin.Engine{
	router := gin.New()
	//apiRouter := router.Group("api")
	//{
	//	apiRouter.GET("/test/index", controllers.IndexApi)
	//}
	//
	//api := router.Group("/api")
	//api.GET("/index", controllers.IndexApi)
	//api.GET("/cookie/set/:userid", controllers.CookieSetExample)

	appv5:= router.Group("appv5")
	{
		appv5.GET("/", GetProjects)
		appv5.GET("/people/:id", GetPerson)
		appv5.POST("/people", CreatePerson)
		appv5.PUT("/people/:id", UpdatePerson)
		appv5.DELETE("/people/:id", DeletePerson)
	}
	//// cookie auth middleware
	//api.Use(auth.Middleware(auth.CookieAuthDriverKey))
	//{
	//	api.GET("/orm", controllers.OrmExample)
	//	api.GET("/store", controllers.StoreExample)
	//	api.GET("/db", controllers.DBExample)
	//	api.GET("/cookie/get", controllers.CookieGetExample)
	//}
	//
	//jwtApi := router.Group("/api")
	//jwtApi.GET("/jwt/set/:userid", controllers.JwtSetExample)
	//
	//// jwt auth middleware
	//jwtApi.Use(auth.Middleware(auth.JwtAuthDriverKey))
	//{
	//	jwtApi.GET("/jwt/get", controllers.JwtGetExample)
	//}
	return router
}
