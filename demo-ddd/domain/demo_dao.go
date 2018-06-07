package domain

type DemoDAO interface {
	List(demoID int64) DemoModel
}
