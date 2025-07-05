package auth

import (
	"bloghomnay-project/common"
	entityAuth "bloghomnay-project/services/entity/auth"
	entityUser "bloghomnay-project/services/entity/user"
	"context"
	"log"

	"github.com/google/uuid"
	"google.golang.org/api/idtoken"
)

func (bz *BusinessAuth) LoginWithGoogle(ctx context.Context, input *entityAuth.GoogleLoginForm) (*common.TokenResponse, *common.AppError) {
	// Validate Google ID Token
	payload, err := idtoken.Validate(ctx, input.Token, bz.cfg.GoogleClientID) // bz.cfg.GoogleClientID = client ID của bạn
	if err != nil {
		return nil, common.NewAppError(401, "Token Google không hợp lệ", err)
	}

	email := payload.Claims["email"].(string)
	name := payload.Claims["name"].(string)

	// Kiểm tra người dùng đã tồn tại chưa
	auth, err := bz.bzAuth.GetAuthByEmail(ctx, email)
	if err != nil {
		// Nếu user chưa có thì tạo mới (email là duy nhất)
		authEntity := &entityAuth.Auth{
			Email:    email,
			TypeAuth: "google",
		}
		if err := bz.bzAuth.CreateAuth(ctx, authEntity); err != nil {
			return nil, common.NewAppError(500, "Không thể tạo auth", err)
		}

		id, err := bz.bzUser.BzCreateUser(ctx, &entityUser.CreateUserForm{
			Email:     email,
			FirstName: name,
			LastName:  name,
			Phone:     "0394025628",
		})
		if err != nil {
			return nil, err
		}

		authEntity.UserId = id
	}

	// Tạo token
	s := common.NewUID(uint32(auth.UserId), 1)
	sub := s.ToBase58()
	tid := uuid.New().String()
	token := bz.jwt.IssueToken(ctx, sub, tid)

	// Trả profile + token
	user, err_us := bz.bzUser.BzGetUsersById(ctx, auth.UserId)
	if err != nil {
		return nil, err_us
	}
	user.Mask()

	// Lưu vào Redis nếu có
	if err := bz.bzRedis.SaveProfile(ctx, user); err != nil {
		log.Printf("Lỗi lưu profile vào Redis: %v", err)
	}

	return token, nil
}
