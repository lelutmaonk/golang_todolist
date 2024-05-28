package runtodocheck

import "todo_app/domain_todocore/model/repository"

type Outport interface {
	repository.FindOneTodoByIDRepo
	repository.SaveTodoRepo
}
