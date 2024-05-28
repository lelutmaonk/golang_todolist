package payload

import (
	"todo_app/shared/gogen"
)

type Payload struct {
	Data      any                   `json:"data"`
	Publisher gogen.ApplicationData `json:"publisher"`
	TraceID   string                `json:"traceId"`
}
