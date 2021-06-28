package services

import (
	"context"

	entities_pb "github.com/lk153/proto-tracking-gen/go/entities"

	"factory/exam/repo"
)

//TaskServiceInterface
type TaskServiceInterface interface {
	GetList(ctx context.Context, limit int, page int, ids []uint64) []*repo.TaskModel
	GetSingle(ctx context.Context, id int) *repo.TaskModel
	Create(ctx context.Context, data *entities_pb.TaskInfo) *repo.TaskModel
	Transform(input []*repo.TaskModel) []*entities_pb.TaskInfo
	TransformSingle(prod *repo.TaskModel) *entities_pb.TaskInfo
}
