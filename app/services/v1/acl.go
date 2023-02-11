package services

import (
	"log"

	"github.com/Subha-Research/svasthfamily-koham/app/dtos"
	"github.com/Subha-Research/svasthfamily-koham/app/errors"
	models "github.com/Subha-Research/svasthfamily-koham/app/models"
	validators "github.com/Subha-Research/svasthfamily-koham/app/validators"
	"go.mongodb.org/mongo-driver/bson"
)

type ACLService struct {
	Model *models.AccessRelationshipModel
}

func (acl_s *ACLService) CreateAccessRelationship(f_user_id string, token *string, rb validators.ACLPostBody) (*[]dtos.CreateACLDTO, error) {
	var is_head_head_relation = true

	if token != nil {
		// Check if document already exists
		_, err_get_doc_head := acl_s.Model.GetAccessRelationship(&rb.FamilyID, nil, nil, f_user_id, f_user_id)
		if err_get_doc_head != nil {
			return nil, err_get_doc_head
		}
		_, err_get_doc_parent := acl_s.Model.GetAccessRelationship(&rb.FamilyID, nil, nil, rb.ParentUserID, rb.ParentUserID)
		if err_get_doc_parent != nil {
			return nil, err_get_doc_parent
		}
		_, err_get_doc_head_parent := acl_s.Model.GetAccessRelationship(&rb.FamilyID, nil, nil, f_user_id, rb.ParentUserID)
		if err_get_doc_head_parent != nil {
			return nil, err_get_doc_head_parent
		}
		is_head_head_relation = false
	}
	// Indicates getting called from another
	// microservice with x-service-id
	// Hence, request for creating HEAD_HEAD access relationship
	inserted_doc_response, err := acl_s.Model.InsertAllAccessRelationship(f_user_id, is_head_head_relation, rb)
	if err != nil {
		return nil, err
	}
	// log.Println("Inserted document", inserted_doc)
	return inserted_doc_response, nil
}

func (acl_s *ACLService) UpdateAccessRelationship(f_head_user_id string, update_type string, rb validators.ACLPutBody) (*dtos.UpdateACLDTO, error) {
	var doc bson.M
	var err error
	// Get Access relation
	doc, err = acl_s.Model.GetAccessRelationship(nil, nil, &f_head_user_id, rb.ParentUserID, rb.Access.ChildUserId)

	if err != nil {
		return nil, errors.KohamError("SFKSE-4015")
	}

	relation_type := doc["relationship_type"].(string)
	is_parent_head := doc["is_parent_head"].(bool)

	is_update_allowed := acl_s.isUpdateAllowed(relation_type, is_parent_head, rb.UpdateType)
	if !is_update_allowed {
		return nil, errors.KohamError("SFKSE-4015")
	}

	var update_doc_response *dtos.UpdateACLDTO
	var err_update_doc error

	update_doc_response, err_update_doc = acl_s.Model.UpdateAccessRelationship(f_head_user_id, rb)

	if err_update_doc != nil {
		return nil, err_update_doc
	}

	return update_doc_response, nil
}
func (acl_s *ACLService) UpdateFamilyID(f_head_user_id string, update_type string, rb validators.ACLPutBody) (*dtos.UpdateACLDTO, error) {
	var doc bson.M
	var err error

	// Get Access relation
	doc, err = acl_s.Model.GetAccessRelationship(nil, nil, &f_head_user_id, f_head_user_id, f_head_user_id)
	if err != nil {
		return nil, errors.KohamError("KSE-4015")
	}
	relation_type := doc["relationship_type"].(string)
	is_parent_head := doc["is_parent_head"].(bool)
	family_id := doc["family_id"].(string)

	if family_id != "" {
		error_data := map[string]string{
			"field": "family_id",
		}
		return nil, errors.KohamError("SFKSE-4018", error_data)
	}
	is_update_allowed := acl_s.isUpdateAllowed(relation_type, is_parent_head, rb.UpdateType)
	if !is_update_allowed {
		return nil, errors.KohamError("SFKSE-4015")
	}
	var update_doc_response *dtos.UpdateACLDTO
	var err_update_doc error

	update_doc_response, err_update_doc = acl_s.Model.UpdateFamilyID(f_head_user_id, rb)

	if err_update_doc != nil {
		return nil, err_update_doc
	}

	return update_doc_response, nil
}
func (acl_s *ACLService) UpdateFamilyMemberID(f_head_user_id string, update_type string, rb validators.ACLPutBody) (*dtos.UpdateACLDTO, error) {
	var doc bson.M
	var err error

	// Get Access relation
	doc, err = acl_s.Model.GetAccessRelationship(nil, nil, &f_head_user_id, f_head_user_id, f_head_user_id)
	if err != nil {
		return nil, errors.KohamError("KSE-4015")
	}
	relation_type := doc["relationship_type"].(string)
	is_parent_head := doc["is_parent_head"].(bool)
	family_member_id := doc["family_member_id"]

	if family_member_id != "" {
		error_data := map[string]string{
			"field": "family_member_id",
		}
		return nil, errors.KohamError("SFKSE-4018", error_data)
	}
	is_update_allowed := acl_s.isUpdateAllowed(relation_type, is_parent_head, rb.UpdateType)
	if !is_update_allowed {
		return nil, errors.KohamError("SFKSE-4015")
	}
	var update_doc_response *dtos.UpdateACLDTO
	var err_update_doc error

	update_doc_response, err_update_doc = acl_s.Model.UpdateFamilyMemberID(f_head_user_id, rb)

	if err_update_doc != nil {
		return nil, err_update_doc
	}

	return update_doc_response, nil
}
func (acl_s *ACLService) isUpdateAllowed(relation_type string, is_parent_head bool, update_type string) bool {
	switch true {
	case relation_type == "HEAD_HEAD" && (update_type == "UPDATE_FAMILY_ID" || update_type == "UPDATE_FAMILY_MEMBER_ID"):
		return true
	case relation_type == "PARENT_CHILD" && is_parent_head:
		return false
	case relation_type == "PARENT_CHILD":
		return false
	case relation_type == "CHILD_CHILD":
		// TODO :: Raise alert
		log.Println("ALERT:: Invalid case is getting executed")
		return false
	case relation_type == "HEAD_CHILD": // TODO:: test this, this should not be present
		return false
	}
	return true
}
