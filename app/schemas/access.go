package schemas

type Access struct {
	AccessID   string `bson:"access_id" json:"access_id"`
	AccessEnum int    `bson:"access_enum" json:"access_enum"`
	AccessKey  string `bson:"access_key" json:"access_key"`
	IsActive   bool   `bson:"is_active" json:"is_active"`
	IsDelete   bool   `bson:"is_delete" json:"is_delete"`

	// Do we need audits for this collection ?
	// Audit      AuditSchema `bson:"audit" json:"audit"`
}
