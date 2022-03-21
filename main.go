package main

import (
	"fmt"
	"github.com/stdeemene/go-travel2/middleware"
	"github.com/stdeemene/go-travel2/middleware/jwt"
	"github.com/stdeemene/go-travel2/utils/logging"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	dbs "github.com/stdeemene/go-travel2/config/db"
	"github.com/stdeemene/go-travel2/config/env"
	addressInit "github.com/stdeemene/go-travel2/pkg/address/init"
	placeInit "github.com/stdeemene/go-travel2/pkg/place/init"
	travelInit "github.com/stdeemene/go-travel2/pkg/travel/init"
	userInit "github.com/stdeemene/go-travel2/pkg/user/init"
)

func main() {
	env.LoadEnv()
	port := env.GetEnvWithKey("SERVER_PORT")
	routers := mux.NewRouter().StrictSlash(true)
	dbConn := dbs.NewMongoDBConnection()

	registerServices(dbConn, routers)

	handler := registerMiddleWares(routers)

	srv := &http.Server{
		Handler:      handler,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Application is running on Port %v....\n", port)
	log.Fatal(srv.ListenAndServe())

}

func registerMiddleWares(router *mux.Router) *negroni.Negroni {
	logging.Logger()
	n := negroni.Classic()
	n.Use(middleware.Cors())
	router.Use(jwt.ProtectApi)
	n.UseHandler(router)
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("server running"))
		if err != nil {
			return
		}
	})
	return n
}

func registerServices(dbConn *dbs.MongoDB, routers *mux.Router) {
	userServiceImp := userInit.InitUser(dbConn, routers)
	addressServiceImp := addressInit.InitAddress(dbConn, routers)
	placeServiceImp := placeInit.InitPlace(dbConn, routers, addressServiceImp)
	travelInit.InitTravel(dbConn, routers, userServiceImp, placeServiceImp)
}
