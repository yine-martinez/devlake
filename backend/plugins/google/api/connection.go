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
	"github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/google/models"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"time"
)

// TODO Please modify the following code to fit your needs
func TestConnection(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {

	var err errors.Error
	var connection models.GoogleConnection
	errorDecode := mapstructure.Decode(input.Body, &connection)
	if errorDecode != nil {
		return nil, errors.Convert(errorDecode)
	}
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %v", connection.Token),
	}
	// test connection
	apiClient, err := api.NewApiClient(
		context.TODO(),
		connection.GetEndpoint(),
		headers,
		20*time.Second,
		"",
		basicRes)
	if err != nil {
		return nil, err
	}
	res, err := apiClient.Get("spreadsheets/"+connection.SpreadsheetID+"/values/"+connection.FirstValue+":"+connection.LastValue,
		nil, nil)
	if err != nil {
		return nil, err
	}

	resBody := &models.GoogleSpreadSheet{}
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
POST /plugins/Google/connections
{
	"name": "Google data connection name",
	"endpoint": "Google api endpoint, i.e. https://example.com",
	"username": "username, usually should be email address",
	"password": "Google api access token"
}
*/
func PostConnections(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	// update from request and save to database
	connection := &models.GoogleConnection{}
	err := connectionHelper.Create(connection, input)
	if err != nil {
		return nil, err
	}
	return &plugin.ApiResourceOutput{Body: connection, Status: http.StatusOK}, nil
}

//TODO Please modify the folowing code to adapt to your plugin
/*
PATCH /plugins/Google/connections/:connectionId
{
	"name": "Google data connection name",
	"endpoint": "Google api endpoint, i.e. https://example.com",
	"username": "username, usually should be email address",
	"password": "Google api access token"
}
*/
func PatchConnection(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	connection := &models.GoogleConnection{}
	err := connectionHelper.Patch(connection, input)
	if err != nil {
		return nil, err
	}
	return &plugin.ApiResourceOutput{Body: connection}, nil
}

/*
DELETE /plugins/Google/connections/:connectionId
*/
func DeleteConnection(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	connection := &models.GoogleConnection{}
	err := connectionHelper.First(connection, input.Params)
	if err != nil {
		return nil, err
	}
	err = connectionHelper.Delete(connection)
	return &plugin.ApiResourceOutput{Body: connection}, err
}

/*
GET /plugins/Google/connections
*/
func ListConnections(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	fmt.Println("here")
	var connections []models.GoogleConnection
	err := connectionHelper.List(&connections)
	if err != nil {
		return nil, err
	}
	return &plugin.ApiResourceOutput{Body: connections, Status: http.StatusOK}, nil
}

//TODO Please modify the folowing code to adapt to your plugin
/*
GET /plugins/Google/connections/:connectionId
{
	"name": "Google data connection name",
	"endpoint": "Google api endpoint, i.e. https://merico.atlassian.net/rest",
	"username": "username, usually should be email address",
	"password": "Google api access token"
}
*/
func GetConnection(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	connection := &models.GoogleConnection{}
	err := connectionHelper.First(connection, input.Params)
	return &plugin.ApiResourceOutput{Body: connection}, err
}
