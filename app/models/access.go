package models

import (
	"context"
	"fmt"
	"log"

	"github.com/Subha-Research/svasthfamily-koham/app/constants"
	schemas "github.com/Subha-Research/svasthfamily-koham/app/schemas"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AccessModel struct {
	Collection *mongo.Collection
	Session    *mongo.Session
}

func (am *AccessModel) GetAccess(access_enum float64, access_key string) (bson.M, error) {
	var result bson.M
	err := am.Collection.FindOne(
		context.TODO(),
		bson.D{{Key: "access_enum", Value: access_enum}, {Key: "access_key", Value: access_key}},
	).Decode(&result)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	fmt.Printf("found access document %v", result)
	return result, nil
}

func (am *AccessModel) InsertAllAccesses() error {
	// Collection variable is set via Dependency injection from app file
	var access_docs []interface{}
	access_map := constants.HEAD_DEFAULT_ACCESS
	for k, v := range access_map {
		// Check first if the same role already exists
		// If exist then do not insert that
		doc, err := am.GetAccess(k, v)
		if doc != nil {
			continue
		} else if doc == nil && err == nil {
			role := &schemas.Access{
				AccessID:   uuid.NewString(),
				AccessEnum: k,
				AccessKey:  v,
				IsActive:   true,
				IsDelete:   false,
			}
			// Convert role struct to interface
			// roleInterface = role
			access_docs = append(access_docs, role)
		} else {
			log.Fatal("Error in getting access, stopping server", err)
		}
	}

	// Call insert many of mongo
	if len(access_docs) > 0 {
		opts := options.InsertMany().SetOrdered(false)
		res, err := am.Collection.InsertMany(context.TODO(), access_docs, opts)
		if err != nil {
			log.Fatal("Error in inserting access. Stopping server", err)
			return err
		}
		fmt.Printf("Inserted access documents with IDs %v\n", res.InsertedIDs)
	}
	return nil
}
