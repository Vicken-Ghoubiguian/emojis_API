package main

//
import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

//
func main() {

	//
	r := gin.Default()

	//
	r.Use(favicon.New("./favicon.ico")) // set favicon middleware

	//
	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	//
	r.GET("/help", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping pong",
		})
	})

	//
	r.GET("/all_emojis", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping pong",
		})
	})

	//
	r.GET("/all_categories", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping pong",
		})
	})

	//
	r.GET("/all_sub_categories", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping pong",
		})
	})

	//
	r.Run() // listen and serve on 0.0.0.0:8080...
}
