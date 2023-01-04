package interfaces

import (
	"github.com/Subha-Research/svasthfamily-koham/app/dto"
	"github.com/Subha-Research/svasthfamily-koham/app/validators"
)

type ITokenService interface {
	GetToken(string) (*string, error)
	CreateToken(string) (*dto.CreateTokenResponse, error)
	ParseToken(string, string) error
	ValidateTokenAccess(*string, string, validators.TokenRequestBody) (*dto.ValidateTokenResponse, error)
	DeleteToken(*string, *string) error
}
