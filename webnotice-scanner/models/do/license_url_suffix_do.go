/*
user.go

Copyright (c) 2022 The OpenDataology Authors
All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package do

type LicenseUrlSuffix struct {
	Id        int    `gorm:"type:int;primary_key;autoIncrement" json:"id"`
	UrlSuffix string `gorm:"type:varchar" json:"url_suffix"`
}

// TableName 设置表名
func (LicenseUrlSuffix) TableName() string {
	return "t_license_url_suffix"
}
