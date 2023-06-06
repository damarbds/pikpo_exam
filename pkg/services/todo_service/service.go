package todo_service

import (
	"context"
	"pikpo_exam/pkg/db"
	"pikpo_exam/pkg/entities"

	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ITodoService interface {
	Create(ctx context.Context, req *CreateRequest) (*ReadByIdResponse, error)
	Update(ctx context.Context, req *UpdateRequest) (*ReadByIdResponse, error)
	ReadById(ctx context.Context, req *ReadByIdRequest) (*ReadByIdResponse, error)
	ReadAll(ctx context.Context, req *ReadByIdRequest) (*ReadAllResponse, error)
}

type todoService struct {
	dbHandler db.IDatabaseAdapter
}

func NewTodoService(dbHandler db.IDatabaseAdapter) ToDoServiceServer {
	return &todoService{
		dbHandler: dbHandler,
	}
}

func (handler *todoService) Create(ctx context.Context, req *CreateRequest) (*ReadByIdResponse, error) {
	log.Debug().Msgf("[TODO Service] entering create...")

	model := &entities.Todo{
		Name: req.GetName(),
	}
	_, err := handler.dbHandler.Insert(model)
	if err != nil {
		log.Warn().Err(err).Msgf("[TODO Service] error insert data: %s", err.Error())
		return nil, err
	}
	return &ReadByIdResponse{
		Id:   uint32(model.ID),
		Name: model.Name,
	}, nil
}

func (handler *todoService) Update(ctx context.Context, req *UpdateRequest) (*ReadByIdResponse, error) {
	log.Debug().Msgf("[TODO Service] entering update...")

	var model = &entities.Todo{}
	result, err := handler.dbHandler.ReadById(model, uint(req.GetId()))
	if err != nil {
		log.Warn().Err(err).Msgf("[TODO Service] error read data: %s", err.Error())
		return nil, err
	}

	model = result.(*entities.Todo)
	model.Name = req.GetName()

	_, err = handler.dbHandler.Update(model, model.ID)
	if err != nil {
		log.Warn().Err(err).Msgf("[TODO Service] error update data: %s", err.Error())
		return nil, err
	}
	return &ReadByIdResponse{
		Id:   uint32(model.ID),
		Name: model.Name,
	}, nil
}

func (handler *todoService) ReadById(ctx context.Context, req *ReadByIdRequest) (*ReadByIdResponse, error) {
	log.Debug().Msgf("[TODO Service] entering read by id...")

	var model = &entities.Todo{}
	result, err := handler.dbHandler.ReadById(model, uint(req.GetId()))
	if err != nil {
		log.Warn().Err(err).Msgf("[TODO Service] error read data: %s", err.Error())
		return nil, err
	}

	model = result.(*entities.Todo)

	return &ReadByIdResponse{
		Id:   uint32(model.ID),
		Name: model.Name,
	}, nil
}

func (handler *todoService) ReadAll(ctx context.Context, req *emptypb.Empty) (*ReadAllResponse, error) {
	log.Debug().Msgf("[TODO Service] entering read all...")

	var models []entities.Todo
	result, err := handler.dbHandler.ReadAll(models)
	if err != nil {
		log.Warn().Err(err).Msgf("[TODO Service] error read data: %s", err.Error())
		return nil, err
	}

	models = result.([]entities.Todo)

	items := make([]*ReadByIdResponse, 0)

	for _, model := range models {
		items = append(items, &ReadByIdResponse{
			Id:   uint32(model.ID),
			Name: model.Name,
		})
	}

	return &ReadAllResponse{
		Items: items,
	}, nil
}

func (handler *todoService) mustEmbedUnimplementedToDoServiceServer() {

}
