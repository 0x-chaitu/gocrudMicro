package dataservicefactory

import (
	config "github.com/0x-chaitu/gocrudMicro/appconfig"
	"github.com/0x-chaitu/gocrudMicro/container"
)

// each data service needs a seperate builder
var dsFbMap = map[string]dataserviceFbInterface{
	config.Task_DATA: &taskDataServiceFactoryWrapper{},
	config.TX_DATA:   &txDataServiceFactory{},
}

// Build methods return type
type DataServiceInterface interface{}

type dataserviceFbInterface interface {
	Build(container.Container, *config.DataConfig) (DataServiceInterface, error)
}

func GetDataServiceFb(key string) dataserviceFbInterface {
	return dsFbMap[key]
}
