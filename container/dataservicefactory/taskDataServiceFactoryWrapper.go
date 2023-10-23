package dataservicefactory

import (
	"log"

	config "github.com/0x-chaitu/gocrudMicro/appconfig"
	"github.com/0x-chaitu/gocrudMicro/container"
	"github.com/0x-chaitu/gocrudMicro/container/dataservicefactory/taskdataservicefactory"
	"github.com/pkg/errors"
)

type taskDataServiceFactoryWrapper struct{}

func (tdsfw *taskDataServiceFactoryWrapper) Build(c container.Container, dataConfig *config.DataConfig) (DataServiceInterface, error) {
	log.Println("task Data Service Factory")
	key := dataConfig.DataStoreConfig.Code
	udsi, err := taskdataservicefactory.GetUserDataServiceFb(key).Build(c, dataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return udsi, nil
}
