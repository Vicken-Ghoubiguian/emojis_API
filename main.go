package main

//
import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/thinkerou/favicon"
	"os"
)

//
func main() {

	/********************* API configuration *********************/

	//
	var log_logrus = logrus.New()

	//
	log_logrus.Out = os.Stdout

	//
	r := gin.Default()

	//
	r.Use(favicon.New("./favicon.ico")) // set favicon middleware

	/********************* Presentation API route *********************/

	//
	r.GET("", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"title":        "emojis_GoAPI",
			"description":  "API written in Go to get, return, treat and manage all existing emojis",
			"presentation": "Web API to get, return, treat and manage all existing emojis and to make statistical calculations on them",
			"author":       "Vicken Ghoubiguian",
			"version":      1.0,
			"logo":         "https://raw.githubusercontent.com/Vicken-Ghoubiguian/emojis_GoAPI/main/favicon.ico",
			"github":       "https://github.com/Vicken-Ghoubiguian/emojis_GoAPI",
			"docker":       "",
			"manual":       c.Request.Host + c.Request.URL.Path + "manual",
			"uses":         c.Request.Host + c.Request.URL.Path + "manual/uses",
			"help":         c.Request.Host + c.Request.URL.Path + "manual/help",
		})
	})

	/********************** Manual API routes **********************/

	//
	r.GET("/manual", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Manual",
		})
	})

	//
	r.GET("/manual/help", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Help... Manual...",
		})
	})

	//
	r.GET("/manual/uses", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Uses... Manual...",
		})
	})

	/********************** All objects API routes **********************/

	//
	r.GET("/all_emojis", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "List all emojis...",
		})
	})

	//
	r.GET("/all_categories", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "List all categories' name and descprition...",
		})
	})

	//
	r.GET("/all_sub_categories", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "List all categories' name, descprition and category to which it belongs...",
		})
	})

	//
	r.Run() // listen and serve on 0.0.0.0:8080...
}
