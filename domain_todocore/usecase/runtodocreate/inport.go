package runtodocreate

import (
	"todo_app/domain_todocore/model/entity"
	"todo_app/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.TodoCreateRequest
}

type InportResponse struct {
	Todo *entity.Todo
}
