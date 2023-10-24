package taskdataservicefactory

import (
	"log"

	config "github.com/0x-chaitu/gocrudMicro/appconfig"
	"github.com/0x-chaitu/gocrudMicro/container"
	"github.com/0x-chaitu/gocrudMicro/container/datastorefactory"
	"github.com/0x-chaitu/gocrudMicro/dataservice"
	sqldb "github.com/0x-chaitu/gocrudMicro/dataservice/taskdata/postgresqldb"
	"github.com/0x-chaitu/gocrudMicro/tool/gdbc"
	"github.com/pkg/errors"
)

type sqlTaskDataServiceFactory struct{}

func (sudsf *sqlTaskDataServiceFactory) Build(c container.Container, dataConfig *config.DataConfig) (dataservice.TaskDataInterface, error) {
	log.Println("sqlTaskDataServiceFactory")
	dsc := dataConfig.DataStoreConfig
	dsi, err := datastorefactory.GetDataStoreFb(dsc.Code).Build(c, &dsc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	ds := dsi.(gdbc.SqlGdbc)
	tds := sqldb.TaskDataSql{DB: ds}
	log.Printf("uds: %v", tds.DB)
	return &tds, nil

}
