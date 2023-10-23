package main

import (
	"fmt"
	"log"
	"time"

	config "github.com/0x-chaitu/gocrudMicro/appconfig"
	"github.com/0x-chaitu/gocrudMicro/container"
	"github.com/0x-chaitu/gocrudMicro/container/servicecontainer"
	"github.com/0x-chaitu/gocrudMicro/model"
	"github.com/0x-chaitu/gocrudMicro/usecase"
	"github.com/pkg/errors"
)

const (
	DEV_CONFIG string = "/home/chaitu/Projects/goCrud/appconfig/appConfigDev.yaml"
)

func main() {
	testPostgreSql()
}

func testPostgreSql() {
	filename := DEV_CONFIG
	container, err := buildContainer(filename)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	testRegisterTask(container)
}

func testRegisterTask(container container.Container) {
	ruci, err := getRegistrationUseCase(container)
	if err != nil {
		log.Printf("registration interface build failed:%+v\n", err)
	}
	created := time.Now().Format("2018-9-12")

	task := model.Task{Name: "Brian", Created: created}

	resultUser, err := ruci.RegisterTask(&task)
	if err != nil {
		log.Fatalf("user registration failed:%+v\n", err)
	} else {
		log.Fatalf("new user registered: %v", resultUser)
	}
}

func getRegistrationUseCase(c container.Container) (usecase.RegistrationTaskUseCaseInterface, error) {
	key := config.TASK
	value, err := c.BuildUseCase(key)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return value.(usecase.RegistrationTaskUseCaseInterface), nil

}

func buildContainer(filename string) (container.Container, error) {
	factoryMap := make(map[string]interface{})
	appConfig := config.AppConfig{}
	container := servicecontainer.ServiceContainer{FactoryMap: factoryMap, AppConfig: &appConfig}

	err := container.InitApp(filename)
	if err != nil {
		//logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	return &container, nil
}
