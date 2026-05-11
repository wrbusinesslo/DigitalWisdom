package service

import (
	DigitalWisdomApi "DigitalWisdom/service/api/DigitalWisdom"
	"DigitalWisdom/service/controller/aclCtrl"
	"DigitalWisdom/service/controller/purchaseCtrl"
	"DigitalWisdom/service/internal/config"
	"DigitalWisdom/service/internal/database"
	"context"
	"fmt"
	"go.uber.org/dig"
	"net/http"
)

func DigitalWisdom() Service {
	once.Do(func() {
		srv = &DigitalWisdomServer{}
	})

	return srv
}

type DigitalWisdomServer struct{}

func (srv *DigitalWisdomServer) Run() {

	container := dig.New()
	srv.provideConfig(container)
	srv.provideService(container)
	srv.provideController(container)
	srv.provideCore(container)

	srv.invokeApiRoutes(container)

	if err := container.Invoke(srv.run); err != nil {
		panic(err)
	}

}

func (srv *DigitalWisdomServer) provideConfig(container *dig.Container) {

	if err := container.Provide(config.NewDigitalWisdom); err != nil {
		panic(err)
	}
}

func (srv *DigitalWisdomServer) provideService(container *dig.Container) {
	if err := container.Provide(func() context.Context {
		return context.TODO()
	}); err != nil {
		panic(err)
	}

	if err := container.Provide(database.NewDigitalWisdom); err != nil {
		panic(err)
	}

	if err := container.Provide(DigitalWisdomApi.NewServer); err != nil {
		panic(err)
	}

	if err := container.Provide(DigitalWisdomApi.NewRouterRoot); err != nil {
		panic(err)
	}

	if err := container.Provide(DigitalWisdomApi.NewGinEngine); err != nil {
		panic(err)
	}

}

func (srv *DigitalWisdomServer) provideController(container *dig.Container) {
	if err := container.Provide(aclCtrl.NewAcl); err != nil {
		panic(err)
	}
	if err := container.Provide(purchaseCtrl.NewPurchaseCtrl); err != nil {
		panic(err)
	}
}

func (srv *DigitalWisdomServer) invokeApiRoutes(container *dig.Container) {
	if err := container.Invoke(DigitalWisdomApi.NewServer); err != nil {
		panic(err)
	}

	if err := container.Invoke(DigitalWisdomApi.NewGinEngine); err != nil {
		panic(err)
	}
	if err := container.Invoke(DigitalWisdomApi.NewAcl); err != nil {
		panic(err)
	}
	if err := container.Invoke(DigitalWisdomApi.NewPurchase); err != nil {
		panic(err)
	}
}

func (srv *DigitalWisdomServer) provideCore(container *dig.Container) {}

func (srv *DigitalWisdomServer) run(server *http.Server) {
	fmt.Printf("DigitalWisdom starts at %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
