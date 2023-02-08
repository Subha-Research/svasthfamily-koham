package errors

type ErrorStruct struct {
	StatusCode   int
	ErrorCode    string
	ErrorMessage string
}

var ErrorEnums = map[string]*ErrorStruct{
	"SFKSE-5001": {500, "SFKSE-5001", "Interal Server Error"},
	// All 400 error codes
	"SFKSE-4001": {400, "SFKSE-4001", "Content-Type header is not invalid."},
	"SFKSE-4002": {400, "SFKSE-4002", "x-service-id header parsing error."},
	"SFKSE-4003": {403, "SFKSE-4003", "Authorization header missing."},
	"SFKSE-4004": {403, "SFKSE-4004", "Authorization header format is incorrect."},
	"SFKSE-4005": {403, "SFKSE-4005", "You are not authorized to access :resource_type: resource."},
	"SFKSE-4006": {400, "SFKSE-4006", "Invalid value for :key: ."},

	"SFKSE-4007": {400, "SFKSE-4007", "No token found the family user id :f_user_id:."},
	"SFKSE-4008": {400, "SFKSE-4008", "Child user id can't be same as parent user id"},
	"SFKSE-4009": {403, "SFKSE-4009", "Invalid token"},
	"SFKSE-4010": {404, "SFKSE-4010", "No token found for family user id :id:"},
	"SFKSE-4011": {400, "SFKSE-4011", "Access Relationship already exist for the mentioned Child-Parent pair"},
	"SFKSE-4012": {404, "SFKSE-4012", "No access relationship found for family user id :id:"},
	"SFKSE-4013": {404, "SFKSE-4013", "Relationship didn't established"},
	"SFKSE-4014": {404, "SFKSE-4014", "404 URL not found"},
	"SFKSE-4015": {405, "SFKSE-4015", "Invalid operation"},
	"SFKSE-4016": {400, "SFKSE-4016", "All fields can't be empty"},
	"SFKSE-4017": {400, "SFKSE-4017", "No access relatioship found."},
	"SFKSE-4018": {400, "SFKSE-4018", "Can't update! :field: is already updated."},
}
