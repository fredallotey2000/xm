package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	confg "xm/internal/config"
	handler "xm/internal/http/handlers"
	v "xm/internal/http/validator"

	db "xm/internal/storage/mysql"

	lg "xm/internal/logger"

	comp "xm/pkg/company"
	usr "xm/pkg/user"

	"xm/internal/kafka"
)

//channels for shared memory communication betweeen message queue(kafka) and http request
var producerChan = make(chan kafka.Message)

func main() {
	log.Println("starting servers...")
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func openLogFile(logFileName string) *os.File {

	f, err := os.OpenFile("./logs/"+logFileName+".log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("Error opening"+logFileName+"file: %v\n", err)
	}
	return f

}

func Run() error {

	//loads configuration for the entire microservice
	conf := confg.NewViperConfig("")

	//creates a logger
	logger := lg.NewLogrusLogger(
		openLogFile("http-debug"),
		openLogFile("kafka-debug"),
		openLogFile("kafka-error"),
		openLogFile("http-error"),
		openLogFile("fatal"),
		openLogFile("info"),
	)
	//creates json validator for http json objects
	v := v.NewPlayGroundValidator()

	//creates an starts mysql database instance
	mysql, err := db.NewMysql(conf, logger)
	if err != nil {
		log.Println(err)
	}

	//creates repositories for database operation
	companyRepo := db.NewCompanyRepo(mysql)
	userRepo := db.NewUserRepo(mysql)

	//creates the domain service for controlling communcation to the doamin
	companyService := comp.NewCompanyService(companyRepo)
	userService := usr.NewUserService(userRepo)

	cmpH := handler.NewCompanyHandler(producerChan, companyService, v, logger)
	userH := handler.NewUserHandler(userService, v, logger)

	//creates and starts kafka producer
	go kafka.NewKafkaProducer(producerChan, conf, logger)
	//creates and starts a consumer listening on kafka topic
	go kafka.NewKafkaConsumer(producerChan, conf, logger, companyService)

	//signal channel to shut down the server using the console
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	//create and start a mux router to serve http requests
	handler.NewMuxRouter(
		cmpH,
		userH,
		conf,
		sigs,
		logger)
	if err != nil {
		return err
	}

	return nil
}
