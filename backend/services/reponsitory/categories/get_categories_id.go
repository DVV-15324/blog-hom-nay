package categories

import (
	entity "bloghomnay-project/services/entity/categories"
	"context"
	"database/sql"
)

// Lúc nhấn vào post
func (p *CategoriesServiceSQL) GetcategoriesByID(ctx context.Context, id int) (*entity.Categories, error) {
	post := entity.Categories{}
	query := "SELECT id, name, description, tag_id FROM posts where id = @id"
	err := p.db.QueryRowContext(ctx, query, sql.Named("id", id)).Scan(&post.Id, &post.Name, &post.Description)
	if err != nil {
		return nil, err
	}
	return &post, nil
}
