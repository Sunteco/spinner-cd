package mongo

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go-project-template/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
	"time"
)

var _ ResourceConsumedSnapshotRepository = (*ResourceConsumedSnapshotRepositoryImpl)(nil)

type (
	ResourceConsumedSnapshotRepository interface {
		baseRepository[model.ResourceConsumedSnapshot]
		HelloWorld()
	}

	ResourceConsumedSnapshotRepositoryImpl struct {
		*repository
	}
)

func (r ResourceConsumedSnapshotRepositoryImpl) Insert(ctx context.Context, data *model.ResourceConsumedSnapshot) error {
	//TODO implement me
	//TODO implement me
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	_, err := r.conn.InsertOne(ctx, data)
	return err
}

func (r ResourceConsumedSnapshotRepositoryImpl) FindOne(ctx context.Context, id string) (*model.ResourceConsumedSnapshot, error) {
	//TODO implement me
	//TODO implement me
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data model.ResourceConsumedSnapshot

	err = r.conn.FindOne(ctx, &data, bson.M{"_id": oid})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (r ResourceConsumedSnapshotRepositoryImpl) Update(ctx context.Context, data *model.ResourceConsumedSnapshot) error {
	//TODO implement me
	data.UpdateAt = time.Now()
	_, err := r.conn.ReplaceOne(ctx, bson.M{"_id": data.ID}, data)
	return err
}

func (r ResourceConsumedSnapshotRepositoryImpl) Delete(ctx context.Context, id string) error {
	//TODO implement me
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidObjectId
	}

	_, err = r.conn.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}

func (r ResourceConsumedSnapshotRepositoryImpl) HelloWorld() {
	fmt.Println("hello world")
}

func NewResourceConsumedSnapshotRepository() ResourceConsumedSnapshotRepository {
	envErr := godotenv.Load("etc/.env")
	if envErr != nil {
		log.Fatal(envErr)
	}
	fmt.Printf("url>>>>>>>>>>>>>>: %s", os.Getenv("url"))
	conn := mon.MustNewModel(os.Getenv("url"), os.Getenv("database"), "resource_consumed_snapshot")
	return &ResourceConsumedSnapshotRepositoryImpl{
		repository: newRepository(conn),
	}
}
