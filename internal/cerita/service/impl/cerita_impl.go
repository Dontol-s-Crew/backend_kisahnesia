package impl

import (
	"context"
	"path/filepath"

	"github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/cerita/dto"
	RepisotoryApi "github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/cerita/repository/impl"
	"github.com/Dontol-s-Crew/backend_kisahnesia.git/pkg/utils"
	"github.com/google/uuid"
)

type ServiceCeritaImpl struct {
	S RepisotoryApi.RepositoryCeritaImpl
}

func ProvideServiceCeritaImpl(S RepisotoryApi.RepositoryCeritaImpl) *ServiceCeritaImpl {
	return &ServiceCeritaImpl{S: S}
}

func (S ServiceCeritaImpl) ServicePengumpulanCerita(ctx context.Context, data dto.UploadCerita) error {
	iduuid := uuid.New()
	fileExtensionilustrasi := filepath.Ext(data.Ilustrasiname.Filename)
	err := utils.Createfile("image/ilustrasi/"+iduuid.String()+fileExtensionilustrasi, data.Ilustrasi)
	if err != nil {
		return err
	}
	fileExtensionCover := filepath.Ext(data.Covername.Filename)
	err = utils.Createfile("image/cover/"+iduuid.String()+fileExtensionCover, data.Cover)
	if err != nil {
		return err
	}
	data.Id_user = 2 //hapus setelah debungging
	entitycerita, _ := dto.UploadCeritaToEntityCerita(data)
	entitycerita.Ilutrasi = iduuid.String() + fileExtensionilustrasi
	entitycerita.Cover = iduuid.String() + fileExtensionCover
	_, err = S.S.RepositoryInsertCerita(entitycerita, ctx)
	// entity.AddCerita_id(entityisi, id)
	// err = S.S.RepositoryInsertText(entityisi, ctx)
	if err != nil {
		return err
	}
	return nil
}
