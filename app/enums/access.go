package enums

var Head_Default_Accesses = map[float64]string{
	// SF: svasthfamily
	// SFM: svasthfamily member
	101: "ADD_SFM",
	102: "UPDATE_SFM",
	103: "VIEW_SFM",
	104: "DELETE_SFM",
	105: "ADD_SFM_HEALTH_RECORD",
	106: "VIEW_SFM_HEALTH_RECORD",
	107: "ADD_SFM_HEALTH_RECORD_METADATA",
	108: "VIEW_SFM_HEALTH_RECORD_METADATA",
	109: "SUGGEST_SFM_HEALTH_RECORD_METADATA",
}
var Child_Default_Accesses = map[float64]string{
	// SF: svasthfamily
	// SFM: svasthfamily member
	102: "UPDATE_SFM",
	103: "VIEW_SFM",
	105: "ADD_SFM_HEALTH_RECORD",
	106: "VIEW_SFM_HEALTH_RECORD",
	107: "ADD_SFM_HEALTH_RECORD_METADATA",
	108: "VIEW_SFM_HEALTH_RECORD_METADATA",
	109: "SUGGEST_SFM_HEALTH_RECORD_METADATA",
}
