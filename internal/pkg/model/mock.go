// Copyright 2022 Innkeeper Belm(孔令飞) <nosbelm@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://miniblog.

package model

import "time"

const TableNameStubInterface = "stub_interface"
const TableNameStubRule = "stub_rule"

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

// StubRule mapped from table <stub_rule>
type StubRule struct {
	ID          int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	InterfaceID int32     `gorm:"column:interface_id;not null" json:"interface_id"`
	MatchType   int32     `gorm:"column:match_type;not null;comment:1:match request query url, 2:match request body" json:"match_type"` // 1:match request query url, 2:match request body
	MatchRule   string    `gorm:"column:match_rule" json:"match_rule"`
	RespCode    string    `gorm:"column:resp_code" json:"resp_code"`
	RespHeader  string    `gorm:"column:resp_header" json:"resp_header"`
	RespBody    string    `gorm:"column:resp_body" json:"resp_body"`
	DelayTime   int32     `gorm:"column:delay_time;comment:ms" json:"delay_time"` // ms
	Desc        string    `gorm:"column:desc" json:"desc"`
	Meta        string    `gorm:"column:meta" json:"meta"`
	Status      string    `gorm:"column:status;not null" json:"status"`
	CreateTime  time.Time `gorm:"column:create_time;not null;default:current_timestamp()" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time;not null;default:current_timestamp()" json:"update_time"`
}

// TableName StubRule's table name
func (*StubRule) TableName() string {
	return TableNameStubRule
}
