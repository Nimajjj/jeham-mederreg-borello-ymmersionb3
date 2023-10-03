package repo

import (
    "database/sql"
)

type PhotoRepo struct {
	db *sql.DB
}


func NewPhotoRepo() *PhotoRepo {
    return &PhotoRepo{ db: DB() }
}
