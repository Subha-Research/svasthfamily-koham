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

type RoleModel struct {
	Collection *mongo.Collection
	Session    *mongo.Session
}

func (rm *RoleModel) GetRole(role_enum int, role_key string) (bson.M, error) {
	var result bson.M
	err := rm.Collection.FindOne(
		context.TODO(),
		bson.D{{Key: "role_enum", Value: role_enum}, {Key: "role_key", Value: role_key}},
	).Decode(&result)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	fmt.Printf("found document %v", result)
	return result, nil
}

func (rm *RoleModel) InsertAllRoles() error {
	// Colllection variable is set via Dependency injection from app file
	var role_docs []interface{}
	role_map := constants.ROLES
	for k, v := range role_map {
		// Check first if the same role already exists
		// If exist then do not insert that
		doc, err := rm.GetRole(k, v)
		if doc != nil {
			continue
		} else if doc == nil && err == nil {
			role := &schemas.RoleSchema{
				RoleID:   uuid.NewString(),
				RoleEnum: k,
				RoleKey:  v,
				IsActive: true,
				IsDelete: false,
			}
			// Convert role struct to interface
			// roleInterface = role
			role_docs = append(role_docs, role)
		} else {
			log.Fatal("Error in getting roles, stopping server", err)
		}
	}

	// Call insert many of mongo
	if len(role_docs) > 0 {
		opts := options.InsertMany().SetOrdered(false)
		res, err := rm.Collection.InsertMany(context.TODO(), role_docs, opts)
		if err != nil {
			log.Fatal("Error in inserting role. Stopping server", err)
			return err
		}
		fmt.Printf("Inserted documents with IDs %v\n", res.InsertedIDs)
	}
	return nil
}
