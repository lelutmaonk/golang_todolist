package withgorm

import (
	"context"
	"todo_app/domain_todocore/model/entity"
	"todo_app/domain_todocore/model/vo"
	"todo_app/shared/config"
	"todo_app/shared/gogen"
	"todo_app/shared/infrastructure/logger"
)

type gateway struct {
	appData gogen.ApplicationData
	config  *config.Config
	log     logger.Logger
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *gateway {

	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,
	}
}

func (r *gateway) FindAllTodo(ctx context.Context, page, size int) ([]*entity.Todo, int64, error) {
	r.log.Info(ctx, "called")

	return nil, 0, nil
}

func (r *gateway) FindOneTodoByID(ctx context.Context, todoID vo.TodoID) (*entity.Todo, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) SaveTodo(ctx context.Context, obj *entity.Todo) error {
	r.log.Info(ctx, "called")

	return nil
}
