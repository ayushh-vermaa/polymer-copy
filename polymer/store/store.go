package store

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DatabaseName = "polymer"

// MongoStore encapsulates the MongoDB client and collection.
type MongoStore struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

type BaseDocument struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty"`
}

func (b *BaseDocument) SetID() {
	b.ID = primitive.NewObjectID()
}

func (b *BaseDocument) SetCreatedAt() {
	b.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
}

type Document interface {
	SetID()
	SetCreatedAt()
}

// ConnectMongoDB connects to the MongoDB cluster and returns a client.
func ConnectMongoDB() (*mongo.Client, error) {
	mongoURI := "mongodb+srv://ayushpverma:polymer@polymer-cluster.gcwlg.mongodb.net/?retryWrites=true&w=majority&appName=polymer-cluster"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successful database connection")

	return client, nil
}

func GetStore(client *mongo.Client, collectionName string) MongoStore {
	collection := client.Database(DatabaseName).Collection(collectionName)
	return MongoStore{client, collection}
}

// InsertDocument inserts a document into the MongoDB collection.
func (store *MongoStore) InsertDocument(document Document) (
	*mongo.InsertOneResult, error) {

	document.SetCreatedAt()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := store.Collection.InsertOne(ctx, document)
	if err != nil {
		return nil, fmt.Errorf("failed to insert document: %w", err)
	}

	return result, nil
}
