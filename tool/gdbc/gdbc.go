package gdbc

import "database/sql"

type IGdbc interface {
	SqlGdbc
}

type SqlGdbc interface {
	Exec(query string, args ...interface{}) (sql.Result, error)

	Prepare(query string) (*sql.Stmt, error)

	Trasactioner
}

type Trasactioner interface {
	Rollback() error
	Commit() error
	TxEnd(txFunc func() error) error

	//TxBegin gets *sql.DB from the reciever and return SqlGdbc which has a *sql.Tx
	TxBegin() (SqlGdbc, error)
}
