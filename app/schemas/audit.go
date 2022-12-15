package sf_schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type AuditSchema struct {
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	CreatedBy string    `bson:"created_by" json:"created_by"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	UpdatedBy string    `bson:"updated_by" json:"updated_by"`
}

func (as *AuditSchema) MarshalBSON() ([]byte, error) {
	if as.CreatedAt.IsZero() {
		as.CreatedAt = time.Now()
	}
	as.UpdatedAt = time.Now()
	// type my AuditSchema
	type audit AuditSchema
	return bson.Marshal((*audit)(as))
}
