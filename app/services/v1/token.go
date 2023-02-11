package services

import (
	"log"
	"time"

	"github.com/Subha-Research/svasthfamily-koham/app/cache"
	"github.com/Subha-Research/svasthfamily-koham/app/common"
	"github.com/Subha-Research/svasthfamily-koham/app/constants"
	"github.com/Subha-Research/svasthfamily-koham/app/dtos"
	"github.com/Subha-Research/svasthfamily-koham/app/errors"
	"github.com/Subha-Research/svasthfamily-koham/app/models"
	"github.com/Subha-Research/svasthfamily-koham/app/validators"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type TokenService struct {
	Model   *models.TokenModel
	ARModel *models.AccessRelationshipModel
	Cache   *cache.TokenCache
}

type TokenClaims struct {
	jwt.RegisteredClaims
}

func (ts *TokenService) GetTokenDataFromDb(f_user_id string) (*dtos.GetTokenResponse, error) {
	result, err := ts.Model.GetToken(f_user_id)
	if err != nil {
		error_data := map[string]string{
			"id": f_user_id,
		}
		return nil, errors.KohamError("SFKSE-4010", error_data)
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
			return nil, errors.KohamError("SFKSE-4010", error_data)
		}
		tv = &result.TokenKey
	}
	return tv, nil
}

func (ts *TokenService) CreateToken(f_user_id string) (*dtos.CreateTokenResponse, error) {
	err := ts.DeleteToken(&f_user_id, nil)
	if err != nil {
		log.Println("Error in deleting token in create token API", err)
	}

	time_util := common.TimeUtil{}
	signing_key := []byte(constants.TokenSigingKey)
	token_expiry := jwt.NewNumericDate(time_util.CurrentTimeInUTC().Add(constants.TokenExpiryTTL * time.Hour))

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
		return nil, errors.KohamError("SFKSE-5001")
	}
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
		return errors.KohamError("SFKSE-4009")
	}
	claims, ok := token.Claims.(*TokenClaims)
	if ok && token.Valid {
		if claims.RegisteredClaims.Issuer != constants.Issuer {
			log.Println("Token issuer did not match")
			return errors.KohamError("SFKSE-4009")
		}
	} else {
		return errors.KohamError("SFKSE-4009")
	}
	return nil
}
func (ts *TokenService) ValidateTokenAccess(token *string, f_user_id string, rb validators.ValidateTokenRB) (*dtos.ValidateTokenResponse, error) {
	db_token_key, err := ts.GetToken(f_user_id)
	if err != nil {
		return nil, err
	}
	if *db_token_key != *token {
		log.Printf("Given token %s did not match with database", *token)
		return nil, errors.KohamError("SFKSE-4009")
	}
	var acl_doc bson.M

	switch rb.AccessEnum {
	case 115:
		acl_doc, err = ts.ARModel.GetAccessRelationship(nil, nil, &f_user_id, f_user_id, rb.ChildUserID)
	case 101:
		acl_doc, err = ts.ARModel.GetAccessRelationship(&rb.FamilyId, nil, &f_user_id, f_user_id, rb.ChildUserID)
	case 104:
		acl_doc, err = ts.ARModel.GetAccessRelationship(&rb.FamilyId, &rb.FamilyMemberId, nil, f_user_id, rb.ChildUserID)
	case 112:
		acl_doc, err = ts.ARModel.GetAccessRelationship(&rb.FamilyId, nil, nil, f_user_id, rb.ChildUserID)
	case 116:
		acl_doc, err = ts.ARModel.GetAccessRelationship(&rb.FamilyId, nil, nil, f_user_id, rb.ChildUserID)
	case 113:
		acl_doc, err = ts.ARModel.GetAccessRelationship(nil, nil, &f_user_id, f_user_id, rb.ChildUserID)
	case 114:
		acl_doc, err = ts.ARModel.GetAccessRelationship(nil, nil, &f_user_id, f_user_id, rb.ChildUserID)
	default:
		return nil, errors.KohamError("SFKSE-4015")
	}
	// If error
	if err != nil {
		return nil, err
	}
	log.Println("Access realtionship document", acl_doc)

	if acl_doc != nil {
		vtr := dtos.ValidateTokenResponse{}
		vtr.Access = true
		return &vtr, nil
	}
	return nil, errors.KohamError("SFKSE-4009")
}

// *token to this function could be nil, in case getting called from CreateToken
func (ts *TokenService) DeleteToken(f_user_id *string, token *string) error {
	result, err := ts.GetTokenDataFromDb(*f_user_id)
	if err != nil {
		error_data := map[string]string{
			"id": *f_user_id,
		}
		return errors.KohamError("SFKSE-4010", error_data)
	}

	var delete_token_key *string
	if token == nil {
		delete_token_key = &result.TokenKey
	} else if token != nil && result.TokenKey == *token {
		delete_token_key = token
	} else {
		log.Printf("Given token %s did not match with database", *token)
		return errors.KohamError("SFKSE-4009")
	}
	err_del := ts.Model.DeleteToken(f_user_id, delete_token_key)
	if err_del != nil {
		return err_del
	}
	// Delete from redis also
	ts.Cache.InvalidateKey(*f_user_id)
	return nil
}
