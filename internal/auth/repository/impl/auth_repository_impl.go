package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/auth/entity"
)

type AuthRepositoryImpl struct {
	DB *sql.DB
}

func ProvideAuthRepisitoryImpl(DB *sql.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{DB: DB}
}

const (
	SELECT_DATA_USERS = `SELECT * FROM "user"`
	INSERT_DATA_USER  = `INSERT INTO "user" (nama,password,email) VALUES ($1,$2,$3) RETURNING id`
)

func (t AuthRepositoryImpl) GetUserDataRepo(ctx context.Context) (entity.Userdbs, error) {
	stmt, err := t.DB.PrepareContext(ctx, SELECT_DATA_USERS)
	if err != nil {
		return nil, err
	}
	var temp entity.Userdbs
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var temp2 entity.Userdb
		if err := rows.Scan(&temp2.Id, &temp2.Nama, &temp2.Password, &temp2.Email, &temp2.Is_admin, &temp2.Time_created, &temp2.Time_updated); err != nil {
			return nil, err
		}
		temp = append(temp, &temp2)
	}
	return temp, nil
}

func (t AuthRepositoryImpl) InsertDataUserRepo(ctx context.Context, userdata entity.Userdb) (int64, error) {
	stmt, err := t.DB.PrepareContext(ctx, INSERT_DATA_USER)
	if err != nil {
		fmt.Println("prepare err")
		fmt.Println(err)
		return 0, err
	}
	rows, err := stmt.QueryContext(ctx, userdata.Nama, userdata.Password, userdata.Email)
	if err != nil {
		fmt.Println("Querry err")
		return 0, err
	}
	var temp int64
	for rows.Next() {
		if err := rows.Scan(&temp); err != nil {
			return 0, err
		}
	}
	return temp, nil
}
