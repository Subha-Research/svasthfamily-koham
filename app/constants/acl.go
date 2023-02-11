package constants

import "golang.org/x/exp/maps"

type UPDATE_INFO struct {
	Type       string
	AccessEnum float64
}

var UPDATE_TYPE = map[string]*UPDATE_INFO{
	"UPDATE_FAMILY_ID": {
		Type:       "UPDATE_FAMILY_ID",
		AccessEnum: 113,
	},
	"UPDATE_FAMILY_MEMBER_ID": {
		Type:       "UPDATE_FAMILY_MEMBER_ID",
		AccessEnum: 114,
	},
	"UPDATE_SFM_ACCESS": {
		Type:       "UPDATE_SFM_ACCESS",
		AccessEnum: 103,
	},
}

var CHILD_DEFAULT_ACCESS = map[float64]string{
	102: "UPDATE_SFM_DETAILS",
	104: "VIEW_SFM_DETAILS",
	107: "ADD_SFM_HEALTH_RECORD",
	108: "VIEW_SFM_HEALTH_RECORD",
	109: "ADD_SFM_HEALTH_RECORD_METADATA",
	110: "VIEW_SFM_HEALTH_RECORD_METADATA",
	111: "SUGGEST_SFM_HEALTH_RECORD_METADATA",
	112: "VIEW_ALL_SFM_BASIC_DETAILS", // View all members basic details
	116: "GET_FAMILY_DETAILS",
}

type ACLConstants struct {
}

func (acl_c *ACLConstants) GetConstantAccessList(acl_type string) map[float64]string {
	var access_list map[float64]string
	switch acl_type {
	case "CHILD":
		access_list = CHILD_DEFAULT_ACCESS
	case "HEAD":
		head_access := map[float64]string{
			101: "ADD_SFM",
			103: "UPDATE_SFM_ACCESS",
			105: "VIEW_SFM_ACCESS",
			106: "DELETE_SFM",
			113: "UPDATE_FAMILY_ID",
			114: "UPDATE_FAMILY_MEMBER_ID",
			115: "CREATE_FAMILY",
		}
		maps.Copy(head_access, CHILD_DEFAULT_ACCESS)
		access_list = head_access
	default:
		access_list = nil
	}
	return access_list
}
