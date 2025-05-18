package comment

import "database/sql"

type CommentServiceSQL struct {
	db *sql.DB
}

func NewCommentServiceSQL(db *sql.DB) *CommentServiceSQL {
	return &CommentServiceSQL{db: db}
}
