package controllers

import (
	"encoding/json"
	"log"

	"github.com/hrmadani/nmapdojo-report/pkg/models"
	"github.com/streadway/amqp"
)

//Constants
const (
	RabbitMQUserName         = "guest"
	RabbitMQPassword         = "guest"
	RabbitMQServer           = "localhost:5672/"
	RabbitMQName             = "user_report"
	RabbitMQDurable          = false
	RabbitMQDeleteWhenUnused = false
	RabbitMQExclusive        = false
	RabbitMQNoWait           = false
	RabbitMQConsumer         = ""
	RabbitMQAutoAck          = true
	RabbitMQNoLocal          = false
)

var (
	UserReport models.UserReport
	Report     models.Report
	ReportLog  models.ReportLog
)

//Error Handler
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//The main function in this service
//Consume messages from RabbitMQ
//Call the appropriate function to save messages to database
func ConsumeFromRabbit() {
	conn, err := amqp.Dial("amqp://" + RabbitMQUserName + ":" + RabbitMQPassword + "@" + RabbitMQServer)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		RabbitMQName,
		RabbitMQDurable,
		RabbitMQDeleteWhenUnused,
		RabbitMQExclusive,
		RabbitMQNoWait,
		nil, // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		RabbitMQConsumer,
		RabbitMQAutoAck,
		RabbitMQExclusive,
		RabbitMQNoLocal,
		RabbitMQNoWait,
		nil, // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			json.Unmarshal([]byte(d.Body), &UserReport)
			failOnError(err, "Failed to Unmarshal")

			switch UserReport.Action {
			case "add":
				ActionIsAdd()
			default:
				ActionIsNotAdd()
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

//If the action is add :
//Add new Report
//Add new Log
func ActionIsAdd() {
	log.Printf("Action is ADD ==> ")
	//Save new Report
	reportId, _ := Report.Save(UserReport)

	//Save new Log
	ReportLog.Save(UserReport, reportId)
}

//If the action is add :
//Change the expire_time in Report
//Add new Log
func ActionIsNotAdd() {
	log.Printf("Action is NOT ADD ==> ")

	Report.UpdateExpireTime(UserReport)

	//Save new Log
	ReportLog.Save(UserReport, uint(UserReport.ReportId))
}
