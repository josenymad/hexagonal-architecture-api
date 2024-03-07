package mongo

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client     *mongo.Client
	dbName     string
	collection string
}

func GetConnectionString() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	return uri
}

func ConnectToMongoDB(db string) (*mongo.Client, error) {
	uri := GetConnectionString()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Println("Error connecting to MongoDB client", err)
		return nil, err
	}

	if err := client.Database(db).RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		log.Println("Error pinging MongoDB", err)
		return nil, err
	}

	return client, nil
}

func NewMongoDB(db, collection string) (*MongoDB, error) {
	mongo, err := ConnectToMongoDB(db)
	if err != nil {
		log.Println("Error connecting to MongoDB")
		return nil, err
	}

	return &MongoDB{
		client:     mongo,
		dbName:     db,
		collection: collection,
	}, nil
}

func (m *MongoDB) PostData(data interface{}) error {
	mongo := m.client.Database(m.dbName).Collection(m.collection)
	_, err := mongo.InsertOne(context.TODO(), data)
	if err != nil {
		log.Println("Error writing to MongoDB", err)
		return err
	}
	// defer func() {
	// 	if err = m.client.Disconnect(context.TODO()); err != nil {
	// 		log.Println("Error disconnecting from MongoDB", err)
	// 	}
	// }()
	return nil
}

func (m *MongoDB) GetAllData() (interface{}, error) {
	mongo := m.client.Database(m.dbName).Collection(m.collection)
	cursor, err := mongo.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println("Error getting cursor containing all data from MongoDB", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var records []interface{}
	for cursor.Next(context.TODO()) {
		var response bson.M
		if err := cursor.Decode(&response); err != nil {
			log.Println("Error decoding cursor when getting all data")
			return nil, err
		}
		records = append(records, response)
	}

	return records, nil
}
