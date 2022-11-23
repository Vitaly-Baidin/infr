package mb

import (
	"fmt"
	"github.com/Vitaly-Baidin/producer-api/response"
	"github.com/segmentio/kafka-go"
	"io"
	"log"
	"net/http"
)

func NewKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func SendMsg(kafkaWriter *kafka.Writer) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(fmt.Sprintf("kafka.Writer - SendMsg - io.ReadAll: %v", err))
			response.SendErrorResponse(rw, http.StatusInternalServerError, err.Error())
			return
		}
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("address-%s", r.RemoteAddr)),
			Value: body,
		}

		err = kafkaWriter.WriteMessages(r.Context(), msg)
		if err != nil {
			log.Println(fmt.Sprintf("kafka.Writer - SendMsg - kafkaWriter.WriteMessages: %v", err))
			response.SendErrorResponse(rw, http.StatusInternalServerError, err.Error())
			return
		}
		log.Println("SEND OK")
	}
}
