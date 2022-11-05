package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

var ch *amqp.Channel

func errorlog(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
}
func main() {

	//Let's connect to your RabbitMQ server first.
	conn, err := amqp.Dial(os.Getenv("CONNECTION"))
	errorlog(err)
	defer conn.Close()

	//Let's create a channel to send our message.
	ch, err = conn.Channel()
	errorlog(err)

	defer ch.Close()
	count := 5
	for i := 0; i < count; i++ {
		channelss(os.Getenv("TOPIC"), strconv.Itoa(i))
	}

}

func channelss(topic string, mesaj string) {
	//Let's define the queue where we will send our message.
	kuyruk, err := ch.QueueDeclare(
		topic, // name .env
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	errorlog(err)

	//What we need to do to share our message.
	err = ch.Publish(
		"",          // exchange
		kuyruk.Name, // Queue name to send. This way we can get the name of the queue we created earlier.
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain", //Type of our message.
			Body:        []byte(mesaj),
		})
	errorlog(err)

	fmt.Println(mesaj, "->", topic)
}
