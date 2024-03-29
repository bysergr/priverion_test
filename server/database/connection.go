package database

import (
	"context"
	"fmt"

	"github.com/bysergr/priverion_test/server/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


// newConnection is a function that connects to the mongodb database
func newConnection() *mongo.Database {

	env := utils.GetENV()

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/?authSource=admin", env.USER_DB, env.PASSWORD_DB, env.HOST_DB, env.PORT_DB)

	opts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	return client.Database(env.DATABASE)
}
