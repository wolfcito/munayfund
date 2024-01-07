package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoImpl struct over mongo db for future dependencies injection
type MongoImpl struct {
	collection *mongo.Collection
}

func NewMongoDBCollection(ctx context.Context, dbName, ordersCollectionName, indexKey string) (*MongoImpl, error) {
	db, err := newDBCon(ctx, dbName)
	if err != nil {
		return nil, err
	}

	collection := db.Collection(ordersCollectionName)
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			primitive.E{Key: indexKey, Value: 1},
		},
	}

	_, err = collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return nil, err
	}

	return &MongoImpl{collection: collection}, nil
}

func (mc *MongoImpl) Disconnect(ctx context.Context) error {
	return mc.collection.Database().Client().Disconnect(ctx)
}

func newDBCon(ctx context.Context, dbName string) (*mongo.Database, error) {
	client, err := newClient(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database(dbName), nil
}

func newClient(ctx context.Context) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI(os.Getenv("MONGODB_CONNECTION_URI")).
		SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (mc *MongoImpl) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) error {
	_, err := mc.collection.InsertOne(ctx, document, opts...)

	return err
}

func (mc *MongoImpl) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) error {
	_, err := mc.collection.DeleteOne(ctx, filter, opts...)

	return err
}

func (mc *MongoImpl) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) error {
	_, err := mc.collection.UpdateOne(ctx, filter, update, opts...)

	return err
}

func (mc *MongoImpl) FindOne(ctx context.Context, target interface{}, filter interface{}, opts ...*options.FindOneOptions) error {
	sr := mc.collection.FindOne(ctx, filter, opts...)

	return sr.Decode(target)
}

func (mc *MongoImpl) Find(ctx context.Context, target interface{}, filter interface{}, opts ...*options.FindOptions) error {
	cursor, err := mc.collection.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}

	if err = cursor.All(ctx, target); err != nil {
		return err
	}

	return nil
}
