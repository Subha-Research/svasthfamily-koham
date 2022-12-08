package sf_services

import (
	sf_validators "github.com/Subha-Research/koham/app/svasthfamily/validators"
)

type IACLService interface {
}

type ACLService struct {
}

func (acl_s *ACLService) CreateSFRelationship(uID string, rb sf_validators.ACLPostBody) error {
	return nil
}
