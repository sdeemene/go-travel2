package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stdeemene/go-travel2/pkg/travel"
	"github.com/stdeemene/go-travel2/pkg/travel/model"
	"github.com/stdeemene/go-travel2/utils/response"
	"net/http"
)

type TravelApi struct {
	TravelService travel.TravelServiceImp
	router        *mux.Router
}

func NewTravelApi(travelService travel.TravelServiceImp, router *mux.Router) *TravelApi {
	travelApi := &TravelApi{TravelService: travelService, router: router}
	travelApi.TravelRouters()
	return travelApi
}

func (api TravelApi) CreateTravel(w http.ResponseWriter, r *http.Request) {
	payload := new(model.TravelReq)

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := api.TravelService.CreateTravel(payload, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusCreated, res)
}

func (api TravelApi) GetTravel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	result, err := api.TravelService.GetTravel(id, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
	}
	response.BaseResponse(w, http.StatusOK, result)
}

func (api TravelApi) DeleteTravel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	result, err := api.TravelService.DeleteTravel(id, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
	}
	response.BaseResponse(w, http.StatusAccepted, result)
}

func (api TravelApi) UpdateTravel(w http.ResponseWriter, r *http.Request) {
	var t model.Travel
	params := mux.Vars(r)
	id := params["id"]
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := api.TravelService.UpdateTravel(id, t, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusAccepted, res)
}

func (api TravelApi) GetTravels(w http.ResponseWriter, r *http.Request) {
	result, err := api.TravelService.GetTravels(r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusOK, result)
}

func (api TravelApi) SearchTravel(w http.ResponseWriter, r *http.Request) {
	var filter interface{}
	query := r.URL.Query().Get("q")
	if query != "" {
		err := json.Unmarshal([]byte(query), &filter)
		if err != nil {
			response.BaseResponse(w, http.StatusBadRequest, err.Error())
			return
		}
	}
	result, err := api.TravelService.SearchTravel(filter, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusOK, result)
}

func (api TravelApi) GetUserTravels(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	res, err := api.TravelService.GetUserTravels(id, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusOK, res)
}

func (api TravelApi) RateTravel(w http.ResponseWriter, r *http.Request) {
	rateValue := new(model.TravelRateReq)
	params := mux.Vars(r)
	id := params["id"]

	err := json.NewDecoder(r.Body).Decode(&rateValue)
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	res, err := api.TravelService.RateTravel(id, rateValue, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusAccepted, res)
}
