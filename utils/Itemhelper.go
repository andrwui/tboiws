package utils

import (
	"math/rand/v2"

	"tboiws/models"
)

func GetRandomItems(items []models.Item, times int) []models.Item {

	var itemStack []models.Item
	itemsLen := len(items)
	if itemsLen == 0 {
		return nil
	}

	for i := 0; i < times; i++ {
		isDuplicate := false
		randomItem := items[rand.IntN(len(items))]

		for _, item := range itemStack {
			if item.ID == randomItem.ID {
				isDuplicate = true
				i--
				break
			}
		}

		if !isDuplicate {
			itemStack = append(itemStack, randomItem)
		}
	}
	return itemStack
}
