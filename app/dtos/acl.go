package dtos

import (
	"github.com/Subha-Research/svasthfamily-koham/app/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateACLDTO struct {
	AccessRelationshipID string              `json:"access_relationship_id"`
	HeadUserId           string              `json:"head_family_user_id"`
	ParentuserId         string              `json:"parent_family_user_id"`
	ChildUserID          string              `json:"child_family_user_id"`
	AccessEnum           []float64           `json:"access_enums"`
	Audit                schemas.AuditSchema `json:"audit"`
}

type UpdateACLDTO struct {
	AccessRelationshipID string      `json:"access_relationship_id"`
	HeadUserId           string      `json:"head_family_user_id"`
	ParentuserId         string      `json:"parent_family_user_id"`
	ChildUserID          string      `json:"child_family_user_id"`
	AccessEnum           primitive.A `json:"access_enums"`
	Audit                primitive.M `json:"audit"`
	FamilyID             string      `json:"family_id"`
	FamilyMemberID       string      `json:"family_member_id"`
}

type AccessRelation struct {
	ChildUserID string      `json:"child_family_user_id"`
	AccessEnums interface{} `json:"access_enums"`
}

type AccessRelationshipDTO struct {
	AccessList []AccessRelation `json:"access_list"`
}

func (ar *AccessRelationshipDTO) FormatAllAccessRelationship(access_relations []bson.M) ([]AccessRelation, error) {
	return_data := []AccessRelation{}
	// var results []map[string]interface{}
	for i := 0; i < len(access_relations); i++ {
		ar := AccessRelation{}
		ar.ChildUserID = access_relations[i]["child_family_user_id"].(string)
		// This will become []float64 after code merge from ACLs
		ar.AccessEnums = access_relations[i]["access_enums"]
		return_data = append(return_data, ar)
	}
	return return_data, nil
}
