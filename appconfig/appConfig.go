package config

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// represents app config
type AppConfig struct {
	SQLConfig DataStoreConfig `yaml:"sqlConfig"`
	UseCase   UseCaseConfig   `yaml:"useCaseConfig"`
}

// different useCases
type UseCaseConfig struct {
	Task TaskConfig `yaml:"task"`
}

// task useCase
type TaskConfig struct {
	Code           string     `yaml:"code"`
	TaskDataConfig DataConfig `yaml:"taskDataConfig"`
	TxDataConfig   DataConfig `yaml:"txDataConfig"`
}

// represents dataservice
type DataConfig struct {
	Code            string          `yaml:"code"`
	DataStoreConfig DataStoreConfig `yaml:"dataStoreConfig"`
}

type DataStoreConfig struct {
	Code       string `yaml:"code"`
	DriverName string `yaml:"driverName"`
	UrlAdress  string `yaml:"urlAdress"`
	DbName     string `yaml:"dbName"`
}

func ReadConfig(filename string) (*AppConfig, error) {
	var ac AppConfig
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "read error")
	}
	err = yaml.Unmarshal(file, &ac)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}
	err = validateConfig(ac)
	if err != nil {
		return nil, errors.Wrap(err, "validation error")

	}
	fmt.Println("appconfig:", ac)
	return &ac, nil
}
