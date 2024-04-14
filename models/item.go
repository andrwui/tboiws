package models

type Item struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Quality     string   `json:"quality"`
	Image       string   `json:"image"`
	DLC         string   `json:"dlc"`
	Type        string   `json:"type"`
	Quote       string   `json:"quote"`
	Description string   `json:"description"`
	Pools       []string `json:"pools"`
}
