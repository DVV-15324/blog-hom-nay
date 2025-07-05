package user

import (
	entityUser "bloghomnay-project/services/entity/user"
	"context"
)

type ResponsitoryUser interface {
	CreateUser(cxt context.Context, user *entityUser.CreateUserForm) (int, error)
	GetUserById(ctx context.Context, id int) (*entityUser.Users, error)
	DeleteUserById(cxt context.Context, id int) error
	GetUserBySDT(ctx context.Context, phone string) (*entityUser.Users, error)
	UpdateUser(cxt context.Context, user *entityUser.UpdateUserForm, id int) error
}
type BzRedis interface {
	SaveProfile(ctx context.Context, data *entityUser.Users) error
	GetProfileEntity(ctx context.Context, userID string) (*entityUser.Users, error)
}
type BusinessUser struct {
	userReponsitory ResponsitoryUser
	bzRedis         BzRedis
}

func NewBusinessUser(userReponsitory ResponsitoryUser, bzRedis BzRedis) *BusinessUser {
	return &BusinessUser{
		userReponsitory: userReponsitory,
		bzRedis:         bzRedis,
	}
}
