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
	/*
			data := taskCtx.GetData().(*GoogleTaskData)
			//rawDataSubTaskArgs, data := helper.RawDataSubTaskArgs{} RawCreateRawDataSubTaskArgs(taskCtx, RAW_SPREADSHEET_TABLE)
			rawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx:    taskCtx,
			Params: GoogleApiParams{},
			Table:  RAW_SPREADSHEET_TABLE,
		},
			//logger := taskCtx.GetLogger()

			collectorWithState, err := helper.NewApiCollectorWithState(*rawDataSubTaskArgs, data.CreatedDateAfter)
			if err != nil {
				return err
			}
			//incremental := collectorWithState.IsIncremental()

			err = collectorWithState.InitCollector(helper.ApiCollectorArgs{
				//Incremental: incremental,
				ApiClient: data.ApiClient,
				// PageSize:    100,
				// TODO write which api would you want request
				//UrlTemplate: "https://script.googleapis.com/v1/scripts/AKfycbxq_dvZPKTsd9rbR8QZnBcsBvmHiNqCOtUscEXOUxM3JaBBNqC5v4gCs-RMlUJVUqJZrw:run",
				UrlTemplate: "http://localhost:8090/simple",
				Query: func(reqData *helper.RequestData) (url.Values, errors.Error) {
					query := url.Values{}
					//TODO Check if neccesary

						input := reqData.Input.(*helper.DatePair)
						query.Set("start_time", strconv.FormatInt(input.PairStartTime.Unix(), 10))
						query.Set("end_time", strconv.FormatInt(input.PairEndTime.Unix(), 10))


					return query, nil
				},
				ResponseParser: func(res *http.Response) ([]json.RawMessage, errors.Error) {
					// TODO decode result from api request
					print(res)
					return []json.RawMessage{}, nil
				},
			})
			if err != nil {
				return err
			}
			return collectorWithState.Execute()

	*/
	data := taskCtx.GetData().(*GoogleTaskData)

	collector, err := api.NewApiCollector(api.ApiCollectorArgs{
		RawDataSubTaskArgs: api.RawDataSubTaskArgs{
			Ctx:    taskCtx,
			Params: GoogleApiParams{},
			Table:  RAW_SPREADSHEET_TABLE,
		},
		ApiClient:   data.ApiClient,
		Incremental: false,

		UrlTemplate: "/spreadsheets/1TZk0LhUxfhIoRaVMHvOaE5M5iM1uFxXcddUXHMcIKXk/values/B2:J185",

		Query: func(reqData *api.RequestData) (url.Values, errors.Error) {
			query := url.Values{}
			return query, nil
		},
		ResponseParser: func(res *http.Response) ([]json.RawMessage, errors.Error) {
			println(res)

			body := &struct {
				Range string          `json:"range"`
				MajorDimension  string `json:"majorDimension"`
				Values json.RawMessage `json:"values"`
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
