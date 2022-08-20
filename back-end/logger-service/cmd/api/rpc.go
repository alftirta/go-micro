package main

import (
	"context"
	"log"
	"logger/data"
	"time"
)

// RPCServer is the type for our RPC Server.
// Methods that take this as a receiver are available over RPc, as long as they are exported.
type RPCServer struct{}

// RPCPayload is the type for data we receive from RPC.
type RPCPayload struct {
	Name string
	Data string
}

// LogInfo writes our payload to the database (which we use MongoDB).
func (r *RPCServer) LogInfo(payload RPCPayload, response *string) error {
	collection := client.Database("logs").Collection("logs")
	if _, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}); err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	// response is the message sent back to the RPC caller.
	*response = "Processed payload via RPC:" + payload.Name

	return nil
}
