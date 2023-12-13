/*
connect.go

Copyright (c) 2022 The OpenDataology Authors
All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package database

import (
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	conf := config.Get()
	db, err := gorm.Open(mysql.Open(conf.DSN), &gorm.Config{})

	if err == nil {
		DB = db
		return db, err
	}
	return nil, err
}
