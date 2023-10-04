package repo

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"prgc/model"
)

const DEFAULT = 0b0000
const CATEGORIES = 0b0001
const ORDER = 0b0010
const KEYWORD = 0b0100

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

    // insert pebble
    query = "INSERT INTO pebbles (Title, Description, Price, Breed, Weight, Quantity, Creation) VALUES (?, ?, ?, ?, ?, ?, ?);"

    _, err = pr.db.Exec(query, pebble.Title, pebble.Description, pebble.Price, pebble.Breed, pebble.Weight, pebble.Quantity, time.Now().Format("01-02-2006"))
    if err != nil {
      panic(err) 
    }

    // get pebble id
    query = "SELECT ID FROM pebbles WHERE title = ?"
    var id int
    err = pr.db.QueryRow(query, pebble.Title).Scan(&id)
    if err != nil {
      panic(err)
    }

    // insert categories
    categories_repo := NewCategorieRepo()
    for _, cat := range pebble.Categories {
        categories_repo.InsertNewCategorie(cat, id)
    }
    
    // insert photos
    photo_repo := NewPhotoRepo()
    for _, photo := range pebble.Photos {
        photo_repo.InsertNewPhoto(photo, id)
    }

    fmt.Print("try InsertNewPebble -> ", pebble.Title, " -> SUCCESS\n")
}


func (pr *PebbleRepo) GetPebbleById(id int) (model.Pebble, error) {
    var pebble model.Pebble

    query := "SELECT Title, Description, Price, Breed, Weight, Quantity FROM pebbles WHERE ID = ?"

    err := pr.db.QueryRow(query, id).Scan(&pebble.Title, &pebble.Description, &pebble.Price, &pebble.Breed, &pebble.Weight, &pebble.Quantity)
    if err != nil {
        fmt.Println("ERROR while requesting pebble from DB")
        panic(err)
    }

    // Get categories
    categoriesRepo := NewCategorieRepo()
    pebble.Categories, err = categoriesRepo.GetCategoriesForPebble(id)
    if err != nil {
        fmt.Println("ERROR while requesting categories from DB")
        panic(err)
    }

    // Get photos  
    photoRepo := NewPhotoRepo()
    pebble.Photos, err = photoRepo.GetPhotosForPebble(id)
    if err != nil {
        fmt.Println("ERROR while requesting photos from DB")
        panic(err)
    }

    fmt.Println("GetPebbleById SUCCESS -> ", pebble)
    return pebble, nil
}


type Search struct {
    Flags int64
    Categories []string
    Order string
    Keywords string
}


func (pr *PebbleRepo) SearchPebbles(search Search) ([]model.Pebble, error) {
    var pebbles []model.Pebble

    // prepare query
    query := `SELECT p.*
            FROM pebbles p
            INNER JOIN pebbles_categories pc ON p.id = pc.id_pebble  
            INNER JOIN categories c ON pc.id_categorie = c.id`

    // apply categories filters
    var defer_search_term string
    if (search.Flags & CATEGORIES != 0) {
        count := len(search.Categories)
        if (count == 1) {
            query += " WHERE c.title='"
            query += search.Categories[0]
            query += "'"
        } else {
            for i, category := range search.Categories {
                if (i == 0) {
                    query += " WHERE "
                }

                if (i < len(search.Categories) && i != 0) {
                    query += " OR "
                }

                query += "c.title='"
                query += category
                query += "'"
            }

            // defer this
            defer_search_term += " HAVING COUNT(DISTINCT c.Title) = "
            defer_search_term += strconv.Itoa(len(search.Categories))
        }
    }

    // apply search by keywords
    if (search.Flags & KEYWORD != 0) {
        for i, keyword := range strings.Split(search.Keywords, " "){
            if (i == 0) {
                if (strings.Contains(query, "WHERE")) {
                    query += " OR "
                } else {
                    query += " WHERE "
                }
            }

            if (i < len(search.Keywords) && i != 0) {
                query += " OR "
            }

            query += " p.title LIKE '%"
            query += keyword
            query += "%'"
        }

        query += " OR "

        for i, keyword := range strings.Split(search.Keywords, " "){
            if (i < len(search.Keywords) && i != 0) {
                query += " OR "
            }

            query += " p.description LIKE '%"
            query += keyword
            query += "%'"
        }
    }


    query += " GROUP BY p.id"
    query += defer_search_term

    // apply order (last one)
    if (search.Flags & ORDER != 0) {
        switch search.Order {
        case "price_desc": 
            query += " ORDER BY p.price DESC"
        case "price_asc": 
            query += " ORDER BY p.price ASC"
        default: 
            fmt.Println("Incorect ORDER flag")
        }
    }

    fmt.Println("\n----\n", query, "\n----\n")
    

    rows, err := pr.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var pebble model.Pebble
        
        err = rows.Scan(&pebble.ID, &pebble.Title, &pebble.Description, &pebble.Price, &pebble.Creation, &pebble.Quantity, &pebble.Breed, &pebble.Weight)
        if err != nil {
            panic(err)
            return nil, err
        }

        // Get categories
        categoriesRepo := NewCategorieRepo()
        pebble.Categories, err = categoriesRepo.GetCategoriesForPebble(pebble.ID)

        // Get photos
        photoRepo := NewPhotoRepo()
        pebble.Photos, err = photoRepo.GetPhotosForPebble(pebble.ID)

        pebbles = append(pebbles, pebble)
    }

    // run query

    fmt.Println("SearchPebbles SUCCESS ->")
    fmt.Println(pebbles)
    return pebbles, nil
}
