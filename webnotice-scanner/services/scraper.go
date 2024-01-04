/*
user.go

Copyright (c) 2022 The OpenDataology Authors
All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package service

import (
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/handlers"
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/models/dto"
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func GetCopyrightService(c *gin.Context, url string, id int) (h gin.H) {

	log.Printf("url:%s,bomid:%d", url, id)
	copyrightCompalianceHandlerQueset := dto.CopyrightComplianceHandlerRequestDTO{
		AibomId:   id,
		SourceUrl: utils.URLResolve(url, ""),
	}

	//遍历
	copyrightComplianceHandlerList := handlers.CopyrightComplianceHandlerList
	for i := 0; i < len(copyrightComplianceHandlerList); i++ {
		go copyrightComplianceHandlerList[i].Handle(copyrightCompalianceHandlerQueset)
	}
	res := gin.H{
		"data": "task run",
	}
	return res
}

func URLResolve(url string) {
	panic("unimplemented")
}
