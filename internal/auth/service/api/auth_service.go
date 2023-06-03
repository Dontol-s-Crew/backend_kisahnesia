package api

import (
	"context"

	"github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/auth/dto"
)

type AuthServiceInterface interface {
	LoginServiceAuth(ctx context.Context, x dto.User) (string, error)
	RegisterServiceAuth(ctx context.Context, x dto.UserRegister) (string, error)
}
