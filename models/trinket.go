package models

type Trinket struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	DLC         string `json:"dlc"`
	Quote       string `json:"quote"`
	Description string `json:"description"`
}
