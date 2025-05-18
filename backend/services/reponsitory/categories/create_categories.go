package categories

import (
	entity "bloghomnay-project/services/entity/categories"
	"context"
	"database/sql"
)

func (c *CategoriesServiceSQL) CreateCategories(cxt context.Context, categories *entity.CreateCategory) error {
	query := `INSERT INTO categories(description, name)
			values(@description, @name);`
	_, err := c.db.ExecContext(cxt, query,
		sql.Named("description", categories.Description),
		sql.Named("name", categories.Name),
	)
	if err != nil {
		return err
	}
	return nil
}
