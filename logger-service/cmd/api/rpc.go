package main

import (
	"context"
	"log"
	"log-service/data"
	"time"
)

// 1. Create Server
// 2. Create Payload

type RPCServer struct {
}

type RPCPayload struct {
	Name string
	Data string
}

// Writes our payload to Mongo
func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	*resp = "Processed payload via RPC:" + payload.Name
	return nil
}
