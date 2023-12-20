/*
router.go

Copyright (c) 2022 The OpenDataology Authors
All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package routes

import (
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(Cors())
	v1 := router.Group("/api/v1")
	{
		basic := new(controllers.BasicInfo)
		v1.GET("/copyright/compliance/task", basic.GetCopyright)

	}
	return router

}
