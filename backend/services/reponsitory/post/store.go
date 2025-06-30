package post

import "database/sql"

type PostServiceSQL struct {
	db *sql.DB
}

func NewPostServiceSQL(db *sql.DB) *PostServiceSQL {
	return &PostServiceSQL{db: db}
}
