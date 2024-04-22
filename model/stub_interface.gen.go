// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameStubInterface = "stub_interface"

// StubInterface mapped from table <stub_interface>
type StubInterface struct {
	ID            int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	URL           string    `gorm:"column:url;not null" json:"url"`
	DefRespCode   string    `gorm:"column:def_resp_code" json:"def_resp_code"`
	DefRespHeader string    `gorm:"column:def_resp_header" json:"def_resp_header"`
	DefRespBody   string    `gorm:"column:def_resp_body" json:"def_resp_body"`
	Owner         string    `gorm:"column:owner" json:"owner"`
	Desc          string    `gorm:"column:desc" json:"desc"`
	Meta          string    `gorm:"column:meta" json:"meta"`
	Status        string    `gorm:"column:status;not null" json:"status"`
	CreateTime    time.Time `gorm:"column:create_time;not null;default:current_timestamp()" json:"create_time"`
	UpdateTime    time.Time `gorm:"column:update_time;not null;default:current_timestamp()" json:"update_time"`
}

// TableName StubInterface's table name
func (*StubInterface) TableName() string {
	return TableNameStubInterface
}