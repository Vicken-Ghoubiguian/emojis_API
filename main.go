package main

//
import (
	"github.com/gin-gonic/gin"
	//"github.com/thinkerou/favicon"
)

//
func main() {

	//
	r := gin.Default()

	//
	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	//
	r.GET("/all_emojis", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping pong",
		})
	})

	//
	r.GET("/emoji/:name", func(c *gin.Context) {

		name := c.Params.ByName("name")

		c.JSON(200, gin.H{
			"message": "ping pong" + name,
		})
	})

	//
	r.GET("/all_categories", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping pong",
		})
	})

	//
	r.GET("/search", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping pong",
		})
	})

	//
	r.GET("/search/by_category/:category", func(c *gin.Context) {

		category := c.Params.ByName("category")

		c.JSON(200, gin.H{
			"message": "ping pong" + category,
		})
	})

	//
	r.Run() // listen and serve on 0.0.0.0:8080...
}
