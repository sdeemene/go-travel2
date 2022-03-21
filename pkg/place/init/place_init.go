package init

import (
	"github.com/gorilla/mux"
	dbs "github.com/stdeemene/go-travel2/config/db"
	"github.com/stdeemene/go-travel2/pkg/address"
	"github.com/stdeemene/go-travel2/pkg/place"
	"github.com/stdeemene/go-travel2/pkg/place/api"
	"github.com/stdeemene/go-travel2/pkg/place/repository"
	"github.com/stdeemene/go-travel2/pkg/place/service"
)

func InitPlace(dbs *dbs.MongoDB, router *mux.Router, addressService address.AddressServiceImp) place.PlaceServiceImp {
	placeRepo := repository.NewPlaceRepo(dbs)
	placeService := service.NewPlaceService(placeRepo, addressService)
	api.NewPlaceApi(placeService, router)
	return placeService
}
