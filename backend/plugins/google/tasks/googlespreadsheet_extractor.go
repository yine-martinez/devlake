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
	"fmt"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/google/models"
	"strconv"
	"strings"
)

var _ plugin.SubTaskEntryPoint = ExtractGooglespreadsheet

func ExtractGooglespreadsheet(taskCtx plugin.SubTaskContext) errors.Error {
	extractor, err := helper.NewApiExtractor(helper.ApiExtractorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx:    taskCtx,
			Params: GoogleApiParams{},
			Table:  RAW_SPREADSHEET_TABLE,
		},
		Extract: func(resData *helper.RawData) ([]interface{}, errors.Error) {
			extractedModels := make([]interface{}, 0)
			extractedData := make([]interface{}, 0)
			json.Unmarshal(resData.Data, &extractedData)
			for _, value := range extractedData {
				data := &response{}
				b, _ := json.Marshal(value)
				json.Unmarshal(b, &data)
				sprintInt, _ := strconv.Atoi(data.Sprint)
				t, _ := strconv.ParseFloat(data.Throughput, 8)
				l, _ := strconv.ParseFloat(data.LeadTime, 8)
				c, _ := strconv.ParseFloat(data.CycleTime, 8)
				data.FlowEfficiency = strings.Replace(data.FlowEfficiency, ",", ".", -1)
				data.FlowEfficiency = strings.Replace(data.FlowEfficiency, "%", "", -1)
				f, err := strconv.ParseFloat(data.FlowEfficiency, 8)
				if err != nil {
					fmt.Println(err)
				}
				extractedModels = append(extractedModels, &models.GoogleSpreadSheet{
					Team:           data.Team,
					Sprint:         sprintInt,
					Tribe:          data.Tribe,
					Q:              data.Q,
					Dates:          data.Dates,
					Throughput:     t,
					LeadTime:       l,
					CycleTime:      c,
					FlowEfficiency: f,
				})
			}

			return extractedModels, nil
		},
	})
	if err != nil {
		return err
	}

	return extractor.Execute()
}

type response struct {
	Team           string
	Sprint         string
	Tribe          string
	Q              string
	Dates          string
	Throughput     string
	LeadTime       string
	CycleTime      string
	FlowEfficiency string
}

func (r *response) UnmarshalJSON(b []byte) error {
	a := []interface{}{&r.Team, &r.Sprint, &r.Tribe, &r.Q, &r.Dates, &r.Throughput, &r.LeadTime, &r.CycleTime, &r.FlowEfficiency}
	return json.Unmarshal(b, &a)
}

var ExtractGooglespreadsheetMeta = plugin.SubTaskMeta{
	Name:             "ExtractGooglespreadsheet",
	EntryPoint:       ExtractGooglespreadsheet,
	EnabledByDefault: true,
	Description:      "Extract raw data into tool layer table google_googlespreadsheet",
}
