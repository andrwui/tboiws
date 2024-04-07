package main

import (
	"log"
	"os"

	"path/filepath"

	"github.com/gin-gonic/gin"

	"tboiws/api"
	"tboiws/models"
	"tboiws/utils"
)

func main() {

	wd, _ := os.Getwd()

	itemsJsonPath := filepath.Join(wd, "json/items.json")
	trinketsJsonPath := filepath.Join(wd, "json/trinkets.json")
	homeHTMLPath := filepath.Join(wd, "static/index.html")
	notFoundHTMLPath := filepath.Join(wd, "static/404.html")

	router := gin.Default()
	router.LoadHTMLFiles(homeHTMLPath, notFoundHTMLPath)

	var items []models.Item
	var trinkets []models.Trinket

	err := utils.ParseJson(itemsJsonPath, &items)
	if err != nil {
		log.Panicf("error: %s", err)
	}
	err2 := utils.ParseJson(trinketsJsonPath, &trinkets)
	if err2 != nil {
		log.Panicf("error: %s", err2)
	}

	router.GET("/", api.RootHandler)

	router.NoRoute(api.NoRootHandler)

	router.GET("/items", func(c *gin.Context) {
		api.ItemsHandler(c, items)
	})

	router.GET("/items/random", func(c *gin.Context) {
		api.RandomItemHandler(c, items)
	})

	router.GET("/trinkets", func(c *gin.Context) {
		api.TrinketHandler(c, trinkets)
	})

	router.GET("/trinkets/random", func(c *gin.Context) {
		api.RandomTrinketHandler(c, trinkets)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
