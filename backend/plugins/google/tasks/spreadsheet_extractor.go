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
	"strconv"
	"strings"
	"time"

	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/google/models"
)

var _ plugin.SubTaskEntryPoint = ExtractGooglespreadsheet

func ExtractGooglespreadsheet(taskCtx plugin.SubTaskContext) errors.Error {
	logger := taskCtx.GetLogger()
	logger.Info("extract data from google spreadsheet")
	data := taskCtx.GetData().(*GoogleTaskData)
	extractor, err := helper.NewApiExtractor(helper.ApiExtractorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: GoogleApiParams{
				SpreadsheetID: data.SpreadsheetID,
				FirstValue:    data.FirstValue,
				LastValue:     data.LastValue,
			},
			Table: RAW_SPREADSHEET_TABLE,
		},
		Extract: func(resData *helper.RawData) ([]interface{}, errors.Error) {
			extractedModels := make([]interface{}, 0)
			extractedData := make([]interface{}, 0)
			errUnmarshal := json.Unmarshal(resData.Data, &extractedData)
			if errUnmarshal != nil {
				logger.Error(errUnmarshal, "error unmarshalling json")
			}
			for _, line := range extractedData {
				data := &spreadSheetStructure{}
				rawData, errMarshal := json.Marshal(line)
				if errMarshal != nil {
					logger.Error(errUnmarshal, "error marshalling json")
				}
				errUnmarshal = json.Unmarshal(rawData, &data)
				if errUnmarshal != nil {
					logger.Error(errUnmarshal, "error unmarshalling rawData json")
				}

				formatedData, err := formatData(data)
				if err != nil {
					continue
				}

				extractedModels = append(extractedModels, formatedData)
			}

			return extractedModels, nil
		},
	})
	if err != nil {
		return err
	}

	return extractor.Execute()
}

type spreadSheetStructure struct {
	Team           string
	Sprint         string
	Tribe          string
	Q              string
	Dates          string
	Throughput     string
	LeadTime       string
	CycleTime      string
	FlowEfficiency string
	StartSprint    string
	EndSprint      string
}

func (r *spreadSheetStructure) UnmarshalJSON(b []byte) error {
	a := []interface{}{&r.Team, &r.Sprint, &r.Tribe, &r.Q, &r.Dates, &r.Throughput, &r.LeadTime, &r.CycleTime, &r.FlowEfficiency, &r.StartSprint, &r.EndSprint}
	return json.Unmarshal(b, &a)
}

var ExtractGooglespreadsheetMeta = plugin.SubTaskMeta{
	Name:             "ExtractGooglespreadsheet",
	EntryPoint:       ExtractGooglespreadsheet,
	EnabledByDefault: true,
	Description:      "Extract raw data into tool layer table google_googlespreadsheet",
}

func formatData(data *spreadSheetStructure) (*models.GoogleSpreadSheet, errors.Error) {
	formatedData := &models.GoogleSpreadSheet{}
	formatedData.Team = strings.TrimSpace(data.Team)
	formatedData.Q = strings.Replace(data.Q, ",", ".", -1)
	formatedData.Throughput, _ = strconv.ParseFloat(data.Throughput, 8) //nolint
	formatedData.LeadTime, _ = strconv.ParseFloat(data.LeadTime, 8)     //nolint
	formatedData.CycleTime, _ = strconv.ParseFloat(data.CycleTime, 8)   //nolint

	if data.FlowEfficiency != "" {
		flowEfficiency := strings.Replace(data.FlowEfficiency, ",", ".", -1)
		flowEfficiency = strings.Replace(data.FlowEfficiency, "%", "", -1)
		formatedData.FlowEfficiency, _ = strconv.ParseFloat(flowEfficiency, 8) //nolint
	} else {
		return formatedData, errors.BadInput.New("Could not format FlowEfficiency")
	}

	if data.StartSprint != "" {
		format := "2006-01-02"
		formatedData.StartSprint, _ = time.Parse(format, data.StartSprint)
		formatedData.EndSprint, _ = time.Parse(format, data.EndSprint)
	} else {
		return formatedData, errors.BadInput.New("Could not format StartSprint")
	}

	return formatedData, nil
}
