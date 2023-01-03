package services_mock

import (
	"time"

	"github.com/Subha-Research/svasthfamily-koham/app/constants"
	"github.com/Subha-Research/svasthfamily-koham/app/dto"
	"github.com/Subha-Research/svasthfamily-koham/app/services/v1"
	"github.com/Subha-Research/svasthfamily-koham/app/validators"
	models_mock "github.com/Subha-Research/svasthfamily-koham/tests/mocks/models"
	"github.com/golang-jwt/jwt/v4"
)

type TokenServiceTest struct {
	Model   *models_mock.TokenModelMock
	ARModel *models_mock.AccessRelationshipModelMock
}

func (tst *TokenServiceTest) CreateToken(f_user_id string) (*dto.CreateTokenResponse, error) {
	all_access_relations, _ := tst.ARModel.GetAllAccessRelationship(f_user_id)
	signing_key := []byte(constants.TokenSigingKey)
	loc, err := time.LoadLocation("Asia/Kolkata")
	token_expiry := jwt.NewNumericDate(time.Date(2021, 1, 1, 0, 0, 0, 0, loc).Add(constants.TokenExpiryTTL * time.Hour))
	dto := dto.AccessRelationshipDTO{}
	acl_dto, err := dto.FormatAllAccessRelationship(all_access_relations)
	if err != nil {
		return nil, err
	}
	claims := services.TokenClaims{
		FUserID:    f_user_id,
		AccessList: acl_dto,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: token_expiry,
			IssuedAt:  jwt.NewNumericDate(time.Date(2021, 1, 1, 0, 0, 0, 0, loc)),
			Issuer:    constants.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signing_key)
	data, insert_err := tst.Model.InsertToken(f_user_id, ss, token_expiry.Time)
	if insert_err != nil {
		return nil, insert_err
	}
	return data, nil
}

func (tst *TokenServiceTest) GetToken(f_user_id string) (*string, error) {
	return nil, nil
}

func (tst *TokenServiceTest) ParseToken(token_string string, f_user_id string) ([]dto.AccessRelation, error) {
	return nil, nil
}

func (tst *TokenServiceTest) ValidateTokenAccess(token *string, f_user_id string, rb validators.TokenRequestBody) (*dto.ValidateTokenResponse, error) {
	return nil, nil
}

func (tst *TokenServiceTest) DeleteToken(f_user_id *string, tokecn *string) error {
	return nil
}
