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
)

func GetCopyrightService(c *gin.Context, url string, id int) (h gin.H) {

	print("url:" + url + " id :" + string(id))
	copyrightCompalianceHandlerQueset := dto.CopyrightComplianceHandlerRequestDTO{
		AibomId:   id,
		SourceUrl: utils.URLResolve(url, ""),
	}

	//遍历
	copyrightCompalianceHandlerList := handlers.CopyrightComplianceHandlerList
	for i := 0; i < len(copyrightCompalianceHandlerList); i++ {
		go copyrightCompalianceHandlerList[i].Handle(copyrightCompalianceHandlerQueset)
	}
	res := gin.H{
		"data": "task run",
	}
	return res

	// user, err := models.GetCopyright(url)
	// if err != nil {
	// 	if err.Error() == "Url is invalid!" {
	// 		c.JSON(http.StatusOK, gin.H{"err": "Url is invalid!"})
	// 	} else {
	// 		c.JSON(http.StatusOK, gin.H{"err": "DB Error"})
	// 	}
	// 	return
	// }
	// res := gin.H{
	// 	"data": &user,
	// }
	// return res
}

func URLResolve(url string) {
	panic("unimplemented")
}
