package api

import (
	"encoding/json"
	"fmt"
	"github.com/stdeemene/go-travel2/middleware/jwt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stdeemene/go-travel2/pkg/user"
	"github.com/stdeemene/go-travel2/pkg/user/model"
	"github.com/stdeemene/go-travel2/utils/response"
)

type UserApi struct {
	UserService user.UserServiceImp
	router      *mux.Router
}

func NewUserApi(userServiceImp user.UserServiceImp, router *mux.Router) *UserApi {
	userApi := &UserApi{UserService: userServiceImp, router: router}
	userApi.UserAuthRouters()
	userApi.UserRouters()
	return userApi
}

func (api UserApi) CreateUser(w http.ResponseWriter, r *http.Request) {
	payload := new(model.SignupReq)

	err := json.NewDecoder(r.Body).Decode(&payload)
	fmt.Println("user body: ", payload)
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.UserService.CreateUser(payload, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusCreated, result)

}
func (api UserApi) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	credentials := new(model.LoginReq)
	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	authenticateUser, err := api.UserService.AuthenticateUser(credentials, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	tokenDetails, err := jwt.GenerateJwtToken(authenticateUser)
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusCreated, tokenDetails)

}
func (api UserApi) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	result, err := api.UserService.GetUser(id, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusOK, result)
}
func (api UserApi) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	res, err := api.UserService.DeleteUser(id, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusAccepted, res)

}
func (api UserApi) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var u model.User
	params := mux.Vars(r)
	id := params["id"]

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := api.UserService.UpdateUser(id, u, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusAccepted, res)

}
func (api UserApi) GetUsers(w http.ResponseWriter, r *http.Request) {
	result, err := api.UserService.GetUsers(r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusOK, result)
}
func (api UserApi) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	email := params["email"]
	res, err := api.UserService.GetUserByEmail(email, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusAccepted, res)

}
func (api UserApi) SearchUser(w http.ResponseWriter, r *http.Request) {
	var filter interface{}
	query := r.URL.Query().Get("q")
	if query != "" {
		err := json.Unmarshal([]byte(query), &filter)
		if err != nil {
			response.BaseResponse(w, http.StatusBadRequest, err.Error())
			return
		}
	}
	result, err := api.UserService.SearchUser(filter, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusOK, result)
}
