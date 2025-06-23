package categories

import "database/sql"

type CategoriesServiceSQL struct {
	db *sql.DB
}

func NewCategoriesServiceSQL(db *sql.DB) *CategoriesServiceSQL {
	return &CategoriesServiceSQL{db: db}
}
