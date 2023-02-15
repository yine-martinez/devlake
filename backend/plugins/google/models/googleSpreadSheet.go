package models

import (
	"github.com/apache/incubator-devlake/core/models/common"
)

type GoogleSpreadSheet struct {
	Team           string `gorm:"primaryKey"`
	Sprint         int    `gorm:"primaryKey"`
	Tribe          string
	Q              string
	Dates          string `gorm:"type:varchar(255)"`
	Throughput     float64
	LeadTime       float64
	CycleTime      float64
	FlowEfficiency float64
	common.NoPKModel
}

func (GoogleSpreadSheet) TableName() string {
	return "_tool_google_spreadsheet"
}
