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
	"github.com/apache/incubator-devlake/core/context"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/google/api"
	"github.com/apache/incubator-devlake/plugins/google/models"
	"github.com/apache/incubator-devlake/plugins/google/models/migrationscripts"
	"github.com/apache/incubator-devlake/plugins/google/tasks"
	"time"
)

// make sure interface is implemented
var _ plugin.PluginMeta = (*Google)(nil)
var _ plugin.PluginInit = (*Google)(nil)
var _ plugin.PluginTask = (*Google)(nil)
var _ plugin.PluginApi = (*Google)(nil)
var _ plugin.PluginBlueprintV100 = (*Google)(nil)
var _ plugin.CloseablePluginTask = (*Google)(nil)

type Google struct{}

func (p Google) Description() string {
	return "collect some Google data"
}

func (p Google) Init(br context.BasicRes) errors.Error {
	api.Init(br)
	return nil
}

func (p Google) SubTaskMetas() []plugin.SubTaskMeta {
	// TODO add your sub task here
	return []plugin.SubTaskMeta{
		tasks.CollectSpreadsheetMeta,
		tasks.ExtractGooglespreadsheetMeta,
	}
}

func (p Google) PrepareTaskData(taskCtx plugin.TaskContext, options map[string]interface{}) (interface{}, errors.Error) {
	op, err := tasks.DecodeAndValidateTaskOptions(options)
	if err != nil {
		return nil, err
	}
	connectionHelper := helper.NewConnectionHelper(
		taskCtx,
		nil,
	)
	connection := &models.GoogleConnection{}
	err = connectionHelper.FirstById(connection, op.ConnectionId)
	if err != nil {
		return nil, errors.Default.Wrap(err, "unable to get Google connection by the given connection ID")
	}

	// TODO Check if here make sense this
	apiClient, err := tasks.NewGoogleApiClient(taskCtx, connection)
	if err != nil {
		return nil, errors.Default.Wrap(err, "unable to get Google API client instance")
	}
	taskData := &tasks.GoogleTaskData{
		Options:   op,
		ApiClient: apiClient,
	}
	var createdDateAfter time.Time
	if op.CreatedDateAfter != "" {
		createdDateAfter, err = errors.Convert01(time.Parse(time.RFC3339, op.CreatedDateAfter))
		if err != nil {
			return nil, errors.BadInput.Wrap(err, "invalid value for `createdDateAfter`")
		}
	}
	if !createdDateAfter.IsZero() {
		taskData.CreatedDateAfter = &createdDateAfter
		//logger.Debug("collect data updated createdDateAfter %s", createdDateAfter)
	}
	return taskData, nil
}

// PkgPath information lost when compiled as plugin(.so)
func (p Google) RootPkgPath() string {
	return "github.com/apache/incubator-devlake/plugins/google"
}

func (p Google) MigrationScripts() []plugin.MigrationScript {
	return migrationscripts.All()
}

func (p Google) ApiResources() map[string]map[string]plugin.ApiResourceHandler {
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

func (p Google) MakePipelinePlan(connectionId uint64, scope []*plugin.BlueprintScopeV100) (plugin.PipelinePlan, errors.Error) {
	return api.MakePipelinePlan(p.SubTaskMetas(), connectionId, scope)
}

func (p Google) Close(taskCtx plugin.TaskContext) errors.Error {
	data, ok := taskCtx.GetData().(*tasks.GoogleTaskData)
	if !ok {
		return errors.Default.New(fmt.Sprintf("GetData failed when try to close %+v", taskCtx))
	}
	data.ApiClient.Release()
	return nil
}
