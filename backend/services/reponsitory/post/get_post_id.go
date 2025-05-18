package post

import (
	entity "bloghomnay-project/services/entity/posts"
	"context"
	"database/sql"
)

// Lúc nhấn vào post
func (p *PostServiceSQL) GetPostsByID(ctx context.Context, id int) (*entity.Posts, error) {
	post := entity.Posts{}
	query := "SELECT id, user_id, category_id, title, content FROM posts where id = @id"
	err := p.db.QueryRowContext(ctx, query, sql.Named("id", id)).Scan(&post.Id, &post.UserID, &post.CategoryId, &post.Title, &post.Content)
	if err != nil {
		return nil, err
	}
	return &post, nil
}
