package database

import (
	"context"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	database = os.Getenv("DB_NAME")
)

type Database interface {
	Collection(string) Collection
}

type Collection interface {
	InsertOne(context.Context, interface{}) (interface{}, error)
	InsertMany(context.Context, []interface{}) ([]interface{}, error)
	FindOne(context.Context, interface{}) SingleResult
}

type SingleResult interface {
	Decode(interface{}) error
}

type mongoDatabase struct {
	db *mongo.Database
}

type mongoCollection struct {
	coll *mongo.Collection
}

type mongoSingleResult struct {
	sr *mongo.SingleResult
}

func New() Database {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port)))

	if err != nil {
		log.Fatal(err)

	}

	return &mongoDatabase{db: client.Database(database)}
}

func (md *mongoDatabase) Collection(colName string) Collection {
	collection := md.db.Collection(colName)
	return &mongoCollection{coll: collection}
}

func (mc *mongoCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	id, err := mc.coll.InsertOne(ctx, document)
	return id.InsertedID, err
}

func (mc *mongoCollection) InsertMany(ctx context.Context, documents []interface{}) ([]interface{}, error) {
	ids, err := mc.coll.InsertMany(ctx, documents)
	return ids.InsertedIDs, err
}

func (mc *mongoCollection) FindOne(ctx context.Context, filter interface{}) SingleResult {
	sr := mc.coll.FindOne(ctx, filter)
	return &mongoSingleResult{sr: sr}
}

func (sr *mongoSingleResult) Decode(v interface{}) error {
	return sr.sr.Decode(v)
}
