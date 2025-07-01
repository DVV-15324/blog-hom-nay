package user

import (
	"bloghomnay-project/common"
	entityUser "bloghomnay-project/services/entity/user"
	"context"
)

type BusinessUser interface {
	BzCreateUser(ctx context.Context, cu *entityUser.CreateUserForm) (int, *common.AppError)
	DeleteUserById(ctx context.Context, id int) *common.AppError
	BzGetUsersById(ctx context.Context, id int) (*entityUser.Users, *common.AppError)
	BzGetUsersBySDT(ctx context.Context, phone string) (*entityUser.Users, *common.AppError)
	BzUpdateUser(ctx context.Context, up *entityUser.UpdateUserForm, id int) *common.AppError
}
type ApiUser struct {
	bz BusinessUser
}

func NewApiUser(bz BusinessUser) *ApiUser {
	return &ApiUser{
		bz: bz,
	}
}
