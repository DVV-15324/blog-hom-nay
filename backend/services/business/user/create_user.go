package user

import (
	"bloghomnay-project/common"
	entityUser "bloghomnay-project/services/entity/user"
	"context"
	"net/http"
)

func (u *BusinessUser) BzCreateUser(ctx context.Context, cu *entityUser.CreateUserForm) (int, *common.AppError) {
	err_v := cu.Validate()
	if err_v != nil {
		app := common.NewAppError(400, http.StatusText(400), err_v)
		return 0, app
	}
	uid, err := u.userReponsitory.CreateUser(ctx, cu)
	if err != nil {
		app := common.NewAppError(500, http.StatusText(500), err)
		return 0, app
	}
	return uid, nil
}
