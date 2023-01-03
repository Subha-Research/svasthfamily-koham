package dto

import (
	"time"

	"github.com/Subha-Research/svasthfamily-koham/app/schemas"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenDTO struct {
}

type CreateTokenResponse struct {
	TokenKey     string              `json:"token_key"`
	TokenExpiry  time.Time           `json:"expires_at"`
	FamilyUserID string              `json:"family_user_id"`
	Audit        schemas.AuditSchema `json:"audit"`
}

type GetTokenResponse struct {
	TokenKey     string              `json:"token_key"`
	TokenExpiry  primitive.DateTime  `json:"expires_at"`
	FamilyUserID string              `json:"family_user_id"`
	Audit        schemas.AuditSchema `json:"audit"`
}

type ValidateTokenResponse struct {
	Access bool `json:"is_access"`
}
