package post

import (
	entity "bloghomnay-project/services/entity/posts"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func (t *PostServiceSQL) GetsPostByTagsID(ctx context.Context, tagID []int) ([]entity.Posts, error) {
	if len(tagID) == 0 {
		return []entity.Posts{}, nil
	}

	args := []interface{}{}
	placeholders := []string{}
	for i := 0; i < len(tagID); i++ {
		paramName := fmt.Sprintf("p%d", i)
		placeholders = append(placeholders, "@"+paramName)
		args = append(args, sql.Named(paramName, tagID[i]))
	}

<<<<<<< HEAD
	query := fmt.Sprintf(`SELECT p.id, p.user_id, p.category_id, p.title, p.content,p.created_at,
		p.updated_at FROM posts as p join post_tags as pt on pt.post_id = p.id WHERE pt.tag_id IN (%s)`, strings.Join(placeholders, ", "))
=======
	query := fmt.Sprintf("SELECT p.id, p.user_id, p.category_id, p.title, p.content FROM posts as p join post_tags as pt on pt.post_id = p.id WHERE pt.tag_id IN (%s)", strings.Join(placeholders, ", "))
>>>>>>> 70a38361bb67beb662f248595a90edb388469f20
	rows, err := t.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listPosts []entity.Posts
	for rows.Next() {
		var data entity.Posts
<<<<<<< HEAD
		if err := rows.Scan(&data.Id, &data.UserID, &data.CategoryId, &data.Title, &data.Content, &data.CreatedAt, &data.UpdatedAt); err != nil {
=======
		if err := rows.Scan(&data.Id, &data.UserID, &data.CategoryId, &data.Title, &data.Content); err != nil {
>>>>>>> 70a38361bb67beb662f248595a90edb388469f20
			return nil, err
		}
		listPosts = append(listPosts, data)
	}

	return listPosts, nil
}
