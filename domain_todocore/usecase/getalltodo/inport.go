package getalltodo

import (
	"todo_app/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	Page int
	Size int
}

type InportResponse struct {
	Count int64
	Items []any
}
