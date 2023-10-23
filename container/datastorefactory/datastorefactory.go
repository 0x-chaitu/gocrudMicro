package datastorefactory

import (
	config "github.com/0x-chaitu/gocrudMicro/appconfig"
	"github.com/0x-chaitu/gocrudMicro/container"
)

// map database code to its interface builder
var dsFbMap = map[string]dsFbInterface{
	config.SQLDB: &sqlFactory{},
}

type DataStoreInterface interface{}

// builder interface for datstore factory
type dsFbInterface interface {
	Build(container.Container, *config.DataStoreConfig) (DataStoreInterface, error)
}

func GetDataStoreFb(key string) dsFbInterface {
	return dsFbMap[key]
}
