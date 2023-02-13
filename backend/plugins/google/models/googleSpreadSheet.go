package models

import (
	"github.com/apache/incubator-devlake/core/models/common"
	"github.com/shopspring/decimal"
)

type GoogleSpreadSheet struct {
	team           string `gorm:"primaryKey"`
	sprint         int    `gorm:"primaryKey"`
	tribe          string
	q              string
	throughput     decimal.Decimal `gorm:"type:decimal(3,1)"`
	leadTime       decimal.Decimal `gorm:"type:decimal(3,1)"`
	cycleTime      decimal.Decimal `gorm:"type:decimal(3,1)"`
	flowEfficiency decimal.Decimal `gorm:"type:decimal(3,1)"`

	common.NoPKModel
}

func (GoogleSpreadSheet) TableName() string {
	return "_tool_google_spreadSheet"
}
