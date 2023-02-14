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
	"github.com/apache/incubator-devlake/core/models/common"
	"github.com/apache/incubator-devlake/core/plugin"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/google/models"
	"github.com/shopspring/decimal"
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
				//var data *models.GoogleSpreadSheet
				//json.Unmarshal(v, data)
				dataModel := models.GoogleSpreadSheet{
					Team:           "",
					Sprint:         0,
					Tribe:          "",
					Q:              "",
					Throughput:     decimal.Decimal{},
					LeadTime:       decimal.Decimal{},
					CycleTime:      decimal.Decimal{},
					FlowEfficiency: decimal.Decimal{},
					NoPKModel:      common.NoPKModel{},
				}
				extractedModels = append(extractedModels, dataModel)
			}

			return extractedModels, nil
		},
	})
	if err != nil {
		return err
	}

	return extractor.Execute()
}

var ExtractGooglespreadsheetMeta = plugin.SubTaskMeta{
	Name:             "ExtractGooglespreadsheet",
	EntryPoint:       ExtractGooglespreadsheet,
	EnabledByDefault: true,
	Description:      "Extract raw data into tool layer table google_googlespreadsheet",
}
