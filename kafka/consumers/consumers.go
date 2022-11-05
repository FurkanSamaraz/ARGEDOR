package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
	kafka "github.com/segmentio/kafka-go"
)

func errorlog(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
}
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Millisecond * 5,

		Dialer: &kafka.Dialer{
			Timeout:   10 * time.Second,
			DualStack: true,
			KeepAlive: 5 * time.Millisecond,
		},
	})
}

func multi(kafkaURL string, topic string, groupID string, id int) {
	reader := getKafkaReader(kafkaURL, topic, groupID)
	defer reader.Close()

	fmt.Println("start consuming ... !!")

	for {

		msg, err := reader.ReadMessage(context.Background())
		errorlog(err)
		reader.CommitMessages(context.Background(), msg)
		log.Printf("Value: %s Id: %s", string(msg.Value), strconv.Itoa(id))
		time.Sleep(time.Millisecond * 5)
	}

}
func main() {

	connection := os.Getenv("CONNECTION")
	topic := os.Getenv("TOPIC")
	group1 := os.Getenv("GROUPNO1")
	group2 := os.Getenv("GROUPNO2")
	group3 := os.Getenv("GROUPNO3")
	group4 := os.Getenv("GROUPNO4")
	group5 := os.Getenv("GROUPNO5")

	// get kafka reader using environment variables.
	var wg sync.WaitGroup
	wg.Add(5)
	go func() {
		go multi(connection, topic, group1, 1)

	}()
	go func() {

		go multi(connection, topic, group2, 2)

	}()
	go func() {

		go multi(connection, topic, group3, 3)

	}()
	go func() {

		go multi(connection, topic, group4, 4)

	}()
	go func() {

		go multi(connection, topic, group5, 5)

	}()
	wg.Wait()
}
