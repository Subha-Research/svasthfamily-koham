package sf_models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type RoleModel struct {
	Collection *mongo.Collection
	Session    *mongo.Session
}

// func (rm *RoleModel) InsertAllRoles() {
// 	var collection = rm.Collection

// 	var roleDocs []interface{}
// 	// roles := [2]string{"sf_head", "sf_member"}
// 	roleEnums := sf_enums.RoleEnums
// 	for i := 0; i < 2; i++ {
// 		var roleInterface interface{}
// 		role := &sf_schemas.RoleSchema{
// 			RoleID:   uuid.NewString(),
// 			RoleEnum: roleEnums[i.(string)],
// 			RoleKey:  "sf_head",
// 			IsActive: true,
// 			IsDelete: false,
// 		}
// 		// Convert role
// 		roleInterface = role
// 		roleDocs = append(roleDocs, roleInterface)
// 	}

// 	// Call insertmany of mongo
// 	opts := options.InsertMany().SetOrdered(false)
// 	res, err = collection.InsertMany(context.TODO(), roleDocs, opts)
// }
