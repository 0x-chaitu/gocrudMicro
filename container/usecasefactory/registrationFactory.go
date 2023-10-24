package usecasefactory

import (
	config "github.com/0x-chaitu/gocrudMicro/appconfig"
	"github.com/0x-chaitu/gocrudMicro/container"
	"github.com/0x-chaitu/gocrudMicro/usecase/registration"
	"github.com/pkg/errors"
)

type RegistrationFactory struct{}

func (rf *RegistrationFactory) Build(c container.Container, appConfig *config.AppConfig, key string) (UseCaseInterface, error) {
	uc := appConfig.UseCase.Task
	udi, err := buildTaskData(c, &uc.TaskDataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	tdi, err := buildTxData(c, &uc.TxDataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	tuc := registration.RegistrationTaskUseCase{TaskDataInterface: udi, TxDataInterface: tdi}
	return &tuc, nil
}
