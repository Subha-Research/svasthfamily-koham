package models

import (
	"context"
	"log"
	"time"

	"github.com/Subha-Research/svasthfamily-koham/app/dto"
	"github.com/Subha-Research/svasthfamily-koham/app/errors"
	"github.com/Subha-Research/svasthfamily-koham/app/schemas"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenModel struct {
	Collection *mongo.Collection
	Session    *mongo.Session
	// InsertReturn dto.CreateTokenResponse
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
	return &ctr, nil
}
