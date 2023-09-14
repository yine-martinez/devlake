package migrationscripts

import (
	"github.com/apache/incubator-devlake/core/context"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/plugins/google/models/migrationscripts/archived"
)

type addSpreadsheet struct {
}

func (script *addSpreadsheet) Up(basicRes context.BasicRes) errors.Error {
	return basicRes.GetDal().AutoMigrate(&archived.GoogleSpreadSheet{})
}

func (*addSpreadsheet) Version() uint64 {
	return 20230213114400
}

func (*addSpreadsheet) Name() string {
	return "create _tool_google_spreadsheet"
}
