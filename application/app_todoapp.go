package application

import (
	"todo_app/domain_todocore/controller/restapi"
	"todo_app/domain_todocore/gateway/withgorm"
	"todo_app/domain_todocore/usecase/getalltodo"
	"todo_app/domain_todocore/usecase/runtodocheck"
	"todo_app/domain_todocore/usecase/runtodocreate"
	"todo_app/shared/config"
	"todo_app/shared/gogen"
	"todo_app/shared/infrastructure/logger"
	"todo_app/shared/infrastructure/token"
)

type todoapp struct{}

func NewTodoapp() gogen.Runner {
	return &todoapp{}
}

func (todoapp) Run() error {

	const appName = "todoapp"

	cfg := config.ReadConfig()

	appData := gogen.NewApplicationData(appName)

	log := logger.NewSimpleJSONLogger(appData)

	jwtToken := token.NewJWTToken(cfg.JWTSecretKey)

	datasource := withgorm.NewGateway(log, appData, cfg)

	primaryDriver := restapi.NewController(appData, log, cfg, jwtToken)

	primaryDriver.AddUsecase(
		//
		getalltodo.NewUsecase(datasource),
		runtodocheck.NewUsecase(datasource),
		runtodocreate.NewUsecase(datasource),
	)

	primaryDriver.RegisterRouter()

	primaryDriver.Start()

	return nil
}
