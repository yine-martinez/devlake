package models

import (
	"github.com/apache/incubator-devlake/core/models/common"
	"github.com/shopspring/decimal"
)

type GoogleSpreadSheet struct {
	Team           string `gorm:"primaryKey"`
	Sprint         int    `gorm:"primaryKey"`
	Tribe          string
	Q              string
	Throughput     decimal.Decimal `gorm:"type:decimal(3,1)"`
	LeadTime       decimal.Decimal `gorm:"type:decimal(3,1)"`
	CycleTime      decimal.Decimal `gorm:"type:decimal(3,1)"`
	FlowEfficiency decimal.Decimal `gorm:"type:decimal(3,1)"`
	common.NoPKModel
}

func (GoogleSpreadSheet) TableName() string {
	return "_tool_google_spreadsheet"
}
