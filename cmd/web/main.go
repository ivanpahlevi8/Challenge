package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/configs"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/handler"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/repository"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/service"
)

var appConfig configs.Config

// create session variable so that can be accessd to entire package inculding middleware
var session *scs.SessionManager

func main() {
	// change wether in production or not
	appConfig.InProduction = false

	// create session manager
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	// assign session
	appConfig.Session = session

	// init db
	//inisialisasi data
	host := "127.0.0.1"
	port := "5432"
	user := "postgres"
	password := "03052001ivan"
	dbname := "UserSynapsisDatabase"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//inisialisasi koneksi
	result, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("error = ", err)
	}

	appConfig.DB = result

	// create shop repo
	shopRepo := repository.InitShopRepo()
	shopRepo.Config = &appConfig

	// create shop service
	shopService := service.InitShopService(shopRepo)
	shopService.Config = &appConfig

	// create user repo
	userRepo := repository.InitUserRepo()
	userRepo.Config = &appConfig

	// create user service
	userService := service.InitUserService(userRepo)
	userService.Config = &appConfig
	userService.ItemsRepo = shopRepo

	// create user handler
	userHandler := handler.IntiUserHandler(userService)
	userHandler.Config = &appConfig

	// create item repo
	itemRepo := repository.InitItemRepo()
	itemRepo.Config = &appConfig

	// create item service
	itemService := service.InitItemService(itemRepo)
	itemService.Config = &appConfig

	// create item handler
	itemHandler := handler.InitItemHandler(itemService)
	itemHandler.Config = &appConfig

	// create user activity handler
	userActivityHandler := handler.IntiUserActivityHandler(userService)
	userActivityHandler.ItemService = itemService
	userActivityHandler.ShopService = shopService
	userActivityHandler.Config = &appConfig

	// init middlware
	InitMiddleware(userService, &appConfig)

	// create route
	srv := &http.Server{
		Addr:    ":2020",
		Handler: route(userHandler, itemHandler, userActivityHandler),
	}

	// start routing
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err.Error())
		log.Fatal("Error happen when starting server")

	}
}
