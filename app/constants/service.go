package constants

const TOKEN_COLLECTION = "tokens"
const ACL_COLLECTION = "access_relationship"
const ROLE_COLLECTION = "roles"
const ACCESS_COLLECTION = "accesses"

var HEADER_VALIDATOR_STRATEGY = map[string]string{
	"service":       "x-service-id",
	"authorization": "authorization",
}
