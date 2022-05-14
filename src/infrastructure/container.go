package infrastructure

import (
	"awesomeProject1/src/infrastructure/database"
	"awesomeProject1/src/infrastructure/repository"
	"awesomeProject1/src/use_case/check"
	"awesomeProject1/src/use_case/user/handler"
	"awesomeProject1/src/use_case/user/service"
	"gorm.io/gorm"
)

type ContainerDI struct {
	Config         Config
	DB             *gorm.DB
	CheckHanler    check.CheckHandler
	RepositoryUser repository.UserRepository
	ServiceUser    service.UserService
	HandlerUser    handler.UserHandler
}

func NewContainerDI(config Config) *ContainerDI {

	container := &ContainerDI{
		Config: config,
	}

	configDB := database.Config{
		Hostname: container.Config.DbHost,
		Port:     container.Config.DbPort,
		UserName: container.Config.DbUser,
		Password: container.Config.DbPassword,
		Database: container.Config.DbDatabase,
	}
	container.DB = database.InitGorm(&configDB)
	container.buildRepositories()
	container.build()
	return container
}

func (c *ContainerDI) buildRepositories() {
	c.RepositoryUser = repository.NewUserRepository(c.DB)
}

func (c *ContainerDI) build() {
	c.CheckHanler = check.NewCheckHandler()

	c.ServiceUser = service.NewUserService(c.RepositoryUser)
	c.HandlerUser = handler.NewUserHandler(c.ServiceUser)
}

func (c *ContainerDI) ShutDown() {}
