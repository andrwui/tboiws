package utils

import (
	"math/rand/v2"

	"tboiws/models"
)

func GetRandomTrinkets(trinkets []models.Trinket, times int) []models.Trinket {

	var trinketStack []models.Trinket
	trinketsLen := len(trinkets)
	if trinketsLen == 0 {
		return nil
	}

	for i := 0; i < times; i++ {
		isDuplicate := false
		randomTrinket := trinkets[rand.IntN(len(trinkets))]

		for _, item := range trinketStack {
			if item.ID == randomTrinket.ID {
				isDuplicate = true
				i--
				break
			}
		}

		if !isDuplicate {
			trinketStack = append(trinketStack, randomTrinket)
		}
	}
	return trinketStack
}
