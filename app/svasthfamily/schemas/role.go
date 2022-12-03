package sf_schemas

type RoleSchema struct {
	RoleID   string `bson:"role_id" json:"role_id"`
	RoleEnum int    `bson:"role_enum" json:"role_enum"`
	RoleKey  string `bson:"role_key" json:"role_key"`
	IsActive bool   `bson:"is_active" json:"is_active"`
	IsDelete bool   `bson:"is_delete" json:"is_delete"`

	// Do we need it?
	// Audit    AuditSchema `bson:"audit" json:"audit"`
}
