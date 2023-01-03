package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Subha-Research/svasthfamily-koham/app/dto"
	"github.com/Subha-Research/svasthfamily-koham/app/errors"
	"github.com/Subha-Research/svasthfamily-koham/app/schemas"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenModel struct {
	Collection *mongo.Collection
	Session    *mongo.Session
}

func (tm *TokenModel) InsertToken(f_user_id string, token string, expiry time.Time) (*dto.CreateTokenResponse, error) {
	as := &schemas.AuditSchema{
		CreatedAt: time.Now(),
		CreatedBy: f_user_id,
		UpdatedAt: time.Now(),
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
		return nil, errors.KohamError("KSE-5001")
	}
	log.Println("Inserted token document with ID", res.InsertedID)
	// Build response data
	ctr := dto.CreateTokenResponse{}
	ctr.TokenKey = token
	ctr.TokenExpiry = expiry
	ctr.FamilyUserID = f_user_id
	ctr.Audit = *as
	return &ctr, nil
}

func (tm *TokenModel) GetToken(f_user_id string) (*dto.GetTokenResponse, error) {
	var result bson.M
	err := tm.Collection.FindOne(
		context.TODO(),
		bson.D{{Key: "family_user_id", Value: f_user_id}},
	).Decode(&result)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if err == mongo.ErrNoDocuments {
			return nil, errors.KohamError("KSE-4007")
		}
		return nil, err
	}
	fmt.Printf("Token document %v", result)
	gtr := dto.GetTokenResponse{}
	gtr.TokenKey = result["token_key"].(string)
	gtr.TokenExpiry = result["expires_at"].(time.Time)
	gtr.FamilyUserID = result["family_user_id"].(string)

	return &gtr, nil
}

func (tm *TokenModel) DeleteToken(f_user_id *string, token *string) error {
	_, err := tm.Collection.DeleteOne(context.TODO(), bson.D{{Key: "family_user_id", Value: f_user_id}, {Key: "token_key", Value: token}})
	if err != nil {
		return errors.KohamError("KSE-5001")
	}
	return nil
}
