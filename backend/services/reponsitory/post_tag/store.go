package posttag

import "database/sql"

type PostTagServiceSQL struct {
	db *sql.DB
}

func NewPostTagServiceSQL(db *sql.DB) *PostTagServiceSQL {
	return &PostTagServiceSQL{db: db}
}
