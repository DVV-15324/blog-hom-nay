package tag

import (
	entity "bloghomnay-project/services/entity/tag"
	"context"
	"database/sql"
)

func (t *TagServiceSQL) CreateTag(ctx context.Context, data *entity.CreateTagForm) error {
	query := "INSERT INTO tags(name) values (@name)"
	_, err := t.db.ExecContext(ctx, query, sql.Named("name", data.Name))
	if err != nil {
		return err
	}
	return nil
}
