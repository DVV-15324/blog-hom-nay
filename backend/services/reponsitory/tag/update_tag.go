package tag

import (
	entity "bloghomnay-project/services/entity/tag"
	"context"
	"database/sql"
	"fmt"
)

func (t *TagServiceSQL) UpdateTag(cxt context.Context, tag *entity.UpdateTagForm, id int) error {
	if tag == nil {
		return fmt.Errorf("tag is nil")
	}
	var name string
	if tag.Name != nil {
		name = *tag.Name
	}

	query := `
		UPDATE tags
		SET name = @name
		WHERE id = @id
	`
	_, err := t.db.ExecContext(cxt, query,
		sql.Named("name", name),
		sql.Named("id", id),
	)
	if err != nil {
		return fmt.Errorf("failed to update tag: %v", err)
	}

	return nil
}
