/*
user.go

Copyright (c) 2022 The OpenDataology Authors
All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package do

type SourceWebPageLicenseKeyword struct {
	Id      int    `gorm:"type:int;primary_key;autoIncrement" json:"id"`
	Keyword string `gorm:"type:varchar" json:"keyword"`
}

// TableName 设置表名
func (SourceWebPageLicenseKeyword) TableName() string {
	return "t_source_web_page_license_keyword"
}
