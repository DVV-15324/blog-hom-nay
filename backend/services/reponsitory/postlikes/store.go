package postlikes

import "database/sql"

type PostLikeServiceSQL struct {
	db *sql.DB
}

func NewPostLikesServiceSQL(db *sql.DB) *PostLikeServiceSQL {
	return &PostLikeServiceSQL{db: db}
}
