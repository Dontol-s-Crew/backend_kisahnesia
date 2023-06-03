package api

import (
	"context"

	"github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/cerita/dto"
)

type ServiceCeritaImpl interface {
	ServicePengumpulanCerita(ctx context.Context, data dto.UploadCerita) error
	ServiceGetAllCerita(ctx context.Context) (dto.CeritaResponses, error)
}
