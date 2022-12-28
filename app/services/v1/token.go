package services

import (
	"fmt"
	"time"

	"github.com/Subha-Research/svasthfamily-koham/app/cache"
	"github.com/Subha-Research/svasthfamily-koham/app/constants"
	"github.com/Subha-Research/svasthfamily-koham/app/dto"
	"github.com/Subha-Research/svasthfamily-koham/app/errors"
	"github.com/Subha-Research/svasthfamily-koham/app/models"
	"github.com/Subha-Research/svasthfamily-koham/app/validators"
	"github.com/golang-jwt/jwt/v4"
)

type TokenService struct {
}

type TokenClaims struct {
	FUserID    string
	AccessList []dto.AccessRelation
	jwt.RegisteredClaims
}

func (ts TokenService) GetToken(f_user_id string) (*string, error) {
	// TODO:: call cache to check if token is present there
	r := cache.Redis{}
	redis_client := r.SetupTokenRedisDB()
	tc := cache.TokenCache{}
	// Dependeny injection
	tc.RedisClient = &redis_client
	tv, _ := tc.Get(f_user_id)

	if tv == nil {
		// TODO:: if not found in cache, check DB
		database := models.Database{}
		t_coll, _, err := database.GetCollectionAndSession("sf_tokens")
		tfromdb := models.TokenModel{}
		tfromdb.Collection = t_coll
		result, err := tfromdb.GetToken(f_user_id)
		if err != nil {
			error_data := map[string]string{
				"id": f_user_id,
			}
			return nil, errors.KohamError("KSE-4010", error_data)
		}

		tokenkey := result.TokenKey
		return &tokenkey, nil
	}
	return tv, nil
}

func (ts TokenService) CreateToken(f_user_id string) (*dto.CreateTokenResponse, error) {
	// TODO :: Before proceeding check if token already exist
	// for the f_user_id and if exist then do not create and
	// return the existing token
	arm := models.AccessRelationshipModel{}
	database := models.Database{}
	tm := models.TokenModel{}

	// TODO:: Move this to constants
	signing_key := []byte("OUR_SECRET_KEY")
	token_expiry := jwt.NewNumericDate(time.Now().Add(constants.TokenExpiryTTL * time.Hour))

	ar_coll, _, err := database.GetCollectionAndSession("sf_access_relationship")
	if err != nil {
		return nil, err
	}
	arm.Collection = ar_coll
	all_access_relations, err := arm.GetAllAccessRelationship(f_user_id)
	if err != nil {
		return nil, err
	}
	dto := dto.AccessRelationshipDTO{}
	acl_dto, err := dto.FormatAllAccessRelationship(all_access_relations)

	if err != nil {
		return nil, err
	}

	// Create the claims
	claims := TokenClaims{
		f_user_id,
		acl_dto,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: token_expiry,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    constants.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signing_key)
	// TODO :: Use log instead of fmt
	fmt.Printf("%v %v", ss, err)

	token_coll, _, err := database.GetCollectionAndSession("sf_tokens")
	if err != nil {
		return nil, err
	}
	tm.Collection = token_coll
	data, insert_err := tm.InsertToken(f_user_id, ss, token_expiry.Time)
	if insert_err != nil {
		return nil, insert_err
	}
	// Cache call
	// Dependency injection pattern
	r := cache.Redis{}
	redis_client := r.SetupTokenRedisDB()
	tc := cache.TokenCache{}
	tc.RedisClient = &redis_client
	tc.Set(f_user_id, data.TokenKey, constants.TokenExpiryTTL)
	return data, nil
}

func (ts TokenService) ParseToken(token_string string, f_user_id string) ([]dto.AccessRelation, error) {
	token, err := jwt.ParseWithClaims(token_string, &TokenClaims{}, func(*jwt.Token) (secret interface{}, err error) {
		// TODO:: Move this to constants
		return []byte("OUR_SECRET_KEY"), nil
	})
	if err != nil {
		return nil, errors.KohamError("KSE-4009")
	}
	claims, ok := token.Claims.(*TokenClaims)
	if ok && token.Valid {
		// TODO:: Use log instead of fmt
		// fmt.Printf("%v %v", claims.FUserID, claims.RegisteredClaims.Issuer)
		if claims.FUserID != f_user_id && claims.RegisteredClaims.Issuer == constants.Issuer {
			return nil, errors.KohamError("KSE-4009")
		}
	} else {
		return nil, errors.KohamError("KSE-4009")
	}
	return claims.AccessList, nil
}
func (ts TokenService) ValidateTokenAccess(token *string, f_user_id string, rb validators.TokenRequestBody) (*dto.ValidateTokenResponse, error) {
	db_token_key, err := ts.GetToken(f_user_id)
	if err != nil {
		return nil, err
	}
	if &db_token_key != &token {
		return nil, errors.KohamError("KSE-4009")
	}
	acesslist, _ := ts.ParseToken(*token, f_user_id)
	for _, v := range acesslist {
		if v.ChildMemberID == rb.ChildmemberID {
			for _, e := range v.AccessEnums.([]float64) {
				if e == rb.AccessEnum {
					// Build Response
					vtr := dto.ValidateTokenResponse{}
					vtr.Access = true
					return &vtr, nil
				}
			}
		}
	}
	return nil, errors.KohamError("KSE-4009")
}
