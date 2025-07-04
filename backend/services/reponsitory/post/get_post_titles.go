package post

import (
	entity "bloghomnay-project/services/entity/posts"
	"context"
	"database/sql"
)

// Tìm kiếm post qua tiêu đề title
func (t *PostServiceSQL) GetPostByTitles(ctx context.Context, keyword string) ([]entity.Posts, error) {
	query := `SELECT 
		p.id, 
		p.user_id, 
		p.category_id, 
		p.description, 
		p.title, 
		p.content,
		COUNT(DISTINCT pl.id) AS [like],
		COUNT(DISTINCT c.id) AS [count_comment],
		p.slug,
		p.created_at,
		p.updated_at
	FROM posts AS p
	LEFT JOIN post_likes AS pl ON pl.post_id = p.id
	LEFT JOIN comments AS c ON c.post_id = p.id
	WHERE 
		p.title LIKE @q OR 
		p.description LIKE @q OR 
		p.content LIKE @q
	GROUP BY 
		p.id, p.user_id, p.category_id, p.description, p.title, p.content, p.slug, p.created_at, p.updated_at
	ORDER BY 
		CASE 
			WHEN p.title LIKE @q THEN 1
			WHEN p.description LIKE @q THEN 2
			WHEN p.content LIKE @q THEN 3
			ELSE 4
		END,
		p.created_at DESC;`

	pattern := "%" + keyword + "%"
	rows, err := t.db.QueryContext(ctx, query, sql.Named("q", pattern))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []entity.Posts
	for rows.Next() {
		var p entity.Posts
		if err := rows.Scan(&p.Id, &p.UserID, &p.CategoryId, &p.Description, &p.Title, &p.Content, &p.Like, &p.CountComment, &p.Slug, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
