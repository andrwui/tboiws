package api

import (
	"net/http"
	"strconv"
	"strings"

	"tboiws/models"
	"tboiws/utils"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func NoRootHandler(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusNotFound, "404.html", gin.H{})
}

func ItemsHandler(c *gin.Context, items []models.Item) {
	filteredItems := items

	id := c.DefaultQuery("id", "")
	if id != "" {
		intID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid item ID.",
			})
			return
		}
		filteredItems = utils.Filter(filteredItems, func(item models.Item) bool {
			return item.ID == intID
		})
	}

	name := c.DefaultQuery("name", "")
	if name != "" {
		filteredItems = utils.Filter(filteredItems, func(item models.Item) bool {
			return strings.EqualFold(item.Name, name)
		})
	}

	dlc := c.DefaultQuery("dlc", "")
	if dlc != "" {
		filteredItems = utils.Filter(filteredItems, func(item models.Item) bool {
			return strings.EqualFold(item.DLC, dlc)
		})
	}
	quality := c.DefaultQuery("quality", "")
	if quality != "" {
		filteredItems = utils.Filter(filteredItems, func(item models.Item) bool {
			return item.Quality == quality
		})
	}
	itemType := c.DefaultQuery("type", "")
	if itemType != "" {
		filteredItems = utils.Filter(filteredItems, func(item models.Item) bool {
			return strings.EqualFold(item.Type, itemType)
		})
	}

	if len(filteredItems) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No items found",
		})
		return
	}

	c.JSON(http.StatusOK, filteredItems)
}

func RandomItemHandler(c *gin.Context, items []models.Item) {
	quantityStr := c.DefaultQuery("qty", "1")
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid quantity.",
		})
		return
	}

	filteredItems := items

	dlc := c.DefaultQuery("dlc", "")
	if dlc != "" {
		filteredItems = utils.Filter(filteredItems, func(item models.Item) bool {
			return strings.EqualFold(item.DLC, dlc)
		})
	}
	quality := c.DefaultQuery("quality", "")
	if quality != "" {
		filteredItems = utils.Filter(filteredItems, func(item models.Item) bool {
			return item.Quality == quality
		})
	}
	itemType := c.DefaultQuery("type", "")
	if itemType != "" {
		filteredItems = utils.Filter(filteredItems, func(item models.Item) bool {
			return strings.EqualFold(item.Type, itemType)
		})
	}

	queryItems := utils.GetRandomItems(filteredItems, quantity)
	if queryItems == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No items found",
		})
		return
	}

	c.JSON(http.StatusOK, queryItems)
}
