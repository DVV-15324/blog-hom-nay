package user

import (
	"bloghomnay-project/common"
	entityUser "bloghomnay-project/services/entity/user"
	"context"
	"net/http"
)

func (u *BusinessUser) BzGetUsersById(ctx context.Context, id int) (*entityUser.Users, *common.AppError) {
	s := common.NewUID(uint32(id), 1) // Tạo FakeId
	userMask := s.ToBase58()

	// Ưu tiên lấy từ Redis
	userRedis, err := u.bzRedis.GetProfileEntity(ctx, userMask)
	if err == nil && userRedis != nil {
		return userRedis, nil
	}

	// Nếu không có trong Redis, fallback sang DB
	user, err := u.userReponsitory.GetUserById(ctx, id)
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return nil, app
	}

	user.Mask() // Mask lại để có FakeId nhất quán

	// Lưu lại vào Redis cho lần sau
	if err := u.bzRedis.SaveProfile(ctx, user); err != nil {
		// Có thể log lỗi nhưng không cần return lỗi, vì lấy DB đã thành công
	}

	return user, nil
}
