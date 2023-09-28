package service

import (
	"context"
	"fmt"
	"go-project-template/internal/database/mongo"
	"go-project-template/internal/model"
)

func Test(ctx context.Context, svcCtx *ServiceContext) ([]model.SyncResourceInfo, error) {
	conn := mongo.NewSyncResourceRepository()
	data, err := conn.FindAll(ctx)
	if err != nil {
		panic(err)
		return nil, err
	}
	mongo.NewResourceConsumedSnapshotRepository().HelloWorld()
	fmt.Printf("datas: %v\n", data)
	return data, nil
}
