package repo

import (
	"database/sql"
	"fmt"
	"prgc/model"
	"strconv"
)

type CartRepo struct {
	db *sql.DB
}

func NewCartRepo() *CartRepo {
	return &CartRepo{ db: DB() }
}

func (cr *CartRepo) SelectCartFromUser(user_id int) model.Cart {
    var cart model.Cart
    cart.Content = map[*model.Pebble]int{}

    query := "SELECT ID FROM cart WHERE FK_ID_User = ?"
    err := cr.db.QueryRow(query, user_id).Scan(&cart.ID)
    if (err != nil) {
        fmt.Println("ERROR while requesting cart from DB")
        panic(err)
    }

    query = "SELECT ID_Caillou, Quantity FROM pebbles_cart WHERE ID_Basket = ?"
    rows, err := cr.db.Query(query, cart.ID)
    if (err != nil) {
        fmt.Println("ERROR while requesting pebbles_cart from DB")
        panic(err)
    }
    defer rows.Close()

    pIdQt := map[int]int{}
    for rows.Next() {
        var pId int
        var qt int
        err := rows.Scan(&pId, &qt)
        if err != nil {
            panic(err)
        }
        pIdQt[pId] = qt
    }

    pebble_repo := NewPebbleRepo()
    for pId, qt := range pIdQt {
        pebble, err := pebble_repo.GetPebbleById(pId)
        if err != nil {
            panic(err)
        }
        cart.Content[&pebble] = qt
    }

    fmt.Println("SelectCartFromUser SUCCESS -> ", cart)
    return cart
}


func (cr *CartRepo) InsertNewCart(cart model.Cart) {
    var count int
    query := "SELECT COUNT(*) FROM cart WHERE FK_ID_User = ?"
    err := cr.db.QueryRow(query, cart.UserID).Scan(&count)
    if err != nil {
      panic(err)
    }
    if (count != 0) {
        fmt.Println("INSERT FAILED: Cart already exists")
        return
    }

    query = "INSERT INTO Cart (FK_ID_User) VALUES (?)"
    _, err = cr.db.Exec(query, cart.UserID)
    if err != nil {
        fmt.Println("INSERT FAILED: ", query)
        panic(err) 
    }
    
    query = "SELECT ID FROM cart WHERE FK_ID_User = ?"
    cr.db.QueryRow(query, cart.UserID).Scan(&cart.ID)

    fmt.Println(cart)
    for p, qt := range cart.Content {
        query = "INSERT INTO pebbles_cart (ID_Caillou, ID_Basket, Quantity) VALUES (?, ?, ?)"
        _, err = cr.db.Exec(query, p.ID, cart.ID, qt)
        if err != nil {
            fmt.Println("INSERT FAILED: ", query)
            fmt.Println(p.ID, cart.ID, qt)
            panic(err) 
        }
    }
}

func (cr *CartRepo) UserHasCart(user_id int) bool {
    var count int
    query := "SELECT COUNT(*) FROM cart WHERE FK_ID_User = ?"
    err := cr.db.QueryRow(query, user_id).Scan(&count)
    if err != nil {
      panic(err)
    }

    return (count != 0)
}

func (cr *CartRepo) AddPebbleToCart(user_id, pebble_id, quantity int) {
    new_item := true
    cart := cr.SelectCartFromUser(user_id)

    for pebble, _ := range cart.Content {
        if (pebble.ID == pebble_id) {
            new_item = false
            break
        }
    }

    if (new_item) {
        cr.AddItemToCart(cart.ID, pebble_id, quantity)
    } else {
        cr.AddToItemQt(cart.ID, pebble_id, quantity)
    }
}

func (cr *CartRepo) AddItemToCart(cart_id, pebble_id, quantity int) {
    query := "INSERT INTO pebbles_cart (id_caillou, id_basket, quantity) VALUES (?, ?, ?)"
    _, err := cr.db.Exec(query, pebble_id, cart_id, quantity)
    if err != nil {
        panic(err)
    }
}


func (cr *CartRepo) AddToItemQt(cart_id, pebble_id, quantity int) {
    var old_quantity int
    query := "SELECT quantity FROM pebbles_cart"
    query += " WHERE ID_Caillou = " + strconv.Itoa(pebble_id)
    query += " AND ID_Basket = " + strconv.Itoa(cart_id)
    err := cr.db.QueryRow(query).Scan(&old_quantity)
    if err != nil {
      panic(err)
    }

    if (old_quantity + quantity <= 0) {
        query = "DELETE FROM pebbles_cart"
        query += " WHERE ID_Caillou = " + strconv.Itoa(pebble_id)
        query += " AND ID_Basket = " + strconv.Itoa(cart_id)
        _, err = cr.db.Exec(query)
        if err != nil {
          panic(err)
        }
        return
    }

    query = "UPDATE pebbles_cart"
    query += " SET quantity = " + strconv.Itoa(old_quantity + quantity)
    query += " WHERE ID_Caillou = " + strconv.Itoa(pebble_id)
    query += " AND ID_Basket = " + strconv.Itoa(cart_id)
    _, err = cr.db.Exec(query)
    if err != nil {
      panic(err)
    }
}
