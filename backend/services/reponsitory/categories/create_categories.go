package categories

import (
	entity "bloghomnay-project/services/entity/categories"
	"context"
	"database/sql"
)

func (c *CategoriesServiceSQL) CreateCategories(cxt context.Context, categories *entity.CreateCategory) error {
	query := `INSERT INTO categories(description, name, img)
			values(@description, @name, @img);`
	_, err := c.db.ExecContext(cxt, query,
		sql.Named("description", categories.Description),
		sql.Named("name", categories.Name),
		sql.Named("img", categories.Img),
	)
	if err != nil {
		return err
	}
	return nil
}
