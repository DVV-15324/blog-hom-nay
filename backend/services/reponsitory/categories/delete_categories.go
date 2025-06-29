package categories

import (
	"context"
	"database/sql"
)

func (c *CategoriesServiceSQL) DeleteCategories(cxt context.Context, id int) error {
	query := "DELETE categories Where id = @id"
	_, err := c.db.ExecContext(cxt, query, sql.Named("id", id))
	if err != nil {
		return err
	}
	return nil
}
