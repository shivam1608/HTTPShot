package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shivam1608/httpshot/handler"
)

// RegisterRoutes : Registers all the routes on /API
func RegisterRoutes(api *gin.RouterGroup) {
	api.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"staus": "Server API Operational 😁",
		})
	})

	api.GET("/capture", func(ctx *gin.Context) {
		url := ctx.DefaultQuery("url", "https://http.cat/400")
		width := ctx.DefaultQuery("width", "1280")
		height := ctx.DefaultQuery("height", "720")
		fullscreen := ctx.DefaultQuery("fullscreen", "false")

		_, err := http.Get(url)
		if err != nil {
			url = "https://http.cat/400"
		}
		ctx.Data(200, "image/png", handler.CaptureHTTPShot(url, width, height, fullscreen))
	})

}
