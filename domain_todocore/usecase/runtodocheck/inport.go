package runtodocheck

import (
	"todo_app/domain_todocore/model/entity"
	"todo_app/domain_todocore/model/vo"
	"todo_app/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	TodoID vo.TodoID
}

type InportResponse struct {
	Todo *entity.Todo
}
