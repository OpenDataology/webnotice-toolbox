/*
user.go

Copyright (c) 2022 The OpenDataology Authors
All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package controllers

import (
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/utils"
	"net/http"
	"strconv"

	service "github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/services"

	"github.com/gin-gonic/gin"
)

func (a *BasicInfo) GetCopyright(c *gin.Context) {
	// url := c.GetString("url")
	// id, _ := strconv.Atoi(c.GetString("id"))

	if c.Query("url") == "" ||
		c.Query("id") == "" {

		a.JsonFail(c, http.StatusBadRequest, "not allow")
		return
	}

	url := utils.URLResolve(c.Query("url"), "")
	id, _ := strconv.Atoi(c.Query("id"))

	res := service.GetCopyrightService(c, url, id)
	a.JsonSuccess(c, http.StatusOK, res)
}
