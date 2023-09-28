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

var _ SyncResourceRepository = (*syncResourceRepositoryImpl)(nil)

type (
	// SyncResourceRepository TestModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTestModel.
	SyncResourceRepository interface {
		baseRepository[model.SyncResourceInfo]
		FindAll(ctx context.Context) ([]model.SyncResourceInfo, error)
	}

	syncResourceRepositoryImpl struct {
		*repository
	}
)

func (m syncResourceRepositoryImpl) Insert(ctx context.Context, data *model.SyncResourceInfo) error {
	//TODO implement me
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreatedDate = time.Now().UnixNano() / 1e6
	}

	_, err := m.conn.InsertOne(ctx, data)
	return err
}

func (m syncResourceRepositoryImpl) FindOne(ctx context.Context, id string) (*model.SyncResourceInfo, error) {
	//TODO implement me
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data model.SyncResourceInfo

	err = m.conn.FindOne(ctx, &data, bson.M{"_id": oid})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m syncResourceRepositoryImpl) FindAll(ctx context.Context) ([]model.SyncResourceInfo, error) {
	//TODO implement me
	var data []model.SyncResourceInfo
	err := m.conn.Find(ctx, &data, bson.M{})
	switch err {
	case nil:
		return data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m syncResourceRepositoryImpl) Update(ctx context.Context, data *model.SyncResourceInfo) error {
	//TODO implement me
	data.CreatedDate = time.Now().UnixNano() / 1e6

	_, err := m.conn.ReplaceOne(ctx, bson.M{"_id": data.ID}, data)
	return err
}

func (m syncResourceRepositoryImpl) Delete(ctx context.Context, id string) error {
	//TODO implement me
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidObjectId
	}

	_, err = m.conn.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}

func NewSyncResourceRepository() SyncResourceRepository {
	envErr := godotenv.Load("etc/.env")
	if envErr != nil {
		log.Fatal(envErr)
	}
	fmt.Printf("url>>>>>>>>>>>>>>: %s", os.Getenv("url"))
	conn := mon.MustNewModel(os.Getenv("url"), os.Getenv("database"), "sync_resource_info")
	return &syncResourceRepositoryImpl{
		repository: newRepository(conn),
	}
}
