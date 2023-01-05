package dto

import (
	"go.mongodb.org/mongo-driver/bson"
)

type AccessRelation struct {
	ChildUserID string      `json:"child_user_id"`
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
