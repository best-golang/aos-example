package service

import (
	"github.com/aos-stack/aos-example/demo-ddd/domain"
	"fmt"
)

type DemoServiceImplements struct { //接口的定义
	DemoDAO domain.DemoDAO //要使用的接口起名字
}

func NewDemoServiceImplements(demoDAO domain.DemoDAO) *DemoServiceImplements {
	return &DemoServiceImplements{DemoDAO: demoDAO}
}

func (mys *DemoServiceImplements) List(demoID int64) domain.DemoModel {
	fmt.Println("============")
	fmt.Println(demoID)
	fmt.Println("come")
	fmt.Println("============")
	return mys.DemoDAO.List(demoID)
}
