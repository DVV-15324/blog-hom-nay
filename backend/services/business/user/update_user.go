package user

import (
	"bloghomnay-project/common"
	entityUser "bloghomnay-project/services/entity/user"
	"context"
	"log"

	"net/http"
)

func (u *BusinessUser) BzUpdateUser(ctx context.Context, up *entityUser.UpdateUserForm, id int) *common.AppError {

	if err := up.Validate(); err != nil {
		return common.NewAppError(400, http.StatusText(400), err)
	}

	if err := u.userReponsitory.UpdateUser(ctx, up, id); err != nil {
		return common.NewAppError(500, http.StatusText(500), err)
	}

	user, err := u.userReponsitory.GetUserById(ctx, id)
	if err != nil {
		return common.NewAppError(404, http.StatusText(404), err)
	}

	user.Mask()

	if err := u.bzRedis.SaveProfile(ctx, user); err != nil {
		log.Printf("Lỗi lưu profile vào Redis: %v", err)

	}

	return nil
}
