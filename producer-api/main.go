package main

import (
	"fmt"
	"github.com/Vitaly-Baidin/producer-api/httpserver"
	"github.com/Vitaly-Baidin/producer-api/mb"
	"github.com/Vitaly-Baidin/producer-api/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	port := os.Getenv("PORT")

	kafkaURL := "kafka:9092" // os.Getenv("KAFKA_URL")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")

	// kafka
	kw := mb.NewKafkaWriter(kafkaURL, kafkaTopic)

	// middlewares
	vu := middleware.ValidateUser(mb.SendMsg(kw))
	ba := middleware.BasicAuth(username, password, vu)

	// mux
	mux := http.NewServeMux()

	mux.HandleFunc("/set", ba)

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

	// Shutdown
	err := httpServer.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("app - Run - httpServerSet.Shutdown: %w", err))
	}
}
