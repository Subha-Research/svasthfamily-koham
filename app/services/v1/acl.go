package services

import (
	"log"

	"github.com/Subha-Research/svasthfamily-koham/app/enums"
	models "github.com/Subha-Research/svasthfamily-koham/app/models"
	validators "github.com/Subha-Research/svasthfamily-koham/app/validators"
)

type IACLService interface {
	CreateSFRelationship(string, validators.ACLPostBody) error
}

type ACLService struct {
	ARModel models.AccessRelationshipModel
}

func (acl_s *ACLService) CreateAccessRelationship(f_user_id string, rb validators.ACLPostBody) error {
	database := models.Database{}
	ar_coll, _, err := database.GetCollectionAndSession("access_relationship")
	if err != nil {
		log.Fatal("Errro in  getting collection and session. Stopping server", err)
	}
	// Dependency injection pattern
	acl_s.ARModel.Collection = ar_coll
	if enums.Roles[rb.RoleEnum] != "FAMILY_HEAD" {
		_, err_get_doc_head := acl_s.ARModel.GetAccessRelationship(f_user_id, f_user_id)
		if err_get_doc_head != nil {
			return err_get_doc_head
		}
		_, err_get_doc_parent := acl_s.ARModel.GetAccessRelationship(rb.ParentMemberID, rb.ParentMemberID)
		if err_get_doc_parent != nil {
			return err_get_doc_parent
		}
		_, err_get_doc_head_parent := acl_s.ARModel.GetAccessRelationship(f_user_id, rb.ParentMemberID)
		if err_get_doc_head_parent != nil {
			return err_get_doc_head_parent
		}
	}
	inserted_doc, err := acl_s.ARModel.InsertAllAccessRelationship(f_user_id, rb)
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
	acl_s.ARModel.Collection = ar_coll
	_, err_update_doc := acl_s.ARModel.UpdateAccessRelationship(f_head_user_id, rb)
	if err_update_doc != nil {
		return err_update_doc
	}

	return nil
}
