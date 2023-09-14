/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package models

import (
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
)

// TODO Please modify the following code to fit your needs
// This object conforms to what the frontend currently sends.
type GoogleConnection struct {
	helper.RestConnection `mapstructure:",squash"`
	helper.BaseConnection `mapstructure:",squash"`
	AccessToken           `mapstructure:",squash"`
	Name                  string `gorm:"type:varchar(100);uniqueIndex" json:"name" validate:"required"`
	Proxy                 string `json:"proxy"`
	RateLimitPerHour      int    `comment:"api request rate limit per hour"`
	SpreadsheetID         string `json:"spreadsheetID"`
	FirstValue            string `json:"firstValue"`
	LastValue             string `json:"lastValue"`
}

type TestConnectionRequest struct {
	Endpoint    string `json:"endpoint"`
	AccessToken `mapstructure:",squash"`
}

type RestConnection struct {
	helper.RestConnection `mapstructure:",squash"`
	helper.AccessToken    `mapstructure:",squash"`
	Endpoint              string `mapstructure:"endpoint" validate:"required" json:"endpoint"`
}

type AccessToken struct {
	Token string `mapstructure:"token" validate:"required" json:"token" encrypt:"yes"`
}

func (GoogleConnection) TableName() string {
	return "_tool_google_connections"
}
