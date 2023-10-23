package usecasefactory

import (
	config "github.com/0x-chaitu/gocrudMicro/appconfig"
	"github.com/0x-chaitu/gocrudMicro/container"
)

// use map to build factories for each use case
var UseCaseFactoryBuilderMap = map[string]UseCaseFbInterface{
	"task": &RegistrationFactory{},
}

type UseCaseInterface interface{}

// every factory needs to implement the builder pattern
type UseCaseFbInterface interface {
	Build(c container.Container, appConfig *config.AppConfig, key string) (UseCaseInterface, error)
}

func GetUseCaseFb(key string) UseCaseFbInterface {
	return UseCaseFactoryBuilderMap[key]
}
