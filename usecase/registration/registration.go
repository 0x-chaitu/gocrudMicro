package registration

import (
	"github.com/0x-chaitu/gocrudMicro/dataservice"
	"github.com/0x-chaitu/gocrudMicro/model"
	"github.com/pkg/errors"
)

type RegistrationTaskUseCase struct {
	TaskDataInterface dataservice.TaskDataInterface
	TxDataInterface   dataservice.TxDataInterface
}

func (ruc *RegistrationTaskUseCase) RegisterTask(task *model.Task) (*model.Task, error) {
	err := task.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "task validation failed")
	}
	resultTask, err := ruc.TaskDataInterface.Insert(task)
	if err != nil {
		return nil, errors.Wrapf(err, "task insertion failed")
	}
	return resultTask, nil
}
