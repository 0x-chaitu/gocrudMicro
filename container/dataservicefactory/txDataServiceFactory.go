package dataservicefactory

import (
	"log"

	config "github.com/0x-chaitu/gocrudMicro/appconfig"
	"github.com/0x-chaitu/gocrudMicro/container"
	"github.com/0x-chaitu/gocrudMicro/container/datastorefactory"
	"github.com/0x-chaitu/gocrudMicro/dataservice/txdataservice"
	"github.com/0x-chaitu/gocrudMicro/tool/gdbc"
	"github.com/pkg/errors"
)

// receiver for Build method
type txDataServiceFactory struct{}

func (tsdf *txDataServiceFactory) Build(c container.Container, dataConfig *config.DataConfig) (DataServiceInterface, error) {
	log.Println("txDataServiceFactory")
	dsc := dataConfig.DataStoreConfig
	dsi, err := datastorefactory.GetDataStoreFb(dsc.Code).Build(c, &dataConfig.DataStoreConfig)
	if err != nil {
		errors.Wrap(err, "")
	}
	ds := dsi.(gdbc.SqlGdbc)
	tdm := txdataservice.TxDataSql{DB: ds}
	log.Println(tdm.DB)
	return &tdm, nil
}
