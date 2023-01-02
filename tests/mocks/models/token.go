package models_mock

import (
	"time"

	"github.com/Subha-Research/svasthfamily-koham/app/dto"
)

type TokenModelMock struct {
}

func (tm *TokenModelMock) InsertToken(f_user_id string, token string, expiry time.Time) (*dto.CreateTokenResponse, error) {
	ctr := dto.CreateTokenResponse{}
	ctr.FamilyUserID = f_user_id
	ctr.TokenExpiry = expiry
	ctr.TokenKey = token
	return &ctr, nil
}
