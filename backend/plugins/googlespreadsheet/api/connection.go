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

package api

import (
	"context"
	"fmt"
	"github.com/apache/incubator-devlake/core/errors"
	"net/http"
	"time"

	"github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/googlespreadsheet/models"
)

//TODO Please modify the following code to fit your needs
func TestConnection(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	// decode
	var err errors.Error
	var connection models.TestConnectionRequest
	if err = api.Decode(input.Body, &connection, vld); err != nil {
		return nil, err
	}
	// test connection
	apiClient, err := api.NewApiClient(
		context.TODO(),
		connection.Endpoint,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", connection.Token),
		},
		3*time.Second,
		connection.Proxy,
		basicRes,
	)
	if err != nil {
		return nil, err
	}

	res, err := apiClient.Get("user", nil, nil)
	if err != nil {
		return nil, err
	}
	resBody := &models.ApiUserResponse{}
	err = api.UnmarshalResponse(res, resBody)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.HttpStatus(res.StatusCode).New(fmt.Sprintf("unexpected status code: %d", res.StatusCode))
	}
	return nil, nil
}

//TODO Please modify the folowing code to adapt to your plugin
/*
POST /plugins/Googlespreadsheet/connections
{
	"name": "Googlespreadsheet data connection name",
	"endpoint": "Googlespreadsheet api endpoint, i.e. https://example.com",
	"username": "username, usually should be email address",
	"password": "Googlespreadsheet api access token"
}
*/
func PostConnections(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	// update from request and save to database
	connection := &models.GooglespreadsheetConnection{}
	err := connectionHelper.Create(connection, input)
	if err != nil {
		return nil, err
	}
	return &plugin.ApiResourceOutput{Body: connection, Status: http.StatusOK}, nil
}

//TODO Please modify the folowing code to adapt to your plugin
/*
PATCH /plugins/Googlespreadsheet/connections/:connectionId
{
	"name": "Googlespreadsheet data connection name",
	"endpoint": "Googlespreadsheet api endpoint, i.e. https://example.com",
	"username": "username, usually should be email address",
	"password": "Googlespreadsheet api access token"
}
*/
func PatchConnection(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	connection := &models.GooglespreadsheetConnection{}
	err := connectionHelper.Patch(connection, input)
	if err != nil {
		return nil, err
	}
	return &plugin.ApiResourceOutput{Body: connection}, nil
}

/*
DELETE /plugins/Googlespreadsheet/connections/:connectionId
*/
func DeleteConnection(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	connection := &models.GooglespreadsheetConnection{}
	err := connectionHelper.First(connection, input.Params)
	if err != nil {
		return nil, err
	}
	err = connectionHelper.Delete(connection)
	return &plugin.ApiResourceOutput{Body: connection}, err
}

/*
GET /plugins/Googlespreadsheet/connections
*/
func ListConnections(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	var connections []models.GooglespreadsheetConnection
	err := connectionHelper.List(&connections)
	if err != nil {
		return nil, err
	}
	return &plugin.ApiResourceOutput{Body: connections, Status: http.StatusOK}, nil
}

//TODO Please modify the folowing code to adapt to your plugin
/*
GET /plugins/Googlespreadsheet/connections/:connectionId
{
	"name": "Googlespreadsheet data connection name",
	"endpoint": "Googlespreadsheet api endpoint, i.e. https://merico.atlassian.net/rest",
	"username": "username, usually should be email address",
	"password": "Googlespreadsheet api access token"
}
*/
func GetConnection(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	connection := &models.GooglespreadsheetConnection{}
	err := connectionHelper.First(connection, input.Params)
	return &plugin.ApiResourceOutput{Body: connection}, err
}