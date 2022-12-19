package services

import (
	"log"

	sf_models "github.com/Subha-Research/svasthfamily-koham/app/models"
	validators "github.com/Subha-Research/svasthfamily-koham/app/validators"
)

type IACLService interface {
}

type ACLService struct {
	ar_model sf_models.SFAccessRelationshipModel
}

func (acl_s *ACLService) CreateSFRelationship(sf_user_id string, rb validators.ACLPostBody) error {
	// 1. Check if this sf_user_id exist or not in family_relationship collection.
	// 2. Check if given role and access is supported by us
	// 3.
	database := sf_models.Database{}
	ar_coll, _, err := database.GetCollectionAndSession("sf_access_relationship")
	if err != nil {
		log.Fatal("Errro in  getting collection and session. Stopping server", err)
	}
	// Dependency injection pattern
	acl_s.ar_model.Collection = ar_coll
	acl_s.ar_model.InsertAllSFAccessRelationship(sf_user_id, rb)
	log.Println("rb", rb)
	return nil
}
