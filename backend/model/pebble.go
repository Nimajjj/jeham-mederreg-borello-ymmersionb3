package model

import (
	"fmt"
)

type Pebble struct {
	ID          int      `json:"ID"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Breed       string   `json:"breed"`
	Quantity    int      `json:"quantity"`
	Weight      float64  `json:"weight"`
    Categories  []string `json:"categorie"`
    Photos      []string `json:"image"`
}

func (p Pebble) String() string {
	return fmt.Sprintf(
        "Pebble[ID:%d,Title:%s, Description:%s, Price:%f€, Weight:%fg, QuantityLeft:%d, Categories:%s]",
		p.ID, p.Title, p.Description, p.Price, p.Weight, p.Quantity, p.Categories)
}
