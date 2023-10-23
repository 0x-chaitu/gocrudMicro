//entry point for application's business logic
//should only depend on models package
//only cmd package should depend on this package

package usecase

import "github.com/0x-chaitu/gocrudMicro/model"

type RegistrationTaskUseCaseInterface interface {
	// Registers a task to db
	RegisterTask(task *model.Task) (resultTask *model.Task, err error)
}

// transaction interface for the usecase layer
type EnableTxer interface {
	EnableTx()
}
