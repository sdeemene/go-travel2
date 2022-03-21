package init

import (
	"github.com/gorilla/mux"
	dbs "github.com/stdeemene/go-travel2/config/db"
	"github.com/stdeemene/go-travel2/pkg/address"
	"github.com/stdeemene/go-travel2/pkg/address/api"
	"github.com/stdeemene/go-travel2/pkg/address/repository"
	"github.com/stdeemene/go-travel2/pkg/address/service"
)

func InitAddress(dbs *dbs.MongoDB, router *mux.Router) address.AddressServiceImp {
	addressRepo := repository.NewAddressRepo(dbs)
	addressService := service.NewAddressService(addressRepo)
	api.NewAddressApi(addressService, router)
	return addressService
}
