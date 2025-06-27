package auth

import (
	"time"
)

type TypeAuth string

var (
	LocalAuth TypeAuth = "local"
	Facebook  TypeAuth = "facebook"
	Google    TypeAuth = "google"
)

type Auth struct {
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	UserId     int       `json:"-"`
	Salt       string    `json:"salt"`
	TypeAuth   TypeAuth  `json:"type_auth"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Deleted_at time.Time `json:"deleted_at,omitempty"`
}

func TableName() string {
	return "auths"
}
