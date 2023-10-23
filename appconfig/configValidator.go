package config

import (
	"github.com/pkg/errors"
)

// database code. Need to map to the database code (DataStoreConfig) in the configuration yaml file.
const (
	SQLDB      string = "sqldb"
	COUCHDB    string = "couch"
	CACHE_GRPC string = "cacheGrpc"
	USER_GRPC  string = "userGrpc"
)

// use case code. Need to map to the use case code (UseCaseConfig) in the configuration yaml file.
// Client app use those to retrieve use case from the container
const (
	TASK string = "task"
)

// data service code. Need to map to the data service code (DataConfig) in the configuration yaml file.
const (
	Task_DATA string = "taskData"
	TX_DATA   string = "txData"
)

func validateConfig(appConfig AppConfig) error {
	err := validateDataStore(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	useCase := appConfig.UseCase
	err = validateUseCase(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func validateDataStore(appConfig AppConfig) error {
	sc := appConfig.SQLConfig
	key := sc.Code
	scMsg := " in validateDataStore doesn't match key = "
	if SQLDB != key {
		errMsg := SQLDB + scMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateUseCase(useCase UseCaseConfig) error {
	err := validateTask(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func validateTask(useCaseConfig UseCaseConfig) error {
	tc := useCaseConfig.Task
	key := tc.Code
	tcMsg := " in validateTask doesn't match key = "
	if TASK != key {
		errMsg := TASK + tcMsg + key
		return errors.New(errMsg)
	}
	key = tc.TaskDataConfig.Code
	if Task_DATA != key {
		errMsg := Task_DATA + tcMsg + key
		return errors.New(errMsg)
	}
	key = tc.TxDataConfig.Code
	if TX_DATA != key {
		errMsg := TX_DATA + tcMsg + key
		return errors.New(errMsg)
	}
	return nil
}
