package main

import (
	"fmt"
	"github.com/Vitaly-Baidin/storage-api/httpserver"
	"github.com/Vitaly-Baidin/storage-api/mb"
	"github.com/Vitaly-Baidin/storage-api/storage"
	"github.com/Vitaly-Baidin/storage-api/user"
	"github.com/Vitaly-Baidin/storage-api/user/usrcommand"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	port := os.Getenv("PORT")
	kafkaURL := os.Getenv("KAFKA_URL")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	kafkaGroupID := os.Getenv("KAFKA_GROUP_ID")

	db := storage.NewStorage[string, user.UserGrade]()

	// repositories
	ur := user.NewRepoStorage(db)

	// services
	us := user.NewService(ur)

	// controllers
	uc := user.NewController(us)

	// user commands
	cmdSet := usrcommand.NewSetCommand(us)

	// kafka
	kr := mb.GetKafkaReader(kafkaURL, kafkaTopic, kafkaGroupID)

	go func() {
		mb.Start(kr, cmdSet)
	}()

	mux := http.NewServeMux()

	mux.HandleFunc("/get", uc.Get)

	httpServer := httpserver.New(mux,
		httpserver.Port(port),
	)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		log.Println(fmt.Errorf("app - Run - httpServerSet.Notify: %w", err))
	}

	err := httpServer.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("app - Run - httpServerSet.Shutdown: %w", err))
	}
}
