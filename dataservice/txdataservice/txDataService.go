package txdataservice

import (
	"github.com/0x-chaitu/gocrudMicro/dataservice"
	"github.com/0x-chaitu/gocrudMicro/tool/gdbc"
)

type TxDataSql struct {
	DB gdbc.SqlGdbc
}

func (tds *TxDataSql) TxEnd(txFunc func() error) error {
	return tds.DB.TxEnd(txFunc)
}

func (tds *TxDataSql) TxBegin() (dataservice.TxDataInterface, error) {
	sqltx, error := tds.DB.TxBegin()
	tdi := TxDataSql{DB: sqltx}
	tds.DB = tdi.DB
	return &tdi, error
}

func (tds *TxDataSql) GetTx() gdbc.SqlGdbc {
	return tds.DB
}
