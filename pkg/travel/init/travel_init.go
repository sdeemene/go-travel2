package init

import (
	"github.com/gorilla/mux"
	dbs "github.com/stdeemene/go-travel2/config/db"
	"github.com/stdeemene/go-travel2/pkg/place"
	"github.com/stdeemene/go-travel2/pkg/travel"
	"github.com/stdeemene/go-travel2/pkg/travel/api"
	"github.com/stdeemene/go-travel2/pkg/travel/repository"
	"github.com/stdeemene/go-travel2/pkg/travel/service"
	"github.com/stdeemene/go-travel2/pkg/user"
)

func InitTravel(dbs *dbs.MongoDB, router *mux.Router, userService user.UserServiceImp, placeService place.PlaceServiceImp) travel.TravelServiceImp {
	travelRepo := repository.NewTravelRepo(dbs)
	travelService := service.NewTravelService(travelRepo, userService, placeService)
	api.NewTravelApi(travelService, router)
	return travelService
}
