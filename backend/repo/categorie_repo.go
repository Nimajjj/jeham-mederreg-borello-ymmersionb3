package repo

import (
	"database/sql"
	"fmt"
)

type CategorieRepo struct {
	db *sql.DB
}


func NewCategorieRepo() *CategorieRepo {
    return &CategorieRepo{ db: DB() }
}


func (cr *CategorieRepo) InsertNewCategorie(categorie_title string) {
    fmt.Print("try InsertNewCategorie -> ", categorie_title)

    var count int
    query := "SELECT COUNT(*) FROM categories WHERE title = ?"
    err := cr.db.QueryRow(query, categorie_title).Scan(&count)
    if err != nil {
      panic(err)
    }

    if count > 0 {
        fmt.Print(" -> FAILED: ALREADY EXISTS\n")
      // categorie already exists
      return 
    }

    // categorie does not exist, insert
    query = "INSERT INTO categories (title) VALUES (?);"

    _, err = cr.db.Exec(query, categorie_title)
    if err != nil {
      panic(err) 
    }
    
    fmt.Print(" -> SUCCESS\n")
}
