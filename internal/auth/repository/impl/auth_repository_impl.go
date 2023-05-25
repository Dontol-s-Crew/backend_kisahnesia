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
	SELECT_DATA_USERS = `SELECT * FROM users`
	INSERT_DATA_USER  = `INSERT INTO users(password,email) VALUES ($1,$2)`
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
		if err := rows.Scan(&temp2.Id, &temp2.Password, &temp2.Email, &temp2.Is_admin, &temp2.Time_created, &temp2.Time_updated); err != nil {
			return nil, err
		}
		temp = append(temp, &temp2)
	}
	return temp, nil
}

func (t AuthRepositoryImpl) InsertDataUserRepo(ctx context.Context, userdata entity.Userdb) error {
	stmt, err := t.DB.PrepareContext(ctx, INSERT_DATA_USER)
	if err != nil {
		fmt.Println("prepare err")
		return err
	}
	_, err = stmt.QueryContext(ctx, userdata.Password, userdata.Email)
	if err != nil {
		fmt.Println("Querry err")
		return err
	}
	return nil
}
