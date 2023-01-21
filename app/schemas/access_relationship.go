package schemas

type AccessRelationshipSchema struct {
	AccessRelationshipID string      `bson:"access_relationship_id" json:"access_relationship_id"`
	HeadFamilyUserID     string      `bson:"head_family_user_id" json:"head_family_user_id"`
	ParentFamilyUserID   string      `bson:"parent_family_user_id" json:"parent_family_user_id"`
	FamilyID             string      `bson:"family_id" json:"family_id"`
	FamilyMemberID       string      `bson:"family_member_id" json:"family_member_id"`
	IsParentHead         bool        `bson:"is_parent_head" json:"is_parent_head"`
	RelationshipType     string      `bson:"relationship_type" json:"relationship_type"`
	ChildFamilyUserID    string      `bson:"child_family_user_id" json:"child_family_user_id"`
	AccessEnum           []float64   `bson:"access_enums" json:"access_enums"`
	IsDelete             bool        `bson:"is_delete" json:"is_delete"`
	Audit                AuditSchema `bson:"audit" json:"audit"`
}
