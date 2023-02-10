package archived

import "github.com/apache/incubator-devlake/core/models/migrationscripts/archived"

type GoogleSpreadSheet struct {
	team           string
	sprint         int
	tribe          string
	q              string
	throughput     float64
	leadTime       float64
	cycleTime      float64
	flowEfficiency float64
	archived.NoPKModel
}

func (GoogleSpreadSheet) TableName() string {
	return "_tool_google_spreadSheet"
}
