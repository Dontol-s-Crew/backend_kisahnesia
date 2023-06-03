package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/cerita/entity"
)

type RepositoryCeritaImpl struct {
	DB *sql.DB
}

func ProvideRepisitoryCeritaImpl(DB *sql.DB) RepositoryCeritaImpl {
	return RepositoryCeritaImpl{DB: DB}
}

const (
	INSERT_CERITA_ENTITY = `INSERT INTO "cerita" (user_id,ilutrasi,cover,populer,daerah,judul,genre) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`
	INSERT_ISI_ISI       = `INSERT INTO "isi" (cerita_id,order,paragraft) VALUES ($1,$2,$3)`
)

func (r RepositoryCeritaImpl) RepositoryInsertCerita(data entity.Cerita, ctx context.Context) (int64, error) {
	stmt, err := r.DB.PrepareContext(ctx, INSERT_CERITA_ENTITY)
	if err != nil {
		fmt.Println("error context")
		return 0, err
	}
	var temp int64
	rows, err := stmt.QueryContext(ctx, data.User_id, data.Ilutrasi, data.Cover, 0, data.Daerah, data.Judul, data.Genre)
	if err != nil {
		fmt.Println("query error:" + err.Error())
		return 0, err
	}
	for rows.Next() {

		if err := rows.Scan(&temp); err != nil {
			fmt.Println("error Scan")
			return 0, err
		}

	}
	return temp, nil
}

func (r RepositoryCeritaImpl) RepositoryInsertText(data entity.Text, ctx context.Context) error {
	stmt, err := r.DB.PrepareContext(ctx, INSERT_ISI_ISI)
	if err != nil {
		return err
	}
	for _, val := range data {
		_, err = stmt.QueryContext(ctx, val.Cerita_id, val.Order, val.Paragraft)
		if err != nil {
			return err
		}
	}
	return nil
}
