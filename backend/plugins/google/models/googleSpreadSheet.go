package models

import "github.com/apache/incubator-devlake/core/models/common"

type GoogleSpreadSheet struct {
	common.NoPKModel
	team           string
	sprint         int
	tribe          string
	q              string
	throughput     float64
	leadTime       float64
	cycleTime      float64
	flowEfficiency float64
}

func (GoogleSpreadSheet) TableName() string {
	return "_tool_google_spreadSheet"
}
