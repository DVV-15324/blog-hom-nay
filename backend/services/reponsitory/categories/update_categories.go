package categories

import (
	entity "bloghomnay-project/services/entity/categories"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func (c *CategoriesServiceSQL) UpdateCategory(cxt context.Context, categories *entity.UpdateCategory, id int) error {
	var (
		placeholders = []string{}
		args         = []interface{}{}
	)
	if categories == nil {
		return fmt.Errorf("categories is invalid")
	}
	if categories.Name != nil {
		placeholders = append(placeholders, "name=@name")
		args = append(args, sql.Named("name", categories.Name))
	}
	if categories.Description != nil {
		placeholders = append(placeholders, "description=@description")
		args = append(args, sql.Named("description", categories.Description))
	}
	query := fmt.Sprintf("UPDATE categories SET %s Where id = @id", strings.Join(placeholders, ","))
	args = append(args, sql.Named("id", id))
	_, err := c.db.ExecContext(cxt, query,
		args...,
	)
	if err != nil {
		return err
	}
	return nil
}
