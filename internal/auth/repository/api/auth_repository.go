package api

import (
	"context"

	"github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/auth/entity"
)

type AuthRepoInterface interface {
	GetUserDataRepo(ctx context.Context) (entity.Userdbs, error)
	InsertDataUserRepo(ctx context.Context, userdata entity.Userdb) error
}
