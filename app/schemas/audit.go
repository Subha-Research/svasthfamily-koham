package schemas

import (
	"time"

	"github.com/Subha-Research/svasthfamily-koham/app/common"
	"go.mongodb.org/mongo-driver/bson"
)

type AuditSchema struct {
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	CreatedBy string    `bson:"created_by" json:"created_by"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	UpdatedBy string    `bson:"updated_by" json:"updated_by"`
}

func (as *AuditSchema) MarshalBSON() ([]byte, error) {
	time_util := common.TimeUtil{}
	if as.CreatedAt.IsZero() {
		as.CreatedAt = *time_util.CurrentTimeInUTC()
	}
	as.UpdatedAt = *time_util.CurrentTimeInUTC()
	// type my AuditSchema
	type audit AuditSchema
	return bson.Marshal((*audit)(as))
}
