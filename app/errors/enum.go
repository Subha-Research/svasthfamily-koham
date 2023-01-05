package errors

type ErrorStruct struct {
	StatusCode   int
	ErrorCode    string
	ErrorMessage string
}

var ErrorEnums = map[string]*ErrorStruct{
	"KSE-5001": {500, "KSE-5001", "Interal Server Error"},
	// All 400 error codes
	"KSE-4001": {400, "KSE-4001", "Content-Type header is not invalid."},
	"KSE-4002": {400, "KSE-4002", "x-service-id header parsing error."},
	"KSE-4003": {403, "KSE-4003", "Authorization header missing."},
	"KSE-4004": {403, "KSE-4004", "Authorization header format is incorrect."},
	"KSE-4005": {403, "KSE-4005", "You are not authorized to access :resource_type: resource."},
	"KSE-4006": {400, "KSE-4006", "Invalid value for :key: ."},

	"KSE-4007": {400, "KSE-4007", "User ID didn't matched."},
	"KSE-4008": {400, "KSE-4008", "Child member id can't be same as parent member id"},
	"KSE-4009": {403, "KSE-4009", "Invalid token"},
	"KSE-4010": {404, "KSE-4010", "No token found for family user id :id:"},
	"KSE-4011": {400, "KSE-4011", "Access Relationship already exist for the mentioned Child-Parent pair"},
	"KSE-4012": {404, "KSE-4012", "No access relationship found for family user id :id:"},
	"KSE-4013": {404, "KSE-4013", "Relationship didn't established"},
	"KSE-4014": {404, "KSE-4014", "404 URL not found"},
}
