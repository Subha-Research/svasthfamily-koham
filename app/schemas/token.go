package schemas

import "time"

type Token struct {
	TokenID      string      `bson:"token_id" json:"token_id"`
	TokenKey     string      `bson:"token_key" json:"token_key"`
	ExpiresAt    time.Time   `bson:"expires_at" json:"expires_at"`
	FamilyUserID string      `bson:"family_user_id" json:"family_user_id"`
	Audit        AuditSchema `bson:"audit" json:"audit"`
	// RoleEnum     string      `bson:"role_id" json:"role_enum"`
}
