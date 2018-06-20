package container

import (
	"github.com/aos-stack/aos-example/app/controller"
	"github.com/aos-stack/aos-example/demo-ddd/domain"
	"github.com/aos-stack/aos-example/demo-ddd/infrastructure/persistence/xorm"
	"github.com/aos-stack/aos-example/demo-ddd/service"

	aosContainer "github.com/aos-stack/aos/container"
	"github.com/go-xorm/xorm"
)

var containerInstance *Container

func GetContainer() *Container {
	if containerInstance == nil {
		containerInstance = &Container{}
		containerInstance.init()
	}
	return containerInstance
}

type Container struct {
	DemoDDDApi *controller.DemoDDDApi
}

func (c *Container) init() {
	var dbDemoConn *xorm.Engine = initEng(0)
	var demoDAO domain.DemoDAO = persistence.NewDemoDAOXorm(dbDemoConn)
	var demoService = service.NewDemoServiceImplements(demoDAO)

	c.DemoDDDApi = controller.NewDemoDDDApi(demoService)
	aosContainer.ContainerSet("demoDDDApi", c.DemoDDDApi)
}

func initEng(num int) *xorm.Engine {
	return nil
}
