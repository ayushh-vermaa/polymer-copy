package store

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const DomainCollection = "domain"

type BaseDomain struct {
	Name         string `bson:"name"` // e.g. amazon.com
	CategoryID   int    `bson:"category"`
	CategoryName string `bson:"category"`
}

// Domain represents the structure of a domain document in MongoDB.
type Domain struct {
	*BaseDocument `bson:",inline"`
	*BaseDomain   `bson:",inline"`
}

// CreateDomain creates a Domain document from the given baseDomain.
func CreateDomain(baseDomain *BaseDomain) Domain {
	domain := Domain{
		BaseDocument: &BaseDocument{},
		BaseDomain:   baseDomain,
	}
	domain.SetID()
	return domain
}

// InsertDomain inserts a new Domain document into the MongoDB collection.
func InsertDomain(client *mongo.Client, baseDomain *BaseDomain) (
	*mongo.InsertOneResult, error) {

	domain := CreateDomain(baseDomain)
	store := GetStore(client, DomainCollection)
	return store.InsertDocument(domain)
}

// GetDomainByKey retrieves a Domain document by its unique name.
func GetDomainByName(client *mongo.Client, name string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"name": name}

	store := GetStore(client, DomainCollection)
	count, err := store.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to check if domain name exists: %w", err)
	}
	if count == 0 {
		return nil, fmt.Errorf("domain name does not exist: %s", name)
	}

	var domain Domain
	err = store.Collection.FindOne(ctx, filter).Decode(&domain)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no domain found with name: %s", name)
		}
		return nil, fmt.Errorf("failed to find domain: %w", err)
	}

	return &domain, nil
}
