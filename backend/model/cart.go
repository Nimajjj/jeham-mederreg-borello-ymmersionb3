package model

import (
    "fmt"
)

type Cart struct {
    ID int
    UserID int
    Content map[*Pebble]int
}

type JsonCart struct {
    ID int
    UserID int
    Pebbles []Pebble
}

func (c Cart) String() string {
    return fmt.Sprintf("Cart[\n  ID:%d\n  UserID:%d\n  Pebbles:%s\n]", c.ID, c.UserID, c.Content)
}

func (c Cart) JsonCompatible() JsonCart {
    var json_cart JsonCart
    json_cart.ID = c.ID
    json_cart.UserID = c.UserID
    json_cart.Pebbles = []Pebble{}

    for p, qt := range c.Content {
        new_pebble := p
        new_pebble.Quantity = qt
        json_cart.Pebbles = append(json_cart.Pebbles, *new_pebble)
    }

    return json_cart
}
