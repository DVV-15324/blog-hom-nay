package auth

import (
	"bloghomnay-project/common"
	entityAuth "bloghomnay-project/services/entity/auth"
	"bloghomnay-project/services/entity/user"
	"context"
	"net/http"
)

func (bz *BusinessAuth) BzRegisterAuth(ctx context.Context, auth *entityAuth.RegisterForm) *common.AppError {
	err := auth.Validate()
	if err != nil {
		app := common.NewAppError(400, http.StatusText(400), err)
		return app
	}
	// check tài khoản tồn tại
	authEmail, err := bz.bzAuth.GetAuthByEmail(ctx, auth.Email)
	if authEmail != nil {
		app := common.NewAppError(409, http.StatusText(http.StatusConflict), entityAuth.ErrorEmailIsExisted)
		return app
	}
	if err != nil {
		app := common.NewAppError(500, http.StatusText(500), err)
		return app
	}

	//Hash Pass
	randStr, err := common.RandomStr()
	if err != nil {
		app := common.NewAppError(500, http.StatusText(500), err)
		return app
	}
	hash, err := bz.hash.GenerateFromPassword(auth.Password, randStr)
	if err != nil {
		app := common.NewAppError(500, http.StatusText(500), err)
		return app
	}
	// Tạo User
	userUid, err_user := bz.bzUser.BzCreateUser(ctx, &user.CreateUserForm{
		Email:     auth.Email,
		FirstName: auth.FirstName,
		LastName:  auth.LastName,
		Phone:     auth.Phone,
	})
	if err_user != nil {
		return err_user
	}
	// Tạo Auth
	err_auth := bz.bzAuth.CreateAuth(ctx, &entityAuth.Auth{
		Salt:     randStr,
		Email:    auth.Email,
		Password: hash,
		UserId:   userUid,
	})
	if err_auth != nil {
		app := common.NewAppError(500, http.StatusText(500), err_auth)
		return app
	}
	return nil
}
