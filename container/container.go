// package container uses dependency injection to create concrete type
// concrete type is a data type not an interface
package container

type Container interface {
	//appconfig will be cached in container
	//needs to be called again if the configration file changes
	InitApp(fileName string) error

	//buildusecase creates concrete type for each usecase
	//each call will create new instance so not singelton -> only data store handlers are singleton as they are cached in container
	BuildUseCase(code string) (interface{}, error)

	//only data store handler can be returned as its the only cached type
	Get(code string) (interface{}, bool)

	//only data store handler is put inside the container
	Put(code string, value interface{})
}
