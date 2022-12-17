package services

import (
	"fmt"
	"log"
	"time"

	"github.com/Subha-Research/svasthfamily-koham/app/constants"
	"github.com/Subha-Research/svasthfamily-koham/app/models"
	"github.com/golang-jwt/jwt/v4"
)

type TokenService struct {
}

func (ts TokenService) CreateToken(f_user_id string) (string, error) {
	mySigningKey := []byte("OUR_SECRET_KEY")
	token_expiry := jwt.NewNumericDate(time.Now().Add(24 * time.Hour))

	type MyCustomClaims struct {
		FUserID string
		jwt.RegisteredClaims
	}

	// Create the claims
	claims := MyCustomClaims{
		f_user_id,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: token_expiry,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    constants.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)

	tm := models.TokenModel{}
	database := models.Database{}
	token_coll, _, err := database.GetCollectionAndSession("sf_tokens")
	if err != nil {
		log.Println("Errro in  getting collection and session. Stopping server", err)
		return "", err
	}
	tm.Collection = token_coll
	tm.InsertToken(f_user_id, ss, token_expiry.Time)

	return "", nil
}
