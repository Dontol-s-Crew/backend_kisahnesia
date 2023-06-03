package impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/auth/dto"

	repositoryApiPkg "github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/auth/repository/api"
	"github.com/Dontol-s-Crew/backend_kisahnesia.git/pkg/utils"
)

type AuthServiceImpl struct {
	as repositoryApiPkg.AuthRepoInterface
}

func ProvideAuthServiceImpl(as repositoryApiPkg.AuthRepoInterface) *AuthServiceImpl {
	return &AuthServiceImpl{as: as}
}

func (t AuthServiceImpl) LoginServiceAuth(ctx context.Context, x dto.User) (string, error) {
	data, err := t.as.GetUserDataRepo(ctx)
	x.Password = utils.Hashpassword(x.Password) //pasword terhash
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	jsondata := dto.FromEntityToJson(data)
	for _, isi := range jsondata {
		if isi.Email == x.Email {
			if isi.Password != x.Password {
				err := errors.New("wrong password")
				return "", err
			}
		}
		jwttoken, err := utils.Generatejwttoken(isi.Id, isi.Nama, isi.Email, isi.Is_admin)
		if err != nil {
			return "", err
		}
		return jwttoken, nil
	}
	err = errors.New("No account have the email or password")
	return "", err
}

func (t AuthServiceImpl) RegisterServiceAuth(ctx context.Context, x dto.UserRegister) (string, error) {
	// _, err := t.as.GetUserDataRepo(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return "", err
	// }
	dbdata := dto.ChangeUserRegisterToAuthEntity(x)
	id, err := t.as.InsertDataUserRepo(ctx, dbdata)
	dbdata.Id = id
	dbdata.Password = utils.Hashpassword(dbdata.Password) //pasword terhash
	if err != nil {
		return "", err
	}
	tokenjwt, err := utils.Generatejwttoken(dbdata.Id, dbdata.Nama, dbdata.Email, false)
	if err != nil {
		return "", err
	}
	return tokenjwt, nil
}
