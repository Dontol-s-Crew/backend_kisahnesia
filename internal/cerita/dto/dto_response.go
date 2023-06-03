package dto

import (
	"time"

	"github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/cerita/entity"
)

type CeritaResponse struct {
	Id          int64     `json:"Id"`
	Views       int64     `json:"Views"`
	Title       string    `json:"Title"`
	Image       string    `json:"Image"`
	Region      string    `json:"Region"`
	Time_upload time.Time `json:"Time_upload"`
	Ilustration string    `json:"Ilustration"`
	Story       []*string `json:"Story"`
}

type CeritaResponses []*CeritaResponse

func CeritaToCeritaResponse(data *entity.Cerita) CeritaResponse {
	var temp CeritaResponse
	temp.Id = data.Id
	temp.Views = data.Populer
	temp.Image = data.Cover
	temp.Ilustration = data.Ilustrasi
	temp.Title = data.Judul
	temp.Region = data.Daerah
	temp.Time_upload = data.Time_created
	return temp
}
