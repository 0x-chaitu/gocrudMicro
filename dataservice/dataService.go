package dataservice

import (
	"github.com/0x-chaitu/gocrudMicro/model"
	"github.com/0x-chaitu/gocrudMicro/tool/gdbc"
)

type TaskDataInterface interface {

	// insert task in db
	Insert(t *model.Task) (resultTask *model.Task, err error)

	EnableTxer
}

type EnableTxer interface {
	//enable underlying db handler sql.DB with sql.Tx
	EnableTx(dataInterface TxDataInterface)
}

type TxDataInterface interface {
	//start a transaction; gets a db handler and returns a txinterface
	//which has a *sql.Tx
	TxBegin() (TxDataInterface, error)

	//called at end of transaction, depending on error, either it commits
	//or rollbacks the transaction
	//txFunc is the business function
	TxEnd(txFunc func() error) error

	//return the underlying tx handler
	GetTx() gdbc.SqlGdbc
}
