package databasehandler

import "database/sql"

// concrete implementation of sqlGdbc using Db
type SqlDbTx struct {
	DB *sql.DB
}

// concrete implementation of sqlGdbc using Tx
type SqlDBConn struct {
	DB *sql.Tx
}

func (sdt *SqlDbTx) Prepare(query string) (*sql.Stmt, error) {
	return sdt.DB.Prepare(query)
}

func (sdt *SqlDbTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return sdt.DB.Exec(query, args...)
}

func (sdb *SqlDBConn) Prepare(query string) (*sql.Stmt, error) {
	return sdb.DB.Prepare(query)
}

func (sdb *SqlDBConn) Exec(query string, args ...interface{}) (sql.Result, error) {
	return sdb.DB.Exec(query, args...)
}
