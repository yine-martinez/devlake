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

package archived

import (
	"github.com/apache/incubator-devlake/core/models/migrationscripts/archived"
)

// This object conforms to what the frontend currently sends.
type GoogleConnection struct {
	RestConnection `mapstructure:",squash"`
	AccessToken    `mapstructure:",squash"`
}

type RestConnection struct {
	BaseConnection `mapstructure:",squash"`
	Endpoint       string `mapstructure:"endpoint" validate:"required" json:"endpoint"`
}

type BaseConnection struct {
	Name             string `gorm:"type:varchar(100);uniqueIndex" json:"name" validate:"required"`
	Proxy            string `json:"proxy"`
	RateLimitPerHour int    `comment:"api request rate limit per hour"`
	SpreadsheetID    string `gorm:"type:varchar(100)"`
	FirstValue       string `gorm:"type:varchar(10)"`
	LastValue        string `gorm:"type:varchar(10)"`

	archived.Model
}

type AccessToken struct {
	Token string `mapstructure:"token" validate:"required" json:"token" encrypt:"yes"`
}

func (GoogleConnection) TableName() string {
	return "_tool_google_connections"
}
