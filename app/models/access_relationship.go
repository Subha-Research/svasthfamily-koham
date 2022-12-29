package models

import (
	"context"
	"fmt"
	"log"
	"time"

	enums "github.com/Subha-Research/svasthfamily-koham/app/enums"
	"github.com/Subha-Research/svasthfamily-koham/app/errors"
	sf_schemas "github.com/Subha-Research/svasthfamily-koham/app/schemas"
	validators "github.com/Subha-Research/svasthfamily-koham/app/validators"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AccessRelationshipModel struct {
	Collection *mongo.Collection
	Session    *mongo.Session
}

func (arm *AccessRelationshipModel) GetAccessRelationship(f_parent_user_id string, f_child_user_id string) (bson.M, error) {
	var result bson.M
	err := arm.Collection.FindOne(
		context.TODO(),
		bson.D{{Key: "parent_family_user_id", Value: f_parent_user_id}, {Key: "child_family_user_id", Value: f_child_user_id}},
	).Decode(&result)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if err == mongo.ErrNoDocuments {
			return nil, errors.KohamError("KSE-4007")
		}
		return nil, err
	}
	fmt.Printf("User ID matched %v", result)
	return result, nil
}

func (arm *AccessRelationshipModel) InsertAllAccessRelationship(f_head_user_id string, rb validators.ACLPostBody) (bson.M, error) {
	// Colllection variable is set via Dependency injection from app file
	var access_list_docs []interface{}
	access_list := rb.AccessList
	default_child_access := []float64{102, 103, 105, 106, 107, 106, 108}
	for i := 0; i < len(access_list); i++ {
		doc, _ := arm.GetAccessRelationship(rb.ParentMemberID, access_list[i].ChildMemberId)
		if doc != nil {
			return doc, errors.KohamError("KSE-4009")
		}
		child_parent_access_relation := &sf_schemas.AccessRelationshipSchema{
			AccessRelationshipID: uuid.NewString(),
			ChildFamilyUserID:    access_list[i].ChildMemberId,
			ParentFamilyUserID:   rb.ParentMemberID,
			AccessEnum:           access_list[i].AccessEnums,
			IsDelete:             false,
			Audit: sf_schemas.AuditSchema{
				CreatedAt: time.Now(),
				CreatedBy: f_head_user_id,
				UpdatedAt: time.Now(),
				UpdatedBy: f_head_user_id,
			},
		}
		if enums.Roles[rb.RoleEnum] != "FAMILY_HEAD" {
			child_child_access_relation := &sf_schemas.AccessRelationshipSchema{
				AccessRelationshipID: uuid.NewString(),
				ChildFamilyUserID:    access_list[i].ChildMemberId,
				ParentFamilyUserID:   access_list[i].ChildMemberId,
				AccessEnum:           default_child_access,
				IsDelete:             false,
				Audit: sf_schemas.AuditSchema{
					CreatedAt: time.Now(),
					CreatedBy: f_head_user_id,
					UpdatedAt: time.Now(),
					UpdatedBy: f_head_user_id,
				},
			}
			access_list_docs = append(access_list_docs, child_child_access_relation)
		}
		access_list_docs = append(access_list_docs, child_parent_access_relation)
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

func (arm *AccessRelationshipModel) UpdateAccessRelationship(f_head_user_id string, rb validators.ACLPutBody) (bson.M, error) {
	// Colllection variable is set via Dependency injection from app file
	access_list := rb.Access
	access_relation := access_list.AccessEnums

	filter := bson.D{{Key: "child_family_user_id", Value: rb.Access.ChildMemberId}, {Key: "parent_family_user_id", Value: rb.ParentMemberID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "access_enums", Value: access_relation}, {Key: "audit.updated_at", Value: time.Now()}, {Key: "audit.updated_by", Value: f_head_user_id}}}}
	var updatedDocument bson.M
	err := arm.Collection.FindOneAndUpdate(
		context.TODO(),
		filter,
		update,
	).Decode(&updatedDocument)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		log.Fatal(err)
	}
	log.Println("updated document", updatedDocument)
	return updatedDocument, nil
}
