package entity

import "time"

type Userdb struct {
	Id           int64     `db:"Id"`
	Password     string    `db:"Password"`
	Email        string    `db:"Email"`
	Nama         string    `db:"Nama"`
	Is_admin     bool      `db:"Is_admin"`
	Time_updated time.Time `db:"Time_updated"`
	Time_created time.Time `db:"Time_created"`
}

type Userdbs []*Userdb
