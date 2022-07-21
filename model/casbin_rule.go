package model

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

type casbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:512;uniqueIndex:unique_index"`
	V0    string `gorm:"size:512;uniqueIndex:unique_index"`
	V1    string `gorm:"size:512;uniqueIndex:unique_index"`
	V2    string `gorm:"size:512;uniqueIndex:unique_index"`
	V3    string `gorm:"size:512;uniqueIndex:unique_index"`
	V4    string `gorm:"size:512;uniqueIndex:unique_index"`
	V5    string `gorm:"size:512;uniqueIndex:unique_index"`
}

func (c *casbinRule) TableName() string {
	return "casbin_rule"
}

func NewEnforce() *casbin.Enforcer {
	a, _ := gormadapter.NewAdapterByDBWithCustomTable(db, &casbinRule{})
	e, _ := casbin.NewEnforcer("config/casbin_model.conf", a)
	return e
}

func (c *casbinRule) Check() {

}

func (c *casbinRule) done() {

}
