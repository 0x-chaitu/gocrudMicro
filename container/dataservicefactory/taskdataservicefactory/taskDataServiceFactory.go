package taskdataservicefactory

import (
	config "github.com/0x-chaitu/gocrudMicro/appconfig"
	"github.com/0x-chaitu/gocrudMicro/container"
	"github.com/0x-chaitu/gocrudMicro/dataservice"
)

var udsFbMap = map[string]taskDataServiceFbInterface{
	config.SQLDB: &sqlTaskDataServiceFactory{},
}

// The builder interface for factory method pattern
// Every factory needs to implement Build method
type taskDataServiceFbInterface interface {
	Build(container.Container, *config.DataConfig) (dataservice.TaskDataInterface, error)
}

// GetDataServiceFb is accessors for factoryBuilderMap
func GetUserDataServiceFb(key string) taskDataServiceFbInterface {
	return udsFbMap[key]
}
