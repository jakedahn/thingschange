package producer

import (
	"flag"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {

	var (
		rabbit_host = flag.String("rabbit_host", "3.3.3.3", "Specify the rabbit hostname")
		rabbit_user = flag.String("rabbit_user", "guest", "Specify the rabbit username")
		rabbit_pass = flag.String("rabbit_pass", "guest", "Specify the rabbit password")
	)

	flag.Parse()

	var connection_string = fmt.Sprintf("amqp://%s:%s@%s:5672/", *rabbit_user, *rabbit_pass, *rabbit_host)
	conn, err := amqp.Dial(connection_string)

	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	body := bodyFrom(os.Args)
	err = ch.Publish(
		"",           // exchange
		"task_queue", // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
