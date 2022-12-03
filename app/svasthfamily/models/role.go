package sf_models

import (
	"context"
	"fmt"
	"log"

	sf_enums "github.com/Subha-Research/koham/app/svasthfamily/enums"
	sf_schemas "github.com/Subha-Research/koham/app/svasthfamily/schemas"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RoleModel struct {
	Collection *mongo.Collection
	Session    *mongo.Session
}

func (rm *RoleModel) InsertAllRoles(coll *mongo.Collection) {
	// var collection = rco

	var roleDocs []interface{}
	roleEnums := sf_enums.RoleEnums
	for i := 0; i < 2; i++ {
		var roleInterface interface{}
		role := &sf_schemas.RoleSchema{
			RoleID:   uuid.NewString(),
			RoleEnum: i,
			RoleKey:  roleEnums[i],
			IsActive: true,
			IsDelete: false,
		}
		// Convert role struct to interface
		roleInterface = role
		roleDocs = append(roleDocs, roleInterface)
	}

	// Call insert many of mongo
	opts := options.InsertMany().SetOrdered(false)
	res, err := coll.InsertMany(context.TODO(), roleDocs, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inserted documents with IDs %v\n", res.InsertedIDs)
}
