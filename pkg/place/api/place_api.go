package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stdeemene/go-travel2/pkg/place"
	"github.com/stdeemene/go-travel2/pkg/place/model"
	"github.com/stdeemene/go-travel2/utils/response"
	"net/http"
)

type PlaceApi struct {
	PlaceService place.PlaceServiceImp
	router       *mux.Router
}

func NewPlaceApi(placeService place.PlaceServiceImp, router *mux.Router) *PlaceApi {
	placeApi := &PlaceApi{PlaceService: placeService, router: router}
	placeApi.PlaceRouters()
	return placeApi
}

func (api PlaceApi) CreatePlace(w http.ResponseWriter, r *http.Request) {
	payload := new(model.PlaceReq)

	err := json.NewDecoder(r.Body).Decode(&payload)
	fmt.Println("place body: ", payload)
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.PlaceService.CreatePlace(payload, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusCreated, result)
}

func (api PlaceApi) GetPlace(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	result, err := api.PlaceService.GetPlace(id, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
	}
	response.BaseResponse(w, http.StatusOK, result)
}

func (api PlaceApi) DeletePlace(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	result, err := api.PlaceService.DeletePlace(id, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
	}
	response.BaseResponse(w, http.StatusAccepted, result)
}

func (api PlaceApi) UpdatePlace(w http.ResponseWriter, r *http.Request) {
	var p model.Place
	params := mux.Vars(r)
	id := params["id"]
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := api.PlaceService.UpdatePlace(id, p, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusAccepted, res)

}

func (api PlaceApi) GetPlaces(w http.ResponseWriter, r *http.Request) {
	result, err := api.PlaceService.GetPlaces(r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusOK, result)
}

func (api PlaceApi) SearchPlace(w http.ResponseWriter, r *http.Request) {
	var filter interface{}
	query := r.URL.Query().Get("q")
	if query != "" {
		err := json.Unmarshal([]byte(query), &filter)
		if err != nil {
			response.BaseResponse(w, http.StatusBadRequest, err.Error())
			return
		}
	}
	result, err := api.PlaceService.SearchPlace(filter, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusOK, result)
}
