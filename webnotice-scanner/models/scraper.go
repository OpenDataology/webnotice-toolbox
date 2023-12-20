/*
scraper.go

Copyright (c) 2022 The OpenDataology Authors
All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package models

import (
	"fmt"
	"net/http"

	"strings"

	"golang.org/x/net/html"
)

type Copyright struct {
	Id              int    `gorm:"type:int;primary_key;autoIncrement" json:"id"`
	License_name    string `gorm:"type:TEXT" json:"license_name,omitempty"`
	Licensor        string `gorm:"type:TEXT" json:"licensor,omitempty"`
	License_content string `gorm:"type:TEXT" json:"license_content,omitempty"`
	License_url     string `gorm:"type:TEXT" json:"license_url,omitempty"`
}

func fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	z := html.NewTokenizer(resp.Body)

	var content string
	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return content, nil
		case tt == html.TextToken:
			content += z.Token().Data
		}
	}
}

func GetCopyright(url string) (copyright *Copyright, err error) {
	content, err := fetch(url)
	if err != nil {
		return nil, err
	}

	// Possible copyright keywords
	keywords := []string{"Copyright", "©", "版权"}

	for _, keyword := range keywords {
		index := strings.LastIndex(content, keyword)
		if index != -1 {
			copyright = &Copyright{
				License_content: content[index:],
			}
			return copyright, nil
		}
	}

	err = fmt.Errorf("no copyright information found")
	return
}
