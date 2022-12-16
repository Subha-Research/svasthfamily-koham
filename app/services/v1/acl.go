package services

import (
	validators "github.com/Subha-Research/svasthfamily-koham/app/validators"
)

type IACLService interface {
}

type ACLService struct {
}

func (acl_s *ACLService) CreateSFRelationship(sf_user_id string, rb validators.ACLPostBody) error {
	// 1. Check if this sf_user_id exist or not in family_relationship collection.
	// 2. Check if given role and access is supported by us
	// 3.
	return nil
}
