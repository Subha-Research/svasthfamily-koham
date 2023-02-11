package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Subha-Research/svasthfamily-koham/app/common"
	"github.com/Subha-Research/svasthfamily-koham/app/dtos"
	"github.com/Subha-Research/svasthfamily-koham/app/errors"
	"github.com/Subha-Research/svasthfamily-koham/app/schemas"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenModel struct {
	Collection *mongo.Collection
	Session    *mongo.Session
}

func (tm *TokenModel) InsertToken(f_user_id string, token string, expiry time.Time) (*dtos.CreateTokenResponse, error) {
	time_utils := common.TimeUtil{}
	as := &schemas.AuditSchema{
		CreatedAt: *time_utils.CurrentTimeInUTC(),
		CreatedBy: f_user_id,
		UpdatedAt: *time_utils.CurrentTimeInUTC(),
		UpdatedBy: f_user_id,
	}
	ts := &schemas.TokenSchema{
		TokenID:      uuid.NewString(),
		TokenKey:     token,
		ExpiresAt:    expiry,
		FamilyUserID: f_user_id,
		Audit:        *as,
	}

	res, err := tm.Collection.InsertOne(context.TODO(), ts)
	if err != nil {
		log.Println("Error while inserting token... returning internal server error", err)
		return nil, errors.KohamError("SFKSE-5001")
	}
	log.Println("Inserted token document with ID", res.InsertedID)
	// Build response data
	ctr := dtos.CreateTokenResponse{
		TokenKey:     token,
		TokenExpiry:  expiry,
		FamilyUserID: f_user_id,
		Audit:        *as,
	}
	return &ctr, nil
}

func (tm *TokenModel) GetToken(f_user_id string) (*dtos.GetTokenResponse, error) {
	var result bson.M
	err := tm.Collection.FindOne(
		context.TODO(),
		bson.D{
			{Key: "family_user_id", Value: f_user_id},
			{Key: "expires_at", Value: bson.D{{Key: "$gte", Value: primitive.NewDateTimeFromTime(time.Now())}}},
		},
	).Decode(&result)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		err_data := map[string]string{
			"f_user_id": f_user_id,
		}
		if err == mongo.ErrNoDocuments {
			return nil, errors.KohamError("SFKSE-4007", err_data)
		}
		return nil, err
	}
	fmt.Printf("Token document %v", result)
	gtr := dtos.GetTokenResponse{}
	gtr.TokenKey = result["token_key"].(string)
	gtr.TokenExpiry = result["expires_at"].(primitive.DateTime)
	gtr.FamilyUserID = result["family_user_id"].(string)

	return &gtr, nil
}

func (tm *TokenModel) DeleteToken(f_user_id *string, token *string) error {
	_, err := tm.Collection.DeleteOne(context.TODO(), bson.D{{Key: "family_user_id", Value: f_user_id}, {Key: "token_key", Value: token}})
	if err != nil {
		log.Println("Error in deleting token, returning internal service error", err)
		return errors.KohamError("SFKSE-5001")
	}
	return nil
}
