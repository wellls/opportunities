package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Initialize Router
	r := gin.Default()

	//Initialize Routes
	InitializeRoutes(r)

	//Run server on 0.0.0.0:8080
	r.Run()
}
