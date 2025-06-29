package post

import (
	entity "bloghomnay-project/services/entity/posts"
	"context"
	"database/sql"
)

// Lọc post qua Categories
func (p *PostServiceSQL) GetPostByCategories(ctx context.Context, categoryID int) ([]entity.Posts, error) {
	query := `SELECT 
		p.id, 
		p.user_id, 
		p.category_id, 
		p.description, 
		p.title, 
		p.content,
		COUNT(pl.id) AS [like],
<<<<<<< HEAD
		COUNT(c.id) AS [count_comment],
		p.created_at,
		p.updated_at
=======
		COUNT(c.id) AS [count_comment]
>>>>>>> 70a38361bb67beb662f248595a90edb388469f20
		FROM posts AS p
		LEFT JOIN post_likes AS pl ON pl.post_id = p.id
		LEFT JOIN comments AS c ON c.post_id = p.id
		WHERE p.category_id = @category_id
		GROUP BY 
<<<<<<< HEAD
			p.id, p.user_id, p.category_id, p.description, p.title, p.content, p.created_at, p.updated_at;`
=======
			p.id, p.user_id, p.category_id, p.description, p.title, p.content;`
>>>>>>> 70a38361bb67beb662f248595a90edb388469f20
	rows, err := p.db.QueryContext(ctx, query, sql.Named("category_id", categoryID))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listPosts []entity.Posts
	for rows.Next() {
		var data entity.Posts
<<<<<<< HEAD
		if err := rows.Scan(&data.Id, &data.UserID, &data.CategoryId, &data.Description, &data.Title, &data.Content, &data.Like, &data.CountComment, &data.CreatedAt, &data.UpdatedAt); err != nil {
=======
		if err := rows.Scan(&data.Id, &data.UserID, &data.CategoryId, &data.Description, &data.Title, &data.Content, &data.Like, &data.CountComment); err != nil {
>>>>>>> 70a38361bb67beb662f248595a90edb388469f20
			return nil, err
		}
		listPosts = append(listPosts, data)
	}

	// Check lỗi sau khi duyệt xong (ví dụ: lỗi mạng trong lúc đọc)
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listPosts, nil
}
