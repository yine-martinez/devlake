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
	"github.com/shopspring/decimal"
	"strconv"
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
			println("-----------------")
			json.Unmarshal(resData.Data, &extractedData)
			// TODO decode some db models from api result
			println("-----------------")

			for _, value := range extractedData {
				fmt.Println(value)
				data := &response{}
				b, _ := json.Marshal(value)
				json.Unmarshal(b, &data)
				fmt.Println(data.Team)
				fmt.Println(data.Sprint)
				sprintInt, _ := strconv.Atoi(data.Sprint)
				t, error := decimal.NewFromString(data.Throughput)
				if error != nil {
					t = decimal.NewFromInt(0)
				}
				l, error := decimal.NewFromString(data.LeadTime)
				if error != nil {
					l = decimal.NewFromInt(0)
				}
				c, error := decimal.NewFromString(data.CycleTime)
				if data.CycleTime != "" {
					c = decimal.NewFromInt(0)
				}
				f, error := decimal.NewFromString(data.FlowEfficiency)
				if error != nil {
					f = decimal.NewFromInt(0)
				}
				extractedModels = append(extractedModels, &models.GoogleSpreadSheet{
					Team:           data.Team,
					Sprint:         sprintInt,
					Tribe:          data.Tribe,
					Q:              data.Q,
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
	Throughput     string
	LeadTime       string
	CycleTime      string
	FlowEfficiency string
}

func (r *response) UnmarshalJSON(b []byte) error {
	a := []interface{}{&r.Team, &r.Sprint, &r.Tribe, &r.Q, &r.Throughput, &r.LeadTime, &r.CycleTime, &r.FlowEfficiency}
	return json.Unmarshal(b, &a)
}

var ExtractGooglespreadsheetMeta = plugin.SubTaskMeta{
	Name:             "ExtractGooglespreadsheet",
	EntryPoint:       ExtractGooglespreadsheet,
	EnabledByDefault: true,
	Description:      "Extract raw data into tool layer table google_googlespreadsheet",
}