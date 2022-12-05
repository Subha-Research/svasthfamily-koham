package errors

type ErrorStruct struct {
	ErrorCode    string
	ErrorMessage string
}

var ErrorEnums = map[string]*ErrorStruct{
	"KSE-5001": {"KSE-5001", "Interal Server Error"},
	// All 400 error codes
	"KSE-4001": {"KSE-4001", "Content-Type header is not invalid."},
	"KSE-4002": {"KSE-4002", "x-service-id header parsing error."},
	"KSE-4003": {"KSE-4003", "Authorization header missing."},
	"KSE-4004": {"KSE-4004", "Authorization header format is incorrect."},
}
