package persistence

import (
	"aos-example/demo-ddd/domain"

	"github.com/go-xorm/xorm"
)

type DemoDAOXorm struct {
	Table domain.DemoModel
	domain.DemoDAO
	Engine *xorm.Engine
}

func (c *DemoDAOXorm) List(b int64) domain.DemoModel {

	// var bp domain.ProjectModel
	// c.Engine.Table("gd_project").Where("id = 8").Get(&bp)

	return c.Table
}

func NewDemoDAOXorm(eng *xorm.Engine) *DemoDAOXorm {
	var a *DemoDAOXorm = &DemoDAOXorm{}
	a.Engine = eng
	return a
}
