package auth

import (
	"bloghomnay-project/common"
	entityAuth "bloghomnay-project/services/entity/auth"
	entityUser "bloghomnay-project/services/entity/user"
	"context"
)

type BzRedis interface {
	SaveProfile(ctx context.Context, data *entityUser.Users) error
	GetProfileEntity(ctx context.Context, userID string) (*entityUser.Users, error)
}

type ReponsitoryAuth interface {
	CreateAuth(cxt context.Context, auth *entityAuth.Auth) error
	GetAuthByEmail(ctx context.Context, email string) (*entityAuth.Auth, error)
}

type Hash interface {
	GenerateFromPassword(password string, salt string) (string, error)
	CompareHashAndPassword(passwordStr string, password string, salt string) bool
}

type BzUser interface {
	BzCreateUser(ctx context.Context, cu *entityUser.CreateUserForm) (int, *common.AppError)
	BzGetUsersById(ctx context.Context, id int) (*entityUser.Users, *common.AppError)
}

type JwtService interface {
	ParseToken(ctx context.Context, tokenStr string) (*common.JwtClaims, error)
	IssueToken(cxt context.Context, sub string, tid string) *common.TokenResponse
}

type BusinessAuth struct {
	jwt     JwtService
	bzUser  BzUser
	hash    Hash
	bzAuth  ReponsitoryAuth
	bzRedis BzRedis
}

func NewBusinessAuth(jwt JwtService, bzUser BzUser, h Hash, bzAuth ReponsitoryAuth, bzRedis BzRedis) *BusinessAuth {
	return &BusinessAuth{
		jwt:     jwt,
		bzUser:  bzUser,
		hash:    h,
		bzAuth:  bzAuth,
		bzRedis: bzRedis,
	}
}
