package img

import "database/sql"

type ImgServiceSQL struct {
	db *sql.DB
}

func NewImgServiceSQL(db *sql.DB) *ImgServiceSQL {
	return &ImgServiceSQL{db: db}
}
