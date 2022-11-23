package mb

import (
	"context"
	"fmt"
	"github.com/Vitaly-Baidin/storage-api/user/usrcommand"
	"github.com/segmentio/kafka-go"
	"log"
)

func GetKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaURL},

		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

func Start(r *kafka.Reader, c usrcommand.Command) {
	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
			return
		}

		err = c.Execute(m.Value)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Printf("message at topic:%v partition:%v offset:%v	store: %s\n", m.Topic, m.Partition, m.Offset, string(m.Key))
	}
}
