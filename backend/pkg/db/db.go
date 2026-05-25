package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection
var ctx = context.TODO()

type Message struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Type     string             `json:"type" bson:"type"`
	Username string             `json:"username" bson:"username"`
	Body     string             `json:"body" bson:"body"`
	Time     string             `json:"time" bson:"time"`
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

// 🚀 UPDATED: Returns the message with the generated ID
func InsertMessage(msg Message) (Message, error) {
	result, err := Collection.InsertOne(ctx, msg)
	if err != nil {
		return msg, err
	}
	// Inject the new ID into the struct before returning
	if oID, ok := result.InsertedID.(primitive.ObjectID); ok {
		msg.ID = oID
	}
	return msg, nil
}

func DeleteMessageByID(idStr string) error {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return err
	}
	_, err = Collection.DeleteOne(ctx, map[string]interface{}{"_id": id})
	return err
}

func GetHistory() ([]Message, error) {
	var messages []Message
	cursor, err := Collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var msg Message
		if err := cursor.Decode(&msg); err != nil {
			continue
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

func ClearAllMessages() error {
	_, err := Collection.DeleteMany(ctx, map[string]interface{}{})
	return err
}
