package model

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"os"
)

type casbinRule struct {
	ID    uint `gorm:"primaryKey;autoIncrement"`
	Ptype string
	V0    string
	V1    string
	V2    string
	V3    string
	V4    string
}

func (c *casbinRule) TableName() string {
	return "casbin_rule"
}

func NewEnforce() *casbin.Enforcer {
	path, _ := os.Getwd()
	a, _ := gormadapter.NewAdapterByDBWithCustomTable(db, &casbinRule{})
	e, _ := casbin.NewEnforcer(path+"/config/casbin_model.conf", a)
	return e
}

func (c *casbinRule) Check() {

}

func (c *casbinRule) done() {

}
