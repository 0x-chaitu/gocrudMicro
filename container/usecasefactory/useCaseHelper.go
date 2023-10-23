package usecasefactory

import (
	config "github.com/0x-chaitu/gocrudMicro/appconfig"
	"github.com/0x-chaitu/gocrudMicro/container"
	"github.com/0x-chaitu/gocrudMicro/container/dataservicefactory"
	"github.com/0x-chaitu/gocrudMicro/dataservice"
	"github.com/pkg/errors"
)

func buildTaskData(c container.Container, dc *config.DataConfig) (dataservice.TaskDataInterface, error) {
	dsi, err := dataservicefactory.GetDataServiceFb(dc.Code).Build(c, dc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	udi := dsi.(dataservice.TaskDataInterface)
	return udi, nil
}

func buildTxData(c container.Container, dc *config.DataConfig) (dataservice.TxDataInterface, error) {
	dsi, err := dataservicefactory.GetDataServiceFb(dc.Code).Build(c, dc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	tdi := dsi.(dataservice.TxDataInterface)
	return tdi, nil
}
