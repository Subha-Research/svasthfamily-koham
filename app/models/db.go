package sf_models

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/url"

	configs "github.com/Subha-Research/koham/app/configs"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
}

func (d *Database) getURI() string {
	config := configs.LoadConfig()
	port := config["db.port"].(string)
	username := config["db.username"]
	password := url.QueryEscape(config["db.password"].(string))
	host := config["db.host"].(string)
	connPool := config["db.maxPoolSize"]
	hostPort := net.JoinHostPort(host, port)

	uri := fmt.Sprintf("mongodb://%s:%s@%s/?&maxPoolSize=%v&w=majority", username, password, hostPort, connPool)
	return uri
}

func (d *Database) getClient() *mongo.Client {
	uri := d.getURI()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Panic(err)
	}

	// Close the DB connection if required
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	log.Println("Successfully connected and pinged.")
	return client
}

func (d *Database) GetCollectionAndSession(collection string) (*mongo.Collection, *mongo.Session, error) {
	opts := options.Session().SetDefaultReadConcern(readconcern.Majority())
	config := configs.LoadConfig()
	database := config["db.database"].(string)

	client := d.getClient()
	session, err := client.StartSession(opts)
	if err != nil {
		log.Println("Error in starting a session", err)
		return nil, nil, fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
	}
	return client.Database(database).Collection(collection), &session, nil
}
