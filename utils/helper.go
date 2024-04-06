package utils

import (
	"encoding/json"
	"io"
	"os"
)

type FilterFunc[T any] func(T) bool

func ParseJson(path string, reference interface{}) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, reference)
	if err != nil {
		return err
	}
	return nil
}

func Filter[T any](items []T, filterFunc FilterFunc[T]) []T {
	var filteredItems []T
	for _, item := range items {
		if filterFunc(item) {
			filteredItems = append(filteredItems, item)
		}
	}
	return filteredItems
}
