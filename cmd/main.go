package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/fillipehmeireles/order-service/adapters/pg"
	"github.com/fillipehmeireles/order-service/adapters/repositories"
	"github.com/fillipehmeireles/order-service/core/usecases"

	"github.com/fillipehmeireles/order-service/internal/config"
	"github.com/fillipehmeireles/order-service/pkg/handlers/order/api"
	userRepositories "github.com/fillipehmeireles/user-service/adapters/repositories"
)

const ConfigPath = "./"

var (
	binding string
)

func init() {
	flag.StringVar(&binding, "httpbind", ":4001", "address/port to bind listen socket")
	flag.Parse()
}

func setupConfig() (*config.Config, error) {
	return config.NewConfig(ConfigPath)
}
func main() {
	config, err := setupConfig()
	if err != nil {
		log.Fatal(err)
	}

	pgInstance, err := pg.NewPGInstance(config.PgConfig.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer pgInstance.Close()
	pgInstance.Migrate()
	orderRepo := repositories.NewOrderRepository(pgInstance.DB)
	userRepo := userRepositories.NewUserRepository(pgInstance.DB)
	orderUseCase := usecases.NewOrderUseCase(orderRepo, userRepo)
	ws := new(restful.WebService)
	ws = ws.Path("/api")

	api.NewOrderHandler(orderUseCase, ws)
	restful.Add(ws)

	log.Println(binding)

	if err := http.ListenAndServe(binding, nil); err != nil {
		log.Fatal(err)
	}
}
