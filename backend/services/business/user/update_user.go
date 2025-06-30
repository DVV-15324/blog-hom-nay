package user

import (
	"bloghomnay-project/common"
	entityUser "bloghomnay-project/services/entity/user"
	"context"
	"net/http"
)

func (u *BusinessUser) BzUpdateUser(ctx context.Context, up *entityUser.UpdateUserForm, id int) *common.AppError {
	err_up := up.Validate()
	if err_up != nil {
		app := common.NewAppError(400, http.StatusText(400), err_up)
		return app
	}
	err := u.userReponsitory.UpdateUser(ctx, up, id)
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err_up)
		return app
	}
	return nil
}
