package auth

import (
	"bloghomnay-project/common"
	entityAuth "bloghomnay-project/services/entity/auth"
	"context"

	"log"
	"net/http"

	"github.com/google/uuid"
)

func (bz *BusinessAuth) LoginAuth(ctx context.Context, au *entityAuth.LoginForm) (*common.TokenResponse, *common.AppError) {
	err := au.Validate()
	if err != nil {
		app := common.NewAppError(400, http.StatusText(400), err)
		return nil, app
	}

	//Check tai khoan ko ton tai
	auth, err_a := bz.bzAuth.GetAuthByEmail(ctx, au.Email)
	if err_a != nil {
		app := common.NewAppError(500, http.StatusText(500), err_a)
		return nil, app
	}
	//So sánh Hash Password
	ss := bz.hash.CompareHashAndPassword(auth.Password, au.Password, auth.Salt)
	if !ss {
		app := common.NewAppError(500, http.StatusText(500), entityAuth.ErrorEmailAndPassword)
		return nil, app
	}
	//Tạo token
	s := common.NewUID(uint32(auth.UserId), 1) //uid đã mask()
	sub := s.ToBase58()
	tid := uuid.New().String()
	token := bz.jwt.IssueToken(ctx, sub, tid)

	user, err_u := bz.bzUser.BzGetUsersById(ctx, auth.UserId)
	if err_u != nil {
		return nil, err_u
	}

	if err := bz.bzRedis.SaveProfile(ctx, user); err != nil {
		log.Printf("Lỗi lưu profile vào Redis: %v", err)
	}
	return token, nil
}
