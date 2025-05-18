package auth

import (
	"bloghomnay-project/common"
	entityAuth "bloghomnay-project/services/entity/auth"
	"context"
)

type BusinessAuth interface {
	LoginAuth(ctx context.Context, au *entityAuth.LoginForm) (*common.TokenResponse, *common.AppError)
	BzRegisterAuth(ctx context.Context, auth *entityAuth.RegisterForm) *common.AppError
}
type ApiAuth struct {
	bz BusinessAuth
}

func NewApiAuth(bz BusinessAuth) *ApiAuth {
	return &ApiAuth{
		bz: bz,
	}
}
