package models

// import (
// 	"context"
// 	"log"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type LogEntryModelInterface interface {
	
// }

// type LogEntry struct {
// 	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
// 	Name      string    `bson:"name" json:"name"`
// 	Data      string    `bson:"data" json:"data"`
// 	CreatedAt time.Time `bson:"created_at" json:"created_at"`
// 	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
// }

// type LogEntryModel struct {
// 	Client *mongo.Client
// }

// func (m *LogEntryModel) Insert(ctx context.Context, name, data string) error {
// 	collection := m.Client.Database("logs").Collection("logs")

// 	_, err := collection.InsertOne(ctx, LogEntry{
// 		Name:      name,
// 		Data: 	   data,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	})
// 	if err != nil {
// 		log.Println("Error inserting into logs:", err)
// 		return err
// 	}

// 	return nil
// }

// func (m *LogEntryModel) All(ctx context.Context) ([]*LogEntry, error) {
// 	collection := m.Client.Database("logs").Collection("logs")

// 	opts := options.Find()
// 	opts.SetSort(bson.D{{"created_at", -1}})

// 	cursor, err := collection.Find(ctx, bson.D{}, opts)
// 	if err != nil {
// 		log.Println("Finding all docs error:", err)
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	var logs []*LogEntry

// 	for cursor.Next(ctx) {
// 		var item LogEntry

// 		err := cursor.Decode(&item)
// 		if err != nil {
// 			log.Print("Error decoding log into slice:", err)
// 			return nil, err
// 		} else {
// 			logs = append(logs, &item)
// 		}
// 	}

// 	return logs, nil
// }

// func (m *LogEntryModel) GetOne(id string) (*LogEntry, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
// 	defer cancel()

// 	collection := m.Client.Database("logs").Collection("logs")

// 	docID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var entry LogEntry
// 	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&entry)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &entry, nil
// }

// func (m *LogEntryModel) DropCollection(ctx context.Context) error {

// 	collection := m.Client.Database("logs").Collection("logs")

// 	if err := collection.Drop(ctx); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (m *LogEntryModel) Update(ctx context.Context, id string, name, data string) (*mongo.UpdateResult, error) {
// 	collection := m.Client.Database("logs").Collection("logs")

// 	docID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	result, err := collection.UpdateOne(
// 		ctx,
// 		bson.M{"_id": docID},
// 		bson.D{
// 			{"$set", bson.D{
// 				{"name", name},
// 				{"data", data},
// 				{"updated_at", time.Now()},
// 			}},
// 		},
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }