package archived

import (
	"time"

	"github.com/apache/incubator-devlake/core/models/migrationscripts/archived"
)

type GoogleSpreadSheet struct {
	Team           string `gorm:"primaryKey;type:varchar(100)"`
	Sprint         string `gorm:"primaryKey"`
	Tribe          string `gorm:"type:varchar(100)"`
	Q              string `gorm:"type:varchar(100)"`
	Dates          string `gorm:"type:varchar(255)"`
	Throughput     float64
	LeadTime       float64
	CycleTime      float64
	FlowEfficiency float64
	StartSprint		time.Time
	EndSprint		time.Time
	archived.NoPKModel
}

func (GoogleSpreadSheet) TableName() string {
	return "_tool_google_spreadsheet"
}
