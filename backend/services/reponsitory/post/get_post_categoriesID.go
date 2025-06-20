package post

import (
	entity "bloghomnay-project/services/entity/posts"
	"context"
	"database/sql"
)

// Lọc post qua Categories
func (p *PostServiceSQL) GetPostByCategories(ctx context.Context, categoryID int) ([]entity.Posts, error) {
<<<<<<< HEAD
	query := `SELECT 
		p.id, 
		p.user_id, 
		p.category_id, 
		p.description, 
		p.title, 
		p.content,
		COUNT(pl.id) AS [like],
		COUNT(c.id) AS [count_comment]
		FROM posts AS p
		LEFT JOIN post_likes AS pl ON pl.post_id = p.id
		LEFT JOIN comments AS c ON c.post_id = p.id
		WHERE p.category_id = @category_id
		GROUP BY 
			p.id, p.user_id, p.category_id, p.description, p.title, p.content;`
	rows, err := p.db.QueryContext(ctx, query, sql.Named("category_id", categoryID))

=======
	query := "SELECT id, user_id, category_id, title, content FROM posts where category_id = @category_id"
	rows, err := p.db.QueryContext(ctx, query, sql.Named("category_id", categoryID))
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listPosts []entity.Posts
	for rows.Next() {
		var data entity.Posts
<<<<<<< HEAD
		if err := rows.Scan(&data.Id, &data.UserID, &data.CategoryId, &data.Description, &data.Title, &data.Content, &data.Like, &data.CountComment); err != nil {
=======
		if err := rows.Scan(&data.Id, &data.UserID, &data.CategoryId, &data.Title, &data.Content); err != nil {
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
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
