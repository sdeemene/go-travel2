package init

import (
	"github.com/gorilla/mux"
	dbs "github.com/stdeemene/go-travel2/config/db"

	"github.com/stdeemene/go-travel2/pkg/user"
	"github.com/stdeemene/go-travel2/pkg/user/api"
	"github.com/stdeemene/go-travel2/pkg/user/repository"
	"github.com/stdeemene/go-travel2/pkg/user/service"
)

func InitUser(dbs *dbs.MongoDB, router *mux.Router) user.UserServiceImp {
	userRepo := repository.NewUserRepo(dbs)
	userService := service.NewUserService(userRepo)
	api.NewUserApi(userService, router)
	return userService
}
