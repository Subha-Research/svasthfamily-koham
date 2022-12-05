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
	"KSE-4005": {403, "KSE-4005", "You are not authorized to access {resource_type} resource."},
}
