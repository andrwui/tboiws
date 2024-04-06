package api

import (
	"net/http"
	"strconv"
	"strings"

	"tboiws/models"
	"tboiws/utils"

	"github.com/gin-gonic/gin"
)

func TrinketHandler(c *gin.Context, trinkets []models.Trinket) {
	filteredItems := trinkets

	id := c.DefaultQuery("id", "")
	if id != "" {
		intID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid trinket ID.",
			})
			return
		}
		filteredItems = utils.Filter(filteredItems, func(trinket models.Trinket) bool {
			return trinket.ID == intID
		})
	}

	name := c.DefaultQuery("name", "")
	if name != "" {
		filteredItems = utils.Filter(filteredItems, func(trinket models.Trinket) bool {
			return strings.EqualFold(trinket.Name, name)
		})
	}

	dlc := c.DefaultQuery("dlc", "")
	if dlc != "" {
		filteredItems = utils.Filter(filteredItems, func(trinket models.Trinket) bool {
			return strings.EqualFold(trinket.DLC, dlc)
		})
	}

	if len(filteredItems) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No trinkets found",
		})
		return
	}

	c.JSON(http.StatusOK, filteredItems)
}

func RandomTrinketHandler(c *gin.Context, trinkets []models.Trinket) {
	quantityStr := c.DefaultQuery("qty", "1")
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid quantity.",
		})
		return
	}

	filteredTrinkets := trinkets

	dlc := c.DefaultQuery("dlc", "")
	if dlc != "" {
		filteredTrinkets = utils.Filter(filteredTrinkets, func(item models.Trinket) bool {
			return strings.EqualFold(item.DLC, dlc)
		})
	}
	queryItems := utils.GetRandomTrinkets(filteredTrinkets, quantity)
	if queryItems == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No items found",
		})
		return
	}

	c.JSON(http.StatusOK, queryItems)
}
