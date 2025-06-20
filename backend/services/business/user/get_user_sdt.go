package user

import (
	"bloghomnay-project/common"
	entityUser "bloghomnay-project/services/entity/user"
	"context"
	"net/http"
)

func (u *BusinessUser) BzGetUsersBySDT(ctx context.Context, phone string) (*entityUser.Users, *common.AppError) {
	err_get := entityUser.CheckPhone(phone)
	if err_get != nil {
		app := common.NewAppError(400, http.StatusText(400), err_get)
		return nil, app
	}
	user, err := u.userReponsitory.GetUserBySDT(ctx, phone)
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return nil, app
	}
	return user, nil
}
