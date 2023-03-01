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

	"github.com/gorilla/mux"
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

// //creates and start a new consumer
// func NewKafkaConsumer(cmpConsumerChan chan Message, conf confg.Configuration, log lgg.Logger, cs cmp.Service) {
// 	mq := &messageQueue{
// 		service: cs,
// 		lg:      log,
// 		conChan: cmpConsumerChan,
// 	}
// 	c, err := kafka.NewConsumer(&kafka.ConfigMap{
// 		"bootstrap.servers": conf.KafkaConfig.ServerIp,
// 		"group.id":          conf.KafkaConfig.GroupId,
// 		"auto.offset.reset": conf.KafkaConfig.AutoOffsetReset,
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer c.Close()
// 	c.Subscribe(conf.KafkaConfig.Topic, nil)
// 	for {
// 		msg, err := c.ReadMessage(-1)
// 		if err != nil {
// 			//lgg.Logger.MqError("Consumer", msg, err)
// 			fmt.Printf("Consumer error: #{err} (#{msg}\n)")
// 		}
// 		go mq.processMessage(msg)
// 	}
// }

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
	switch m.Request.Method {
	case http.MethodPost:
		mq.addCompany(m)
	case http.MethodGet:
		mq.getCompany(m)
	case http.MethodDelete:
		mq.deleteCompany(m)
	case http.MethodPatch:
		mq.updateCompany(m)
	}
}

//function to add company to db and return http response
func (mq *messageQueue) addCompany(m Message) {
	hops, err := mq.service.AddCompany(m.Request.Context(), cmp.Company(m.Company))
	if err != nil {
		mq.lg.HttpError(m.Request, "CreateCompany", err)
		writeResponse(m.ResponseWriter, http.StatusInternalServerError, nil, err)
		return
	}
	writeResponse(m.ResponseWriter, http.StatusCreated, hops, nil)
}

//function to get company from db and return http response
func (mq *messageQueue) getCompany(m Message) {
	vars := mux.Vars(m.Request)
	compId := vars["companyId"]

	comp, err := mq.service.GetCompany(m.Request.Context(), compId)
	if err != nil {
		mq.lg.HttpError(m.Request, "GetCompany", err)
		writeResponse(m.ResponseWriter, http.StatusInternalServerError, nil, err)
		return
	}

	writeResponse(m.ResponseWriter, http.StatusOK, comp, nil)
}

//function to update company info
func (mq *messageQueue) updateCompany(m Message) {
	cmp, err := mq.service.ModifyCompany(m.Request.Context(), cmp.Company(m.Company), m.Id)
	if err != nil {
		mq.lg.HttpError(m.Request, "updateCompany", err)
		writeResponse(m.ResponseWriter, http.StatusInternalServerError, nil, err)
		return
	}
	writeResponse(m.ResponseWriter, http.StatusCreated, cmp, nil)
}

//function to delete company
func (mq *messageQueue) deleteCompany(m Message) {

	compId, err := mq.service.RemoveCompany(m.Request.Context(), m.Id)
	if err != nil {
		mq.lg.HttpError(m.Request, "DeleteCompany", err)
		writeResponse(m.ResponseWriter, http.StatusInternalServerError, nil, err)
		return
	}
	writeResponse(m.ResponseWriter, http.StatusOK, compId, nil)
}

//helper fucntion for writing http reponse
func writeResponse(w http.ResponseWriter, status int, data interface{}, err error) {
	w.WriteHeader(status)
	if err != nil {
		resp := ErrorResp{
			Error: fmt.Sprint(err),
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			fmt.Fprintf(w, "error encoding resp %v:%s", resp, err)
		}
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Fprintf(w, "error encoding resp %v:%s", data, err)
	}
}

// func (mq *messageQueue) addCompany(m Message) {
// 	//Get the Databank location from the service layer
// 	ctx := context.Background()
// 	var respWriter http.ResponseWriter
// 	var req *http.Request

// 	if m.Request != nil {
// 		ctx = m.Request.Context()
// 		respWriter = m.ResponseWriter
// 		req = m.Request
// 	}

// 	company, err := mq.service.AddCompany(ctx, cmp.Company(m.company))
// 	if err != nil {
// 		mq.lg.HttpError(req, "CreateCompany", err)
// 		writeResponse(respWriter, http.StatusInternalServerError, nil, err)
// 		return
// 	}
// 	//write response wiht datbank location to http response
// 	writeResponse(respWriter, http.StatusCreated, company, nil)
// }
