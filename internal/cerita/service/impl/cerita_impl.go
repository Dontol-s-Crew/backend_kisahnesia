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
	entitycerita.Ilustrasi = iduuid.String() + fileExtensionilustrasi
	entitycerita.Cover = iduuid.String() + fileExtensionCover
	_, err = S.S.RepositoryInsertCerita(entitycerita, ctx)
	// entity.AddCerita_id(entityisi, id)
	// err = S.S.RepositoryInsertText(entityisi, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (S ServiceCeritaImpl) ServiceGetAllCerita(ctx context.Context) (dto.CeritaResponses, error) {
	mapss := make(map[int64]int64)
	cerita, err := S.S.GetAllCeritaRepo(ctx)
	if err != nil {
		return nil, err
	}
	var Response dto.CeritaResponses
	ceritalen := int64(len(cerita))
	for i := int64(0); i < ceritalen; i++ {
		var temp dto.CeritaResponse
		var tempstring []*string
		temp.Story = tempstring
		temp = dto.CeritaToCeritaResponse(cerita[i])
		mapss[cerita[i].Id] = i
		Response = append(Response, &temp)
	}
	isi, err := S.S.GetAllIsiRepo(ctx)
	if err != nil {
		return nil, err
	}
	lenisi := int64(len(isi))
	for i := int64(0); i < lenisi; i++ {
		Response[mapss[isi[i].Id]].Story = append(Response[mapss[isi[i].Id]].Story, &isi[i].Paragraft)
	}
	return Response, nil
}
