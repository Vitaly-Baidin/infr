package main

import (
	"fmt"
	"github.com/Vitaly-Baidin/infr/internal/user"
	"github.com/Vitaly-Baidin/infr/middleware"
	"github.com/Vitaly-Baidin/infr/pkg/httpserver"
	"github.com/Vitaly-Baidin/infr/pkg/storage"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	username = "root"
	password = "rootroot"

	portSet = "8080"
	portGet = "9090"
)

func main() {
	db := storage.NewStorage[string, user.UserGrade]()

	// repositories
	ur := user.NewRepoStorage(db)

	// services
	us := user.NewService(ur)

	// controllers
	uc := user.NewController(us)

	// middlewares

	ba := middleware.NewBasicAuth(username, password, http.HandlerFunc(uc.Set))

	// mux's
	muxSet := http.NewServeMux()
	muxGet := http.NewServeMux()

	muxSet.Handle("/set", ba)

	muxGet.HandleFunc("/get", uc.Get)

	httpServerSet := httpserver.New(muxSet,
		httpserver.Port(portSet),
	)

	httpServerGet := httpserver.New(muxGet,
		httpserver.Port(portGet),
	)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("app - Run - signal: " + s.String())
	case err := <-httpServerSet.Notify():
		log.Println(fmt.Errorf("app - Run - httpServerSet.Notify: %w", err))
	case err := <-httpServerGet.Notify():
		log.Println(fmt.Errorf("app - Run - httpServerGet.Notify: %w", err))
	}

	// Shutdowns
	err := httpServerSet.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("app - Run - httpServerSet.Shutdown: %w", err))
	}

	err = httpServerGet.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("app - Run - httpServerGet.Shutdown: %w", err))
	}
}
