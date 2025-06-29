package user

import (
	"bloghomnay-project/common"
	entityUser "bloghomnay-project/services/entity/user"
	"context"
	"net/http"
)

func (u *BusinessUser) BzGetUsersById(ctx context.Context, id int) (*entityUser.Users, *common.AppError) {
	user, err := u.userReponsitory.GetUserById(ctx, id)

	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return nil, app
	}
	return user, nil
}
