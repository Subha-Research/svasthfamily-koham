package services

import (
	"log"
	"time"

	"github.com/Subha-Research/svasthfamily-koham/app/cache"
	"github.com/Subha-Research/svasthfamily-koham/app/common"
	"github.com/Subha-Research/svasthfamily-koham/app/constants"
	"github.com/Subha-Research/svasthfamily-koham/app/dto"
	"github.com/Subha-Research/svasthfamily-koham/app/errors"
	"github.com/Subha-Research/svasthfamily-koham/app/models"
	"github.com/Subha-Research/svasthfamily-koham/app/validators"
	"github.com/golang-jwt/jwt/v4"
)

type TokenService struct {
	Model   *models.TokenModel
	ARModel *models.AccessRelationshipModel
	Cache   *cache.TokenCache
}

type TokenClaims struct {
	jwt.RegisteredClaims
}

func (ts *TokenService) GetTokenDataFromDb(f_user_id string) (*dto.GetTokenResponse, error) {
	database := models.Database{}
	t_coll, _, err := database.GetCollectionAndSession(constants.TokenCollection)
	if err != nil {
		return nil, errors.KohamError("KSE-5001")
	}

	ts.Model.Collection = t_coll
	result, err := ts.Model.GetToken(f_user_id)
	if err != nil {
		error_data := map[string]string{
			"id": f_user_id,
		}
		return nil, errors.KohamError("KSE-4010", error_data)
	}

	return result, nil
}

func (ts *TokenService) GetToken(f_user_id string) (*string, error) {
	tv, _ := ts.Cache.Get(f_user_id)

	if tv == nil {
		result, err := ts.GetTokenDataFromDb(f_user_id)
		if err != nil {
			error_data := map[string]string{
				"id": f_user_id,
			}
			return nil, errors.KohamError("KSE-4010", error_data)
		}
		tv = &result.TokenKey
	}
	return tv, nil
}

func (ts *TokenService) CreateToken(f_user_id string) (*dto.CreateTokenResponse, error) {
	err := ts.DeleteToken(&f_user_id, nil)
	if err != nil {
		log.Println("Error in deleting token in create token API", err)
	}

	database := models.Database{}
	time_util := common.TimeUtil{}
	signing_key := []byte(constants.TokenSigingKey)
	token_expiry := jwt.NewNumericDate(time_util.CurrentTimeInUTC().Add(constants.TokenExpiryTTL * time.Hour))

	ar_coll, _, err := database.GetCollectionAndSession(constants.ACL_COLLECTION)
	if err != nil {
		return nil, err
	}
	ts.ARModel.Collection = ar_coll

	if err != nil {
		log.Println("Error in formatting access relationship data", err)
		return nil, err
	}

	// Create the claims
	claims := TokenClaims{
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: token_expiry,
			IssuedAt:  jwt.NewNumericDate(*time_util.CurrentTimeInUTC()),
			Issuer:    constants.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signing_key)
	if err != nil {
		log.Println("Error while signing token", err)
		return nil, errors.KohamError("KSE-5001")
	}

	token_coll, _, err := database.GetCollectionAndSession(constants.TokenCollection)
	if err != nil {
		return nil, err
	}
	ts.Model.Collection = token_coll
	data, insert_err := ts.Model.InsertToken(f_user_id, ss, token_expiry.Time)
	if insert_err != nil {
		return nil, insert_err
	}
	// Cache call
	ts.Cache.Set(f_user_id, data.TokenKey, constants.TokenExpiryTTL)
	return data, nil
}

func (ts *TokenService) ParseToken(token_string string, f_user_id string) error {
	token, err := jwt.ParseWithClaims(token_string, &TokenClaims{}, func(*jwt.Token) (secret interface{}, err error) {
		return []byte(constants.TokenSigingKey), nil
	})
	if err != nil {
		return errors.KohamError("KSE-4009")
	}
	claims, ok := token.Claims.(*TokenClaims)
	if ok && token.Valid {
		if claims.RegisteredClaims.Issuer != constants.Issuer {
			log.Println("Token issuer did not match")
			return errors.KohamError("KSE-4009")
		}
	} else {
		return errors.KohamError("KSE-4009")
	}
	return nil
}
func (ts *TokenService) ValidateTokenAccess(token *string, f_user_id string, rb validators.TokenRequestBody) (*dto.ValidateTokenResponse, error) {
	db_token_key, err := ts.GetToken(f_user_id)
	if err != nil {
		return nil, err
	}
	if *db_token_key != *token {
		log.Printf("Given token %s did not match with database", *token)
		return nil, errors.KohamError("KSE-4009")
	}
	all_access_relations, err := ts.ARModel.GetAllAccessRelationship(f_user_id)
	acl_dto := dto.AccessRelationshipDTO{}
	if err != nil {
		return nil, err
	}
	access_list, err := acl_dto.FormatAllAccessRelationship(all_access_relations)

	for _, v := range access_list {
		if v.ChildMemberID == rb.ChildmemberID {
			for _, e := range v.AccessEnums.([]interface{}) {
				if e.(float64) == rb.AccessEnum {
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

// *token to this function could be nil, in case getting called from CreateToken
func (ts *TokenService) DeleteToken(f_user_id *string, token *string) error {
	result, err := ts.GetTokenDataFromDb(*f_user_id)
	if err != nil {
		error_data := map[string]string{
			"id": *f_user_id,
		}
		return errors.KohamError("KSE-4010", error_data)
	}

	database := models.Database{}
	t_coll, _, err := database.GetCollectionAndSession(constants.TokenCollection)
	if err != nil {
		return errors.KohamError("KSE-5001")
	}

	ts.Model.Collection = t_coll
	var delete_token_key *string
	if token == nil {
		delete_token_key = &result.TokenKey
	} else if token != nil && *&result.TokenKey == *token {
		delete_token_key = token
	} else {
		log.Printf("Given token %s did not match with database", *token)
		return errors.KohamError("KSE-4009")
	}
	err_del := ts.Model.DeleteToken(f_user_id, delete_token_key)
	if err_del != nil {
		return err_del
	}
	// Delete from redis also
	ts.Cache.InvalidateKey(*f_user_id)
	return nil
}
