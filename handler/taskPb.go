package handler

import (
	"context"
	"factory/exam/services"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	servicesPb "github.com/lk153/proto-tracking-gen/go/services"
)

//TaskPBHandler
type TaskPBHandler struct {
	servicesPb.UnimplementedTaskServiceServer
	taskService services.TaskServiceInterface
	tracer      trace.Tracer
}

//NewTaskPBHandler
func NewTaskPBHandler(
	taskService services.TaskServiceInterface,
) *TaskPBHandler {
	tracer := otel.Tracer("NewTaskPBHandler")
	return &TaskPBHandler{
		tracer:      tracer,
		taskService: taskService,
	}
}

//Get List of Task
func (s *TaskPBHandler) Get(
	ctx context.Context,
	req *servicesPb.TaskListRequest,
) (*servicesPb.TaskListResponse, error) {
	ctx, span := s.tracer.Start(ctx, "Get")
	defer span.End()

	limit := req.GetLimit()
	page := req.GetPage()
	ids := req.GetIds()
	tasks := s.taskService.GetList(ctx, int(limit), int(page), ids)
	data := s.taskService.Transform(tasks)

	return &servicesPb.TaskListResponse{
		Data: data,
	}, nil
}

//GetSingle Task
func (s *TaskPBHandler) GetSingle(
	ctx context.Context,
	req *servicesPb.SingleTaskRequest,
) (*servicesPb.SingleTaskResponse, error) {
	ctx, span := s.tracer.Start(ctx, "GetSingle")
	defer span.End()

	ID := req.GetId()
	task := s.taskService.GetSingle(ctx, int(ID))
	data := s.taskService.TransformSingle(task)
	return &servicesPb.SingleTaskResponse{
		Data: data,
	}, nil
}

//Create Task
func (s *TaskPBHandler) Create(
	ctx context.Context,
	req *servicesPb.TaskCreateRequest,
) (*servicesPb.TaskCreateResponse, error) {
	ctx, span := s.tracer.Start(ctx, "Create")
	defer span.End()

	data := req.GetData()
	task := s.taskService.Create(ctx, data)
	data = s.taskService.TransformSingle(task)

	return &servicesPb.TaskCreateResponse{
		Data: data,
	}, nil
}
