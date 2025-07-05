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
	payload, err := idtoken.Validate(ctx, input.Token, bz.cfg.GoogleClientID)
	if err != nil {
		log.Printf("Lỗi validate token Google: %v", err)
		return nil, common.NewAppError(401, "Token Google không hợp lệ", err)
	}

	// --- Bắt lỗi claim không tồn tại hoặc không đúng định dạng ---
	emailClaim, ok := payload.Claims["email"]
	if !ok {
		return nil, common.NewAppError(400, "Không tìm thấy email trong token", nil)
	}
	email, ok := emailClaim.(string)
	if !ok {
		return nil, common.NewAppError(400, "Email không hợp lệ", nil)
	}

	nameClaim, ok := payload.Claims["name"]
	if !ok {
		return nil, common.NewAppError(400, "Không tìm thấy name trong token", nil)
	}
	name, ok := nameClaim.(string)
	if !ok {
		return nil, common.NewAppError(400, "Tên không hợp lệ", nil)
	}

	// --- Kiểm tra user đã tồn tại hay chưa ---
	auth, err := bz.bzAuth.GetAuthByEmail(ctx, email)
	if err != nil || auth == nil {
		log.Printf("Người dùng chưa tồn tại hoặc auth nil: %v", err)

		id, err_id := bz.bzUser.BzCreateUser(ctx, &entityUser.CreateUserForm{
			Email:     email,
			FirstName: name,
			LastName:  name,
			Phone:     "0394025628", // TODO: Nếu cần có thể để rỗng
		})
		if err_id != nil {
			log.Printf("Lỗi tạo user mới: %v", err)
			return nil, err_id
		}

		authEntity := &entityAuth.Auth{
			Email:  email,
			UserId: id,
		}
		if err := bz.bzAuth.CreateAuth(ctx, authEntity); err != nil {
			log.Printf("Lỗi tạo auth record: %v", err)
			return nil, common.NewAppError(500, "Không thể tạo auth", err)
		}

	}

	// --- Tạo token ---
	log.Printf("UserID dùng để tạo token: %v", auth.UserId)

	s := common.NewUID(uint32(auth.UserId), 1)
	sub := s.ToBase58()
	tid := uuid.New().String()
	token := bz.jwt.IssueToken(ctx, sub, tid)

	// --- Lấy profile người dùng ---
	user, err_u := bz.bzUser.BzGetUsersById(ctx, auth.UserId)
	if err != nil {
		log.Printf("Lỗi lấy user profile: %v", err)
		return nil, err_u
	}
	user.Mask()

	// --- Lưu vào Redis (tùy chọn) ---
	if err := bz.bzRedis.SaveProfile(ctx, user); err != nil {
		log.Printf("Lỗi lưu profile vào Redis: %v", err)
	}

	return token, nil
}
