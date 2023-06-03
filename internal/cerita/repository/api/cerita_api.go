package api

import (
	"context"

	"github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/cerita/entity"
)

type RepositoryCeritaApi interface {
	RepositoryInsertText(data entity.Text, ctx context.Context) error
	RepositoryInsertCerita(data entity.Cerita, ctx context.Context) (int64, error)
}
