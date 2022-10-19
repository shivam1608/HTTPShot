package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	x "github.com/shivam1608/httpshot/api"
	"github.com/shivam1608/httpshot/util"
)

var PORT string
var ENVIROMENT string

func loadConfig() {
	PORT = util.GetEnv("PORT", "3000")
	ENVIROMENT = util.GetEnv("ENVIROMENT", "dev")
}

func main() {
	loadConfig() // Loading Config Data

	if ENVIROMENT == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()                                           // Create a router
	router.Use(static.Serve("/", static.LocalFile("./public", true))) // Serving Homepage
	api := router.Group("/api")                                       // Creating API Group
	x.RegisterRoutes(api)                                             // Registering API Routes
	router.Run(":" + PORT)                                            // Listening to PORT

}
