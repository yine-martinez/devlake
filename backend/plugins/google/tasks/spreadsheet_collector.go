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

package tasks

import (
	"encoding/json"
	"github.com/apache/incubator-devlake/core/errors"
	core "github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"net/http"
	"net/url"
)

const RAW_SPREADSHEET_TABLE = "google_spreadsheet"

var _ core.SubTaskEntryPoint = CollectSpreadsheet

func CollectSpreadsheet(taskCtx core.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*GoogleTaskData)

	collector, err := api.NewApiCollector(api.ApiCollectorArgs{
		RawDataSubTaskArgs: api.RawDataSubTaskArgs{
			Ctx:    taskCtx,
			Params: GoogleApiParams{},
			Table:  RAW_SPREADSHEET_TABLE,
		},
		ApiClient:   data.ApiClient,
		Incremental: false,

		UrlTemplate: "/spreadsheets/1TZk0LhUxfhIoRaVMHvOaE5M5iM1uFxXcddUXHMcIKXk/values/B3:J9999",

		Query: func(reqData *api.RequestData) (url.Values, errors.Error) {
			query := url.Values{}
			return query, nil
		},
		ResponseParser: func(res *http.Response) ([]json.RawMessage, errors.Error) {
			println(res)

			body := &struct {
				Range          string          `json:"range"`
				MajorDimension string          `json:"majorDimension"`
				Values         json.RawMessage `json:"values"`
			}{}
			err := api.UnmarshalResponse(res, body)
			if err != nil {
				return nil, err
			}
			println("receive data:", len(body.Values))
			return []json.RawMessage{body.Values}, nil
		},
	})
	if err != nil {
		return err
	}

	return collector.Execute()
}

var CollectSpreadsheetMeta = core.SubTaskMeta{
	Name:             "CollectSpreadsheet",
	EntryPoint:       CollectSpreadsheet,
	EnabledByDefault: true,
	Description:      "Collect Spreadsheet data from Google api",
}
