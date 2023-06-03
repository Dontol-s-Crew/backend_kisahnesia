package dto

import (
	"mime/multipart"
	"net/http"

	"github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/cerita/entity"
)

type UploadCerita struct {
	Id_user       int64    `json:"Id_user"`
	Genre         string   `json:"Genre"`
	Region        string   `json:"Region"`
	Title         string   `json:"Title"`
	Text          []string `json:"Text"`
	Cover         multipart.File
	Covername     *multipart.FileHeader
	Ilustrasi     multipart.File
	Ilustrasiname *multipart.FileHeader
}

func UploadCeritaToEntityCerita(data UploadCerita) (entity.Cerita, entity.Text) {
	var temp entity.Cerita
	var temp2 entity.Text
	temp.User_id = int64(data.Id_user)
	temp.Judul = data.Title
	temp.Daerah = data.Region
	temp.Genre = data.Genre
	// for index, val := range data.Text {
	// 	var temp3 entity.Isi
	// 	temp3.Order = int64(index)
	// 	temp3.Paragraft = val
	// 	temp2 = append(temp2, &temp3)
	// }
	return temp, temp2
}

func TransformBodyToUploadCerita(r *http.Request) (UploadCerita, error) {
	var temp UploadCerita
	temp.Genre = r.FormValue("genre")
	temp.Title = r.FormValue("title")
	temp.Region = r.FormValue("region")
	var err error
	temp.Cover, temp.Covername, err = r.FormFile("cover")
	if err != nil {
		return UploadCerita{}, err
	}
	defer temp.Cover.Close()
	temp.Ilustrasi, temp.Ilustrasiname, err = r.FormFile("ilustrasi")
	if err != nil {
		return UploadCerita{}, err
	}
	defer temp.Ilustrasi.Close()
	return temp, nil
}
