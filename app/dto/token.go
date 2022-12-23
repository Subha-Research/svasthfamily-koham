package dto

import (
	"time"

	"github.com/Subha-Research/svasthfamily-koham/app/schemas"
)

type TokenDTO struct {
}

type CreateTokenResponse struct {
	TokenKey     string              `json:"token_key"`
	TokenExpiry  time.Time           `json:"expires_at"`
	FamilyUserID string              `json:"family_user_id"`
	Audit        schemas.AuditSchema `json:"audit"`
}
