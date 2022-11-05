package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

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
		panic(err)
	}
}
func main() {

	var wg sync.WaitGroup
	wg.Add(5)
	//Connecting to our RabbitMQ Server
	conn, err := amqp.Dial(os.Getenv("CONNECTION"))
	errorlog(err)
	defer conn.Close()

	//Let's create a channel to communicate
	ch, err = conn.Channel()
	errorlog(err)

	defer ch.Close()
	//We create 5 simultaneous listeners.
	go func() {
		go channelss(os.Getenv("TOPIC"), os.Getenv("GROUPNO1"))

	}()

	go func() {
		go channelss(os.Getenv("TOPIC"), os.Getenv("GROUPNO2"))

	}()

	go func() {
		go channelss(os.Getenv("TOPIC"), os.Getenv("GROUPNO3"))

	}()
	go func() {
		go channelss(os.Getenv("TOPIC"), os.Getenv("GROUPNO4"))

	}()
	go func() {
		go channelss(os.Getenv("TOPIC"), os.Getenv("GROUPNO5"))

	}()
	wg.Wait() //We are waiting for them.

}
func channelss(topic string, group string) {

	//We define our tail
	_, _ = ch.QueueDeclare(
		topic, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err := ch.Qos(5, 0, false); err != nil {
		log.Fatal("Qos Setting was unsuccessfull")
	}
	//Here we are listening to our tail.
	msgs, err := ch.Consume(
		topic, // .env
		group, // consumer .env
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	errorlog(err)

	/*
		Here, we created a channel so that the program does not
		close while our function that works with goroutine is running.
	*/
	forever := make(chan bool)
	/*
	   we define simultaneous read operation and 5 second lifetime
	*/
	go func() {

		for a := range msgs {
			err = ch.Ack(a.DeliveryTag, false) // <-- Difference
			errorlog(err)

			fmt.Printf("Received Event %s\n", string(a.Body))
			time.Sleep(time.Millisecond * 5)
		}
	}()
	go func() {
		//Here we pull messages from the queue, if any.
		for msg := range msgs {
			log.Printf("Value: %s - GroupID:%s", msg.Body, group)

		}

	}()

	log.Printf("[*] Queue Listening [*]")

	<-forever
}
