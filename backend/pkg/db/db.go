package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection
var ctx = context.TODO()

// Message matches the JSON we send and the BSON we store in Mongo
type Message struct {
	Type     string `json:"type" bson:"type"`
	Username string `json:"username" bson:"username"`
	Body     string `json:"body" bson:"body"`
	Time     string `json:"time" bson:"time"`
}

func InitDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Check if MongoDB is running! Error:", err)
	}

	fmt.Println("✅ MongoDB Connected")
	Collection = client.Database("chat_app").Collection("messages")
}

// InsertMessage saves the message to Mongo
func InsertMessage(msg Message) error {
	_, err := Collection.InsertOne(ctx, msg)
	return err
}

func GetHistory() ([]Message, error) {
	var messages []Message

	// We want to fetch messages. You can add .SetLimit(50) if the chat gets huge.
	cursor, err := Collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var msg Message
		if err := cursor.Decode(&msg); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

func ClearAllMessages() error {
	// Passing an empty filter {} tells Mongo to match (and delete) everything
	_, err := Collection.DeleteMany(ctx, map[string]interface{}{})
	return err
}
