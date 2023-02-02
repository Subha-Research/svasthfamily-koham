package interfaces

import (
	"github.com/Subha-Research/svasthfamily-koham/app/dtos"
	"github.com/Subha-Research/svasthfamily-koham/app/validators"
)

type ITokenService interface {
	GetToken(string) (*string, error)
	CreateToken(string) (*dtos.CreateTokenResponse, error)
	ParseToken(string, string) error
	ValidateTokenAccess(*string, string, validators.ValidateTokenRB) (*dtos.ValidateTokenResponse, error)
	DeleteToken(*string, *string) error
}
