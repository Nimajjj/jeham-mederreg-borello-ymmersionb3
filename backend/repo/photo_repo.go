package repo

import (
    "fmt"
    "database/sql"
)

type PhotoRepo struct {
	db *sql.DB
}


func NewPhotoRepo() *PhotoRepo {
    return &PhotoRepo{ db: DB() }
}

func (pr *PhotoRepo) InsertNewPhoto(photo_path string, pebble_id int) {
    fmt.Print("try InsertNewPhoto -> ", photo_path, "::", pebble_id)

    var count int
    query := "SELECT COUNT(*) FROM photos WHERE FilePath = ?"
    err := pr.db.QueryRow(query, photo_path).Scan(&count)
    if err != nil {
      panic(err)
    }
    if count > 0 {
        fmt.Print(" -> FAILED: ALREADY EXISTS\n")
        return 
    }

    // insert new photo path
    query = "INSERT INTO photos (FilePath) VALUES (?);"

    _, err = pr.db.Exec(query, photo_path)
    if err != nil {
      panic(err) 
    }

    // get photo id
    query = "SELECT ID FROM photos WHERE FilePath = ?"
    var id int
    err = pr.db.QueryRow(query, photo_path).Scan(&id)
    if err != nil {
      panic(err)
    }

    // link photo to pebble
    query = "INSERT INTO pebbles_photos (ID_Pebble, ID_Photo) VALUES (?, ?);"
    _, err = pr.db.Exec(query, pebble_id, id)
    if err != nil {
      panic(err) 
    }

    fmt.Print(" -> SUCCESS\n")
}


func (pr *PhotoRepo) GetPhotosForPebble(pebble_id int) ([]string, error) {
    var photos []string

    query := `SELECT p.filepath 
                FROM pebbles_photos pp
                JOIN photos p ON p.id = pp.id_photo
                WHERE pp.id_pebble = ?`

    rows, err := pr.db.Query(query, pebble_id)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var photo string
        err := rows.Scan(&photo)
        if err != nil {
            return nil, err
        }
        photos = append(photos, photo)
    }

    return photos, nil
}
