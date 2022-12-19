package sf_models

import (
	"context"
	"fmt"
	"log"
	"time"

	sf_schemas "github.com/Subha-Research/svasthfamily-koham/app/schemas"
	validators "github.com/Subha-Research/svasthfamily-koham/app/validators"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SFAccessRelationshipModel struct {
	Collection *mongo.Collection
	Session    *mongo.Session
}

// func (arm *SFAccessRelationshipModel) CreateSFAccessRelationship(rb validators.ACLPostBody) (bson.M, error) {

func (arm *SFAccessRelationshipModel) InsertAllSFAccessRelationship(f_head_user_id string, rb validators.ACLPostBody) (bson.M, error) {
	// Colllection variable is set via Dependency injection from app file
	var access_list_docs []interface{}
	access_list := rb.ChildMemberAccessList
	for i := 0; i < len(access_list); i++ {
		// Check first if the same role already exists
		// If exist then do not insert that
		// doc, err := arm.GetRole(i, access_list[i])

		log.Println("***** Access list", access_list[i])

		access_relation := &sf_schemas.AccessRelationshipSchema{
			AccessRelationshipID: uuid.NewString(),
			ChildFamilyUserID:    access_list[i]["child_member_id"].(string),
			ParentFamilyUserID:   rb.ParentMemberID,
			AccessEnum:           access_list[i]["access_enums"].([]interface{}),
			IsDelete:             false,
			Audit: sf_schemas.AuditSchema{
				CreatedAt: time.Now(),
				CreatedBy: f_head_user_id,
				UpdatedAt: time.Now(),
				UpdatedBy: f_head_user_id,
			},
		}
		access_list_docs = append(access_list_docs, access_relation)
	}

	// Call insert many of mongo
	if len(access_list_docs) > 0 {
		opts := options.InsertMany().SetOrdered(false)
		res, err := arm.Collection.InsertMany(context.TODO(), access_list_docs, opts)
		if err != nil {
			log.Println("Error in inserting access relation", err)
			return nil, err
		}
		fmt.Printf("Inserted documents with IDs %v\n", res.InsertedIDs)
	}
	return nil, nil
}
