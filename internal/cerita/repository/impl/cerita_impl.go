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
	INSERT_CERITA_ENTITY = `INSERT INTO "cerita" (user_id,ilustrasi,cover,populer,daerah,judul,genre) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`
	INSERT_ISI_ISI       = `INSERT INTO "isi" (cerita_id,order,paragraft) VALUES ($1,$2,$3)`
	SELECT_ALL_CERITA    = `SELECT * FROM "cerita"`
	SELECT_ALL_USER      = `SELECT * FROM "user"`
	SELECT_ALL_ISI       = `SELECT * FROM "isi" ORDER BY "cerita_id" ASC, "order" ASC`
)

func (r RepositoryCeritaImpl) RepositoryInsertCerita(data entity.Cerita, ctx context.Context) (int64, error) {
	stmt, err := r.DB.PrepareContext(ctx, INSERT_CERITA_ENTITY)
	if err != nil {
		fmt.Println("error context")
		return 0, err
	}
	var temp int64
	rows, err := stmt.QueryContext(ctx, data.User_id, data.Ilustrasi, data.Cover, 0, data.Daerah, data.Judul, data.Genre)
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

func (r RepositoryCeritaImpl) GetAllCeritaRepo(ctx context.Context) (entity.ArrayCerita, error) {
	stmt, err := r.DB.PrepareContext(ctx, SELECT_ALL_CERITA)
	if err != nil {
		fmt.Println("error context")
		return nil, err
	}
	var CeritaRows entity.ArrayCerita
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		fmt.Println("query error:" + err.Error())
		return nil, err
	}
	for rows.Next() {
		var temp1 entity.Cerita
		if err := rows.Scan(&temp1.Id, &temp1.User_id, &temp1.Populer, &temp1.Ilustrasi, &temp1.Cover, &temp1.Daerah, &temp1.Judul, &temp1.Genre, &temp1.Status, &temp1.Time_created, &temp1.Time_updated); err != nil {
			fmt.Println("error Scan")
			return nil, err
		}
		CeritaRows = append(CeritaRows, &temp1)
	}
	return CeritaRows, nil
}

func (r RepositoryCeritaImpl) GetAllIsiRepo(ctx context.Context) (entity.Text, error) {
	stmt, err := r.DB.PrepareContext(ctx, SELECT_ALL_ISI)
	if err != nil {
		fmt.Println("error context")
		return nil, err
	}
	var IsiRows entity.Text
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		fmt.Println("query error:" + err.Error())
		return nil, err
	}
	for rows.Next() {
		var temp1 entity.Isi
		if err := rows.Scan(&temp1.Id, &temp1.Cerita_id, &temp1.Order, &temp1.Paragraft); err != nil {
			fmt.Println("error Scan")
			return nil, err
		}
		IsiRows = append(IsiRows, &temp1)
	}
	return IsiRows, nil
}
