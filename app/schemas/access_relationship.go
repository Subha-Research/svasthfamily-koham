package schemas

type AccessRelationshipSchema struct {
	AccessRelationshipID string      `bson:"access_relationship_id" json:"access_relationship_id"`
	ParentFamilyUserID   string      `bson:"parent_family_user_id" json:"parent_family_user_id"`
	ChildFamilyUserID    string      `bson:"child_family_user_id" json:"child_family_user_id"`
	AccessEnum           []float64   `bson:"access_enums" json:"access_enums"`
	IsDelete             bool        `bson:"is_delete" json:"is_delete"`
	Audit                AuditSchema `bson:"audit" json:"audit"`
}
