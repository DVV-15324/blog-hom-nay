package post

import (
	entity "bloghomnay-project/services/entity/posts"
	"context"
	"database/sql"
)

// Lúc nhấn vào post
func (p *PostServiceSQL) GetPostsByID(ctx context.Context, id int) (*entity.Posts, error) {
	post := entity.Posts{}
	query := `
		SELECT 
		p.id, 
		p.user_id, 
		p.category_id, 
		p.description, 
		p.title, 
		p.content,
		COUNT(DISTINCT pl.id) AS [like],
	    COUNT(DISTINCT c.id) AS [count_comment],
<<<<<<< HEAD
		p.created_at,
		p.updated_at
=======
		p.created_at
>>>>>>> 70a38361bb67beb662f248595a90edb388469f20
		FROM posts AS p
		LEFT JOIN post_likes AS pl ON pl.post_id = p.id
		LEFT JOIN comments AS c ON c.post_id = p.id
		WHERE p.id = @id
		GROUP BY 
<<<<<<< HEAD
			p.id, p.user_id, p.category_id, p.description, p.title, p.content, p.created_at, p.updated_at;`
	err := p.db.QueryRowContext(ctx, query, sql.Named("id", id)).Scan(&post.Id, &post.UserID, &post.CategoryId, &post.Description, &post.Title, &post.Content, &post.Like, &post.CountComment, &post.CreatedAt, &post.UpdatedAt)
=======
			p.id, p.user_id, p.category_id, p.description, p.title, p.content, p.created_at;`
	err := p.db.QueryRowContext(ctx, query, sql.Named("id", id)).Scan(&post.Id, &post.UserID, &post.CategoryId, &post.Description, &post.Title, &post.Content, &post.Like, &post.CountComment, &post.CreatedAt)
>>>>>>> 70a38361bb67beb662f248595a90edb388469f20
	if err != nil {
		return nil, err
	}
	return &post, nil
}
