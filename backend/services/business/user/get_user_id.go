package user

import (
	"bloghomnay-project/common"
	entityUser "bloghomnay-project/services/entity/user"
	"context"

	"net/http"
)

func (u *BusinessUser) BzGetUsersById(ctx context.Context, id int) (*entityUser.Users, *common.AppError) {
	s := common.NewUID(uint32(id), 1)
	userMask := s.ToBase58()

	userRedis, err := u.bzRedis.GetProfileEntity(ctx, userMask)
	if err == nil && userRedis != nil {
		userRedis.Id = int(common.DecodeFromBase58(userRedis.FakeId).LocalID)
		return userRedis, nil
	}

	user, err := u.userReponsitory.GetUserById(ctx, id)
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return nil, app
	}

	user.Mask()

	if err := u.bzRedis.SaveProfile(ctx, user); err != nil {

	}

	return user, nil
}
