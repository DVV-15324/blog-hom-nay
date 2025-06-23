package auth

import (
	"bloghomnay-project/common"
	"context"
)

func (bz *BusinessAuth) BzIntrospectToken(ctx context.Context, accessToken string) (*common.JwtClaims, error) {
	claims, error := bz.jwt.ParseToken(ctx, accessToken)
	if error != nil {
		return nil, error
	}
	return claims, nil
}
