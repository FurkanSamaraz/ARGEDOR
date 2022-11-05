package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

func errorlog(err error, msg string) {
	if err != nil {
		log.Fatalln(err, msg)

	}
}
func init() {
	err := godotenv.Load("../.env")
	errorlog(err, "")
}
func main() {
	connection := os.Getenv("CONNECTION")
	topic := os.Getenv("TOPIC")
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", connection, topic, partition)
	errorlog(err, "failed to dial leader")

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	count := 5
	for i := 0; i < count; i++ {

		_, err = conn.WriteMessages(

			kafka.Message{Value: []byte(strconv.Itoa(i))},
		)
	}

	errorlog(err, "failed to write messages")

	err = conn.Close()
	errorlog(err, "failed to close writer")

}
