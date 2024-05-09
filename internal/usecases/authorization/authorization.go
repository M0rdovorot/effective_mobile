package authorization

import (
	"context"

	ctxusecase "github.com/M0rdovorot/effective_mobile/internal/usecases/context"
)

const (
	user_token = "user_token"
	admin_token = "admin_token"
)

func AuthorizeUser(token string) (bool, error) {
	if token == user_token {
		return false, nil
	}
	if token == admin_token {
		return true, nil
	}
	return false, ErrUnauthorized
}

func AuthorizeUserCtx(ctx context.Context) (bool, error) {
	token := ctxusecase.GetToken(ctx)
	return AuthorizeUser(token)
}

func AuthorizeAdmin(token string) (error) {
	if token == admin_token {
		return nil
	}
	return ErrUnauthorized
}

func AuthorizeAdminCtx(ctx context.Context) (error) {
	token := ctxusecase.GetToken(ctx)
	return AuthorizeAdmin(token)
}