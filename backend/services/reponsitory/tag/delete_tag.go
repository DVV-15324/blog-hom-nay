package tag

import (
	"context"
	"database/sql"
)

func (u *TagServiceSQL) DeleteTagById(cxt context.Context, id int) error {
	query := "DELETE * from tags Where id = @id"
	_, err := u.db.ExecContext(cxt, query, sql.Named("id", id))
	if err != nil {
		return err
	}
	return nil
}
