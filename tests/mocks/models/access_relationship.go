package models_mock

import "go.mongodb.org/mongo-driver/bson"

type AccessRelationshipModelMock struct {
}

func (armm *AccessRelationshipModelMock) GetAllAccessRelationship(f_user_id string) ([]bson.M, error) {
	//
	var results []bson.M
	single_access_relation := bson.M{
		"access_relationship_id": "35aef7b1-f634-4b1a-9ea7-67c41d0fbcf1",
		"parent_family_user_id":  "8204a616-2131-4a64-97d0-ae3f2b9211be",
		"child_family_user_id":   "8204a616-2131-4a64-97d0-ae3f2b9211be",
		"access_enums":           []float64{101, 102, 103, 104, 105, 106, 107, 108, 109},
	}
	results = append(results, single_access_relation)
	return results, nil
}
