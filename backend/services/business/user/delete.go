package user

import (
	"bloghomnay-project/common"
	"context"
	"net/http"
)

func (u *BusinessUser) DeleteUserById(ctx context.Context, id int) *common.AppError {
	err := u.userReponsitory.DeleteUserById(ctx, id)
	if err != nil {
		app := common.NewAppError(500, http.StatusText(500), err)
		return app
	}
	return nil

}
