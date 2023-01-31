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

package impl

import (
	"fmt"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
    "github.com/apache/incubator-devlake/plugins/googlespreadsheet/api"
    "github.com/apache/incubator-devlake/plugins/googlespreadsheet/models"
    "github.com/apache/incubator-devlake/plugins/googlespreadsheet/models/migrationscripts"
	"github.com/apache/incubator-devlake/plugins/googlespreadsheet/tasks"

	"github.com/apache/incubator-devlake/core/context"
)

// make sure interface is implemented
var _ plugin.PluginMeta = (*Googlespreadsheet)(nil)
var _ plugin.PluginInit = (*Googlespreadsheet)(nil)
var _ plugin.PluginTask = (*Googlespreadsheet)(nil)
var _ plugin.PluginApi = (*Googlespreadsheet)(nil)
var _ plugin.PluginBlueprintV100 = (*Googlespreadsheet)(nil)
var _ plugin.CloseablePluginTask = (*Googlespreadsheet)(nil)



type Googlespreadsheet struct{}

func (p Googlespreadsheet) Description() string {
	return "collect some Googlespreadsheet data"
}

func (p Googlespreadsheet) Init(br context.BasicRes) errors.Error {
	api.Init(br)
	return nil
}

func (p Googlespreadsheet) SubTaskMetas() []plugin.SubTaskMeta {
	// TODO add your sub task here
	return []plugin.SubTaskMeta{
	}
}

func (p Googlespreadsheet) PrepareTaskData(taskCtx plugin.TaskContext, options map[string]interface{}) (interface{}, errors.Error) {
	op, err := tasks.DecodeAndValidateTaskOptions(options)
    if err != nil {
        return nil, err
    }
    connectionHelper := helper.NewConnectionHelper(
        taskCtx,
        nil,
    )
    connection := &models.GooglespreadsheetConnection{}
    err = connectionHelper.FirstById(connection, op.ConnectionId)
    if err != nil {
        return nil, errors.Default.Wrap(err, "unable to get Googlespreadsheet connection by the given connection ID")
    }

    apiClient, err := tasks.NewGooglespreadsheetApiClient(taskCtx, connection)
    if err != nil {
        return nil, errors.Default.Wrap(err, "unable to get Googlespreadsheet API client instance")
    }

    return &tasks.GooglespreadsheetTaskData{
        Options:   op,
        ApiClient: apiClient,
    }, nil
}

// PkgPath information lost when compiled as plugin(.so)
func (p Googlespreadsheet) RootPkgPath() string {
	return "github.com/apache/incubator-devlake/plugins/googlespreadsheet"
}

func (p Googlespreadsheet) MigrationScripts() []plugin.MigrationScript {
	return migrationscripts.All()
}

func (p Googlespreadsheet) ApiResources() map[string]map[string]plugin.ApiResourceHandler {
    return map[string]map[string]plugin.ApiResourceHandler{
        "test": {
            "POST": api.TestConnection,
        },
        "connections": {
            "POST": api.PostConnections,
            "GET":  api.ListConnections,
        },
        "connections/:connectionId": {
            "GET":    api.GetConnection,
            "PATCH":  api.PatchConnection,
            "DELETE": api.DeleteConnection,
        },
    }
}

func (p Googlespreadsheet) MakePipelinePlan(connectionId uint64, scope []*plugin.BlueprintScopeV100) (plugin.PipelinePlan, errors.Error) {
	return api.MakePipelinePlan(p.SubTaskMetas(), connectionId, scope)
}

func (p Googlespreadsheet) Close(taskCtx plugin.TaskContext) errors.Error {
	data, ok := taskCtx.GetData().(*tasks.GooglespreadsheetTaskData)
	if !ok {
		return errors.Default.New(fmt.Sprintf("GetData failed when try to close %+v", taskCtx))
	}
	data.ApiClient.Release()
	return nil
}
