package sf_services

import (
	sf_validators "github.com/Subha-Research/koham/app/svasthfamily/validators"
)

type IACLService interface {
}

type ACLService struct {
}

func (acl_s *ACLService) CreateSFRelationship(sf_user_id string, rb sf_validators.ACLPostBody) error {
	// 1. Check if this sf_user_id exist or not in family_relationship collection.
	// 2. Check if given role and access is supported by us
	// 3.
	return nil
}
