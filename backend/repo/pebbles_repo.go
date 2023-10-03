package repo

import (
    "time"
    "fmt"
    "database/sql"
    
    "prgc/model"
)

type PebbleRepo struct {
	db *sql.DB
}

func NewPebbleRepo() *PebbleRepo {
    return &PebbleRepo{ db: DB() }
}

func (pr *PebbleRepo) InsertNewPebble(pebble *model.Pebble) {
    var count int
    query := "SELECT COUNT(*) FROM pebbles WHERE title = ?"
    err := pr.db.QueryRow(query, pebble.Title).Scan(&count)
    if err != nil {
      panic(err)
    }

    if count > 0 {
      // if pebble already exists in db return
        fmt.Print("try InsertNewPebble -> ", pebble.Title, " -> FAILED ALREADY EXISTS\n")
        return 
    }


    // insert categories
    categories_repo := NewCategorieRepo()
    for _, cat := range pebble.Categories {
        categories_repo.InsertNewCategorie(cat)
    }

    // insert photos


    // finally insert pebble
    query = "INSERT INTO pebbles (Title, Description, Price, Breed, Weight, Quantity, Creation) VALUES (?, ?, ?, ?, ?, ?, ?);"

    _, err = pr.db.Exec(query, pebble.Title, pebble.Description, pebble.Price, pebble.Breed, pebble.Weight, pebble.Quantity, time.Now().Format("01-02-2006"))
    if err != nil {
      panic(err) 
    }

    fmt.Print("try InsertNewPebble -> ", pebble.Title, " -> SUCCESS\n")
}
