package service

import (
	"github.com/aos-stack/aos-example/demo-ddd/domain"
)

type DemoService interface {
	List(demoID int64) domain.DemoModel
}
