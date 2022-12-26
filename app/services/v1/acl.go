package services

import (
	"log"

	"github.com/Subha-Research/svasthfamily-koham/app/enums"
	models "github.com/Subha-Research/svasthfamily-koham/app/models"
	validators "github.com/Subha-Research/svasthfamily-koham/app/validators"
)

type IACLService interface {
}

type ACLService struct {
	ar_model models.AccessRelationshipModel
}

func (acl_s *ACLService) CreateAccessRelationship(f_user_id string, rb validators.ACLPostBody) error {
	// 1. Check if this sf_user_id exist or not in family_relationship collection.
	// 2. Check if given role and access is supported by us
	// 3.
	database := models.Database{}
	ar_coll, _, err := database.GetCollectionAndSession("access_relationship")
	if err != nil {
		log.Fatal("Errro in  getting collection and session. Stopping server", err)
	}
	// Dependency injection pattern
	acl_s.ar_model.Collection = ar_coll
	if enums.Roles[rb.RoleEnum] != "FAMILY_HEAD" {
		_, err_get_doc_head := acl_s.ar_model.GetAccessRelationship(f_user_id, f_user_id)
		if err_get_doc_head != nil {
			return err_get_doc_head
		}
		_, err_get_doc_parent := acl_s.ar_model.GetAccessRelationship(rb.ParentMemberID, rb.ParentMemberID)
		if err_get_doc_parent != nil {
			return err_get_doc_parent
		}
		_, err_get_doc_head_parent := acl_s.ar_model.GetAccessRelationship(f_user_id, rb.ParentMemberID)
		if err_get_doc_head_parent != nil {
			return err_get_doc_head_parent
		}
	}
	inserted_doc, err := acl_s.ar_model.InsertAllAccessRelationship(f_user_id, rb)
	if err != nil {
		return err
	}
	log.Println("Inserted document", inserted_doc)
	return nil
}

func (acl_s *ACLService) UpdateAccessRelationship(f_head_user_id string, rb validators.ACLPutBody) error {
	database := models.Database{}
	ar_coll, _, err := database.GetCollectionAndSession("access_relationship")
	if err != nil {
		log.Fatal("Errro in  getting collection and session. Stopping server", err)
	}
	// Dependency injection pattern
	acl_s.ar_model.Collection = ar_coll
	_, err_get_doc_child_parent := acl_s.ar_model.UpdateAccessRelationship(f_head_user_id, rb)
	if err_get_doc_child_parent != nil {
		return err_get_doc_child_parent
	}

	return nil
}
