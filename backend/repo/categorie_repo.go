package repo

import (
	"database/sql"
	"fmt"
	"prgc/model"
)

type CategorieRepo struct {
	db *sql.DB
}


func NewCategorieRepo() *CategorieRepo {
    return &CategorieRepo{ db: DB() }
}


func (cr *CategorieRepo) InsertNewCategorie(categorie_title string, pebble_id int) {
    fmt.Print("try InsertNewCategorie -> ", categorie_title, "::", pebble_id)

    var count int
    query := "SELECT COUNT(*) FROM categories WHERE title = ?"
    err := cr.db.QueryRow(query, categorie_title).Scan(&count)
    if err != nil {
      panic(err)
    }

    if count == 0 {
        // categorie does not exist, insert
        query = "INSERT INTO categories (title) VALUES (?);"
        _, err = cr.db.Exec(query, categorie_title)
        if err != nil {
          panic(err) 
        }
        fmt.Print(" -> INSERT SUCCESS")
    } else {
        fmt.Print(" -> INSERT FAILED: ALREADY EXISTS")
    }

    // get categorie id
    query = "SELECT ID FROM categories WHERE Title = ?"
    var id int
    err = cr.db.QueryRow(query, categorie_title).Scan(&id)
    if err != nil {
      panic(err)
    }

    // link categoreie to pebble
    query = "INSERT INTO pebbles_categories (ID_Pebble, ID_Categorie) VALUES (?, ?);"
    _, err = cr.db.Exec(query, pebble_id, id)
    if err != nil {
      panic(err) 
    }
    
    fmt.Print(" -> LINK PEBBLE SUCCESS\n")
}

func (cr *CategorieRepo) GetCategoriesForPebble(pebble_id int) ([]string, error) {
    var categories []string
    query := `SELECT c.title
                FROM pebbles_categories pc
                JOIN categories c ON c.id = pc.id_categorie 
                WHERE pc.id_pebble = ?`

    rows, err := cr.db.Query(query, pebble_id)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var category string
        err := rows.Scan(&category)
        if err != nil {
            return nil, err
        }
        categories = append(categories, category)
    }
    return categories, nil
}


func (cr *CategorieRepo) GetAllCategories() ([]model.Categorie, error) {
    var categories []model.Categorie

    query := "SELECT * FROM categories"
    rows, err := cr.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var cat model.Categorie
        err := rows.Scan(&cat.ID, &cat.Title)
        if err != nil {
            return nil, err
        }
        categories = append(categories, cat)
    }

    return categories, nil
}
