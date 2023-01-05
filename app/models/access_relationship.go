package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Subha-Research/svasthfamily-koham/app/common"
	"github.com/Subha-Research/svasthfamily-koham/app/constants"
	"github.com/Subha-Research/svasthfamily-koham/app/errors"
	schemas "github.com/Subha-Research/svasthfamily-koham/app/schemas"
	validators "github.com/Subha-Research/svasthfamily-koham/app/validators"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/exp/maps"
)

type AccessRelationshipModel struct {
	Collection *mongo.Collection
	Session    *mongo.Session
}

type UserIDs struct {
	HeadUserId   string
	ChildUserId  string
	ParentUserId string
}

func (arm *AccessRelationshipModel) GetAllAccessRelationship(f_user_id string) ([]bson.M, error) {
	var results []bson.M
	cursor, err := arm.Collection.Find(context.TODO(), bson.D{{Key: "parent_family_user_id", Value: f_user_id}})
	if err != nil {
		log.Println("Error while getting all access relationship", err)
		return nil, errors.KohamError("KSE-5001")
	}

	// Get a list of all returned documents and print them out.
	// See the mongo.Cursor documentation for more examples of using cursors.
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Println("Error while getting all access relationship", err)
		return nil, errors.KohamError("KSE-5001")
	}
	if len(results) == 0 {
		error_data := map[string]string{
			"id": f_user_id,
		}
		log.Printf("No access relationship found for id %s", f_user_id)
		return nil, errors.KohamError("KSE-4012", error_data)
	}
	return results, nil
}

func (arm *AccessRelationshipModel) GetAccessRelationship(f_head_user_id *string, f_parent_user_id string, f_child_user_id string) (bson.M, error) {
	var filter = bson.D{{Key: "parent_family_user_id", Value: f_parent_user_id}, {Key: "child_family_user_id", Value: f_child_user_id}}
	if f_head_user_id != nil {
		filter = append(filter, bson.E{Key: "head_family_user_id", Value: f_head_user_id})
	}
	var result bson.M
	err := arm.Collection.FindOne(
		context.TODO(),
		filter,
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

func (arm *AccessRelationshipModel) InsertAllAccessRelationship(f_head_user_id string, is_head_head bool, rb validators.ACLPostBody) (bson.M, error) {
	// time_util := common.TimeUtil{}
	var access_list_docs []interface{}
	access_list := rb.AccessList
	for i := 0; i < len(access_list); i++ {
		var access_enums = access_list[i].AccessEnums
		var err error
		var access_relation_parent_child *schemas.AccessRelationshipSchema
		// If access already created in parent user id and child user id
		doc, _ := arm.GetAccessRelationship(nil, rb.ParentUserID, access_list[i].ChildUserId)
		if doc != nil {
			return doc, errors.KohamError("KSE-4009")
		}

		u_ids := UserIDs{
			HeadUserId:   f_head_user_id,
			ChildUserId:  access_list[i].ChildUserId,
			ParentUserId: rb.ParentUserID,
		}
		if is_head_head && access_enums == nil {
			access_relation_parent_child, err = arm.getAccessRelation(u_ids, "HEAD_HEAD", *rb.IsParentHead, access_enums)
		} else {
			access_relation_parent_child, err = arm.getAccessRelation(u_ids, "PARENT_CHILD", *rb.IsParentHead, access_enums)
		}
		access_list_docs = append(access_list_docs, access_relation_parent_child)
		if err != nil {
			return nil, err
		}
		if !is_head_head {
			//Inserting child child relation
			u_ids.ParentUserId = access_list[i].ChildUserId
			access_relation_child_child, err := arm.getAccessRelation(u_ids, "CHILD_CHILD", *rb.IsParentHead, nil)
			if err != nil {
				return nil, err
			}
			access_list_docs = append(access_list_docs, access_relation_child_child)
			if !*rb.IsParentHead {
				// Head child relation
				u_ids.ParentUserId = f_head_user_id
				access_relation_head_child, err := arm.getAccessRelation(u_ids, "HEAD_CHILD", *rb.IsParentHead, nil)
				if err != nil {
					return nil, err
				}
				access_list_docs = append(access_list_docs, access_relation_head_child)
			}
		}
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
	time_util := common.TimeUtil{}
	access_list := rb.Access
	access_relation := access_list.AccessEnums

	filter := bson.D{{Key: "child_family_user_id", Value: rb.Access.ChildUserId}, {Key: "parent_family_user_id", Value: rb.ParentUserID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "access_enums", Value: access_relation}, {Key: "audit.updated_at", Value: *time_util.CurrentTimeInUTC()}, {Key: "audit.updated_by", Value: f_head_user_id}}}}
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

func (arm *AccessRelationshipModel) getSchema(ids UserIDs,
	relation_type string, is_parent_head bool, access []float64) (*schemas.AccessRelationshipSchema, error) {

	access_relation := &schemas.AccessRelationshipSchema{
		AccessRelationshipID: uuid.NewString(),
		HeadFamilyUserID:     ids.HeadUserId,
		ChildFamilyUserID:    ids.ChildUserId,
		ParentFamilyUserID:   ids.ParentUserId,
		RelationshipType:     relation_type,
		IsParentHead:         is_parent_head,
		AccessEnum:           access,
		IsDelete:             false,
		Audit: schemas.AuditSchema{
			CreatedAt: time.Now(),
			CreatedBy: ids.HeadUserId,
			UpdatedAt: time.Now(),
			UpdatedBy: ids.HeadUserId,
		},
	}
	return access_relation, nil
}

func (arm *AccessRelationshipModel) getAccessRelation(ids UserIDs,
	relation_type string, is_parent_head bool, access []float64) (*schemas.AccessRelationshipSchema, error) {

	var default_access = maps.Keys(constants.HEAD_DEFAULT_ACCESS)
	var access_relation *schemas.AccessRelationshipSchema
	var err error

	switch relation_type {
	case "HEAD_HEAD":
		access_relation, err = arm.getSchema(ids, relation_type, is_parent_head, default_access)
	case "PARENT_CHILD":
		if access != nil {
			default_access = access
		} else {
			default_access = maps.Keys(constants.CHILD_DEFAULT_ACCESS)
		}
		access_relation, err = arm.getSchema(ids, relation_type, is_parent_head, default_access)
	case "HEAD_CHILD":
		default_access = maps.Keys(constants.CHILD_DEFAULT_ACCESS)
		access_relation, err = arm.getSchema(ids, relation_type, is_parent_head, default_access)
	case "CHILD_CHILD":
		default_access = maps.Keys(constants.CHILD_DEFAULT_ACCESS)
		access_relation, err = arm.getSchema(ids, relation_type, is_parent_head, default_access)
	default:
		return nil, errors.KohamError("KSE-4013")
	}
	return access_relation, err
}
