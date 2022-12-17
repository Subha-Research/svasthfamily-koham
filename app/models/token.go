package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Subha-Research/svasthfamily-koham/app/schemas"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenModel struct {
	Collection *mongo.Collection
	Session    *mongo.Session
}

func (tm *TokenModel) InsertToken(f_user_id string, token string, expiry time.Time) error {

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
		log.Fatal(err)
	}
	fmt.Printf("inserted document with ID %v\n", res.InsertedID)

	return nil
}
