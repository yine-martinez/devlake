package models

import (
	"time"

	"github.com/apache/incubator-devlake/core/models/common"
)

type GoogleSpreadSheet struct {
	Team           string `gorm:"primaryKey"`
	Sprint         string `gorm:"primaryKey"`
	Tribe          string
	Q              string
	Dates          string `gorm:"type:varchar(255)"`
	Throughput     float64
	LeadTime       float64
	CycleTime      float64
	FlowEfficiency float64
	StartSprint		time.Time
	EndSprint		time.Time	
	common.NoPKModel
}

func (GoogleSpreadSheet) TableName() string {
	return "_tool_google_spreadsheet"
}
