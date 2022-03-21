package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stdeemene/go-travel2/pkg/address"
	"github.com/stdeemene/go-travel2/pkg/address/model"
	"github.com/stdeemene/go-travel2/utils/response"
)

type AddressApi struct {
	AddressService address.AddressServiceImp
	router         *mux.Router
}

func NewAddressApi(addressServiceImp address.AddressServiceImp, router *mux.Router) *AddressApi {
	addressApi := &AddressApi{AddressService: addressServiceImp, router: router}
	addressApi.AddressRouters()
	return addressApi
}

func (api AddressApi) CreateAddress(w http.ResponseWriter, r *http.Request) {

	var addr model.Address

	err := json.NewDecoder(r.Body).Decode(&addr)
	fmt.Println("1 address body: ", addr)
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.AddressService.CreateAddress(addr, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusCreated, result)
}

func (api AddressApi) GetAddress(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]
	result, err := api.AddressService.GetAddress(id, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusOK, result)
}

func (api AddressApi) DeleteAddress(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	res, err := api.AddressService.DeleteAddress(id, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusAccepted, res)

}

func (api AddressApi) UpdateAddress(w http.ResponseWriter, r *http.Request) {

	var addr model.Address
	params := mux.Vars(r)
	id := params["id"]

	err := json.NewDecoder(r.Body).Decode(&addr)
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := api.AddressService.UpdateAddress(id, addr, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusAccepted, res)
}

func (api AddressApi) GetAddresses(w http.ResponseWriter, r *http.Request) {
	result, err := api.AddressService.GetAddresses(r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusOK, result)
}

func (api AddressApi) SearchAddress(w http.ResponseWriter, r *http.Request) {
	var filter interface{}
	query := r.URL.Query().Get("q")
	if query != "" {
		err := json.Unmarshal([]byte(query), &filter)
		if err != nil {
			response.BaseResponse(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	result, err := api.AddressService.SearchAddress(filter, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusOK, result)
}
