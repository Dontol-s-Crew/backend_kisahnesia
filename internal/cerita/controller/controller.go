package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/cerita/dto"
	ServiceApi "github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/cerita/service/impl"
	"github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/global"
)

type CeritaController struct {
	X *mux.Router
	S ServiceApi.ServiceCeritaImpl
}

func ProvideCeritaController(X *mux.Router, S ServiceApi.ServiceCeritaImpl) CeritaController {
	return CeritaController{X: X, S: S}
}

func (CC CeritaController) HandlerUploadPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20) // Set the maximum memory to 32MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	UploadData, err := dto.TransformBodyToUploadCerita(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	CC.S.ServicePengumpulanCerita(ctx, UploadData)
}

func (CC CeritaController) GetImage(w http.ResponseWriter, r *http.Request) {
	hashmap := mux.Vars(r)
	fileBytes, err := ioutil.ReadFile("image/" + hashmap["catagory"] + "/" + hashmap["id"])
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return
}
func (CC CeritaController) GetAllCerita(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	CeritaAll, err := CC.S.ServiceGetAllCerita(ctx)
	if err != nil {
		panic(err)
	}
	arr, err := json.Marshal(CeritaAll)
	if err != nil {
		panic(err)
	}
	arr, err = json.MarshalIndent(CeritaAll, "", " ")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(arr)
	return
}

func (CC CeritaController) InitializeController() {
	router := CC.X.PathPrefix(global.API_PATH_ROOT_CERITA).Subrouter()
	router.HandleFunc(global.API_PATH_POST_UPLOAD, CC.HandlerUploadPost).Methods(http.MethodPost)
	router.HandleFunc(global.API_PATH_GET_IMAGE, CC.GetImage).Methods(http.MethodGet)
	router.HandleFunc(global.API_PATH_GET_ALL, CC.GetAllCerita).Methods(http.MethodGet)
}
