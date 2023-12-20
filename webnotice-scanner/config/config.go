/*
config.go

Copyright (c) 2022 The OpenDataology Authors
All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"fmt"
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/config/database"
	"os"
	"strconv"

	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/routes"
	"gorm.io/gorm"
)

type Config struct {
	Addr        string
	DSN         string
	TOKEN       string
	MaxIdleConn int
}

var config Config

func Load() error {
	config.Addr = os.Getenv("ADDR")
	config.DSN = os.Getenv("DSN")
	config.TOKEN = os.Getenv("TOKEN")
	max_idle, err := strconv.Atoi(os.Getenv("MAX_IDLE_CONN"))

	//config.Addr = "127.0.0.1:8080"
	//config.DSN = "root:passw@tcp(127.0.0.1:3000)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	//config.TOKEN = "1"
	//max_idle, err := strconv.Atoi("100")

	if err != nil {
		fmt.Println("err max_idle_conn")
		return err
	}
	config.MaxIdleConn = max_idle
	return nil
}

func Get() Config {
	return config
}

func InitDB() (*gorm.DB, error) {
	conf := Get()
	return database.InitDB(conf.DSN)
}

func InitAndStart() {

	if err := Load(); err != nil {
		fmt.Println("Failed to load configuration")
		return
	}

	fmt.Println("config is loaded")
	_, err := InitDB()
	if err != nil {
		fmt.Println("err open databases")
		return
	}
	fmt.Println("DB is connected")

	router := routes.InitRouter()
	router.Run(Get().Addr)
}
