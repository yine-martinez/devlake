package archived

import (
	"github.com/shopspring/decimal"
)

type GoogleSpreadSheet struct {
	Team           string          `gorm:"primaryKey;type:varchar(100)"`
	Sprint         int             `gorm:"primaryKey"`
	Tribe          string          `gorm:"type:varchar(100)"`
	Q              string          `gorm:"type:varchar(100)"`
	Throughput     decimal.Decimal `gorm:"type:decimal(3,1)"`
	LeadTime       decimal.Decimal `gorm:"type:decimal(3,1)"`
	CycleTime      decimal.Decimal `gorm:"type:decimal(3,1)"`
	FlowEfficiency decimal.Decimal `gorm:"type:decimal(3,1)"`
}

func (GoogleSpreadSheet) TableName() string {
	return "_tool_google_spreadSheet"
}
