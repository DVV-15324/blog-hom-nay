package post

import (
	entity "bloghomnay-project/services/entity/posts"
	"context"
	"database/sql"
)

// int la uid cua user
func (p *PostServiceSQL) CreatePost(cxt context.Context, post *entity.CreatePost) (int, error) {
	var uid int
<<<<<<< HEAD
	query := `INSERT INTO posts(user_id, category_id, title, content, description)
	OUTPUT INSERTED.id
			values(@user_id, @category_id, @title, @content, @description);`
=======
	query := `INSERT INTO posts(user_id, category_id, title, content)
	OUTPUT INSERTED.id
			values(@user_id, @category_id, @title, @content);`
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	u := p.db.QueryRowContext(cxt, query,
		sql.Named("user_id", post.UserId),
		sql.Named("category_id", post.CategoryId),
		sql.Named("title", post.Title),
		sql.Named("content", post.Content),
<<<<<<< HEAD
		sql.Named("description", post.Description),
=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	)
	err := u.Scan(&uid)
	if err != nil {
		return 0, err
	}
	return uid, nil
}
