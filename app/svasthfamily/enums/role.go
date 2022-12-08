package sf_enums

type RoleInfo struct {
	Key  string
	Enum int
}

var RoleEnums = []string{"SF_HEAD", "SF_MEMBER"}
var Roles = map[int]string{
	0: "SF_HEAD",
	1: "SF_MEMBER",
}
