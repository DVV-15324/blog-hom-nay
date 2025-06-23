package tag

import "database/sql"

type TagServiceSQL struct {
	db *sql.DB
}

func NewTagServiceSQL(db *sql.DB) *TagServiceSQL {
	return &TagServiceSQL{db: db}
}
