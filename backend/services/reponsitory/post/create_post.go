package post

import (
	"bloghomnay-project/common"
	entity "bloghomnay-project/services/entity/posts"
	"context"
	"database/sql"
	"fmt"
)

// int la uid cua user
func (p *PostServiceSQL) CreatePost(cxt context.Context, post *entity.CreatePost) (int, error) {
	var uid int
	query := `INSERT INTO posts(user_id, category_id, title, content, description, slug)
	OUTPUT INSERTED.id
	VALUES(@user_id, @category_id, @title, @content, @description, @slug);`

	u := p.db.QueryRowContext(cxt, query,
		sql.Named("user_id", post.UserId),
		sql.Named("category_id", post.CategoryId),
		sql.Named("title", post.Title),
		sql.Named("content", post.Content),
		sql.Named("description", post.Description),
		sql.Named("slug", ""), // slug tạm
	)

	err := u.Scan(&uid)
	if err != nil {
		return 0, err
	}

	postIdFake := common.NewUID(uint32(uid), 3)
	slug := fmt.Sprintf("%s-%s", common.GenerateSlug(post.Title), postIdFake.ToBase58())

	queryUpdateSlug := `UPDATE posts SET slug=@slug WHERE id=@id`
	_, err = p.db.ExecContext(cxt, queryUpdateSlug, sql.Named("slug", slug), sql.Named("id", uid))
	if err != nil {
		return 0, err
	}

	return uid, nil
}
