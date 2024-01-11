package port

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	InsertOne(ctx context.Context, user interface{}, opts ...*options.InsertOneOptions) error
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) error
	FindOne(ctx context.Context, target interface{}, filter interface{}, opts ...*options.FindOneOptions) error
	Find(ctx context.Context, target interface{}, filter interface{}, opts ...*options.FindOptions) error
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) error
}

type ProjectRepository interface {
	InsertOne(ctx context.Context, project interface{}, opts ...*options.InsertOneOptions) error
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) error
	FindOne(ctx context.Context, target interface{}, filter interface{}, opts ...*options.FindOneOptions) error
	Find(ctx context.Context, target interface{}, filter interface{}, opts ...*options.FindOptions) error
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) error
}

type HistoryRepository interface {
	InsertOne(ctx context.Context, history interface{}, opts ...*options.InsertOneOptions) error
	Find(ctx context.Context, target interface{}, filter interface{}, opts ...*options.FindOptions) error
	FindOne(ctx context.Context, target interface{}, filter interface{}, opts ...*options.FindOneOptions) error
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) error
}

type MultimediaRepository interface {
	UploadFile(filePath string) (string, error)
}
