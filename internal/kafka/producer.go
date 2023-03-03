package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	confg "xm/internal/config"

	"github.com/segmentio/kafka-go"

	lgg "xm/internal/logger"
)

func NewKafkaProducer(cmpProducerChan chan Message, conf confg.Configuration, log lgg.Logger) {

	w := &kafka.Writer{
		Addr:     kafka.TCP(conf.KafkaConfig.ServerIp),
		Topic:    conf.KafkaConfig.Topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer w.Close()
	//comunicating by sharing memoryreceiving messages from http request
	for {
		m := <-cmpProducerChan
		//fmt.Printf("prod: %+v\n", m)
		encoded, err := json.Marshal(m)
		if err != nil {
			panic(err)
		}
		err = w.WriteMessages(context.Background(), kafka.Message{Value: encoded})
		if err != nil {
			panic(err)
		}
		//fmt.Printf("Sent message: %s\n", m)
		fmt.Printf("Producer sent: %d bytes\n", w.Stats().Bytes)
	}

}

// //creates and starts a new kafka producer
// func NewKafkaProducer(cmpProducerChan chan Message, conf confg.Configuration, log lgg.Logger) {
// 	cfg := kafka.ConfigMap{"bootstrap.servers": conf.KafkaConfig.ServerIp}
// 	p, err := kafka.NewProducer(&cfg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer p.Close()

// 	topic := conf.KafkaConfig.Topic
// 	m := <-cmpProducerChan

// 	encoded, err := json.Marshal(m)
// 	if err != nil {
// 		panic(err)
// 	}

// 	p.Produce(&kafka.Message{
// 		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartisionA},
// 		Value:          encoded,
// 	}, nil)
// 	p.Flush(1000)
// }

// p.Produce(&kafka.Message{
// 	TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartisionA},
// 	Value:          encoded,
// }, nil)
