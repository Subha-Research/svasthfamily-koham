package dto

import (
	"go.mongodb.org/mongo-driver/bson"
)

type AccessRelation struct {
	ChildMemberID string    `json:"child_member_id"`
	AccessEnums   []float64 `json:"access_enums"`
}

type AccessRelationshipDTO struct {
	AccessList []AccessRelation `json:"access_list"`
}

func (ar *AccessRelationshipDTO) FormatAllAccessRelationship(access_relations []bson.M) (AccessRelationshipDTO, error) {
	dto := AccessRelationshipDTO{}
	// var results []map[string]interface{}
	for i := 0; i < len(access_relations); i++ {
		ar := AccessRelation{}
		ar.ChildMemberID = access_relations[i]["child_member_user_id"].(string)
		ar.AccessEnums = access_relations[i]["access_enums"].([]float64)
		dto.AccessList = append(dto.AccessList, ar)
	}
	return dto, nil
}
