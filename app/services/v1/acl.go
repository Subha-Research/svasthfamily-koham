package services

import (
	"log"

	"github.com/Subha-Research/svasthfamily-koham/app/errors"
	models "github.com/Subha-Research/svasthfamily-koham/app/models"
	validators "github.com/Subha-Research/svasthfamily-koham/app/validators"
)

type ACLService struct {
	Model *models.AccessRelationshipModel
}

func (acl_s *ACLService) CreateAccessRelationship(f_user_id string, token *string, rb validators.ACLPostBody) error {
	var is_head_head_relation = true
	if token != nil {
		_, err_get_doc_head := acl_s.Model.GetAccessRelationship(nil, f_user_id, f_user_id)
		if err_get_doc_head != nil {
			return err_get_doc_head
		}
		_, err_get_doc_parent := acl_s.Model.GetAccessRelationship(nil, rb.ParentUserID, rb.ParentUserID)
		if err_get_doc_parent != nil {
			return err_get_doc_parent
		}
		_, err_get_doc_head_parent := acl_s.Model.GetAccessRelationship(nil, f_user_id, rb.ParentUserID)
		if err_get_doc_head_parent != nil {
			return err_get_doc_head_parent
		}
		is_head_head_relation = false
	}
	// Indicates getting called from another
	// microservice with x-service-id
	// Hence, request for creating HEAD_HEAD access relationship
	inserted_doc, err := acl_s.Model.InsertAllAccessRelationship(f_user_id, is_head_head_relation, rb)
	if err != nil {
		return err
	}
	log.Println("Inserted document", inserted_doc)
	return nil
}

func (acl_s *ACLService) UpdateAccessRelationship(f_head_user_id string, rb validators.ACLPutBody) error {
	// Get Access relation
	doc, err := acl_s.Model.GetAccessRelationship(&f_head_user_id, rb.ParentUserID, rb.Access.ChildUserId)
	if err != nil {
		return errors.KohamError("KSE-4015")
	}

	relation_type := doc["relationship_type"]
	is_parent_head := doc["is_parent_head"].(bool)
	switch true {
	case relation_type == "HEAD_HEAD":
		return errors.KohamError("KSE-4015")
	case relation_type == "PARENT_CHILD" && is_parent_head:
		return errors.KohamError("KSE-4015")
	case relation_type == "CHILD_CHILD":
		// TODO :: Raise alert
		log.Println("ALERT:: Invalid case is getting executed")
		return errors.KohamError("KSE-4015")
	case relation_type == "HEAD_CHILD":
		return errors.KohamError("KSE-4015")
	}

	_, err_update_doc := acl_s.Model.UpdateAccessRelationship(f_head_user_id, rb)
	if err_update_doc != nil {
		return err_update_doc
	}
	return nil
}
