package models_mock

import (
	"time"

	"github.com/Subha-Research/svasthfamily-koham/app/dtos"
)

type TokenModelMock struct {
}

func (tm *TokenModelMock) InsertToken(f_user_id string, token string, expiry time.Time) (*dtos.CreateTokenResponse, error) {
	ctr := dtos.CreateTokenResponse{}
	ctr.FamilyUserID = f_user_id
	ctr.TokenExpiry = expiry
	ctr.TokenKey = token
	return &ctr, nil
}
