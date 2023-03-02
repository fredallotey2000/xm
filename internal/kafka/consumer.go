package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	confg "xm/internal/config"

	//router "xm/internal/http/handlers"

	lgg "xm/internal/logger"

	cmp "xm/pkg/company"

	"github.com/segmentio/kafka-go"
)

type ErrorResp struct {
	Error string
}

func NewKafkaConsumer(cmpConsumerChan chan Message, conf confg.Configuration, log lgg.Logger, cs cmp.Service) {
	mq := &messageQueue{
		service: cs,
		lg:      log,
		conChan: cmpConsumerChan,
	}
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{conf.KafkaConfig.ServerIp},
		Partition: 0,
		Topic:     conf.KafkaConfig.Topic,
		GroupID:   conf.KafkaConfig.GroupId,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})

	defer r.Close()

	for {
		m, err := r.FetchMessage(context.Background())
		if err != nil {
			break
		}
		if err := r.CommitMessages(context.Background(), m); err != nil {
			panic(err)
		}
		//lg.
		//fmt.Printf("Topic %s msg: %s\n", m.Topic, m.Value)
		go mq.processMessage(m)

	}
}

//Message queue and all the dependencies required for MQ operations
type messageQueue struct {
	service cmp.Service
	lg      lgg.Logger
	conChan chan Message
}

//go routine to process read messages off topic
func (mq *messageQueue) processMessage(msg kafka.Message) {
	var m Message
	err := json.Unmarshal(msg.Value, &m)
	if err == nil {
		//lgg.Logger.MqError("Consumer", msg.TopicPartition, err)
		//fmt.Printf("Message on %s: %s\n", msg.TopicPartition, m)
	} else {
		// The client will automatically try to recover from all errors.
		fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		//lgg.Logger.MqError("Consumer", msg, err)
	}
	switch m.RequestMethod {
	case http.MethodDelete:
		mq.deleteCompany(m)
	case http.MethodPatch:
		mq.updateCompany(m)
	}
}

//function to update company info
func (mq *messageQueue) updateCompany(m Message) {
	_, err := mq.service.ModifyCompany(context.Background(), cmp.Company(m.Company), m.Id)
	if err != nil {
		mq.lg.MqError("Patching", "", err)
	}
}

//function to delete company
func (mq *messageQueue) deleteCompany(m Message) {

	_, err := mq.service.RemoveCompany(context.Background(), m.Id)
	if err != nil {
		mq.lg.MqError("Deletion", "", err)
	}

}
