package entity

import "time"

type Cerita struct {
	Id           int64     `db:"Id"`
	User_id      int64     `db:"User_id"`
	Ilutrasi     string    `db:"Ilutrasi"`
	Cover        string    `db:"Cover"`
	Populer      int64     `db:"Populer"`
	Judul        string    `db:"Judul"`
	Daerah       string    `db:"Daerah"`
	Status       bool      `db:"Status"`
	Genre        string    `db:"Genre"`
	Time_updated time.Time `db:"Time_updated"`
	Time_created time.Time `db:"Time_created"`
}

type Isi struct {
	Id        int64  `db:"Id"`
	Cerita_id int64  `db:"Cerita_id"`
	Order     int64  `db:"Order"`
	Paragraft string `db:"Paragraft"`
}

type Text []*Isi

func AddCerita_id(data Text, id int64) {
	for index := range data {
		data[index].Cerita_id = id
	}
}
