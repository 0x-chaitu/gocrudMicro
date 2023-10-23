package databasehandler

import (
	"log"

	"github.com/0x-chaitu/gocrudMicro/tool/gdbc"
)

func (sdt *SqlDbTx) Rollback() error {
	return nil
}

func (sdt *SqlDbTx) Commit() error {
	return nil
}

func (sdt *SqlDbTx) TxEnd(txFunc func() error) error {
	return nil
}

func (sdt *SqlDbTx) TxBegin() (gdbc.SqlGdbc, error) {
	log.Println("transaction begin")
	tx, err := sdt.DB.Begin()
	sct := SqlDBConn{tx}
	return &sct, err
}

func (sct *SqlDBConn) Rollback() error {
	return sct.DB.Rollback()
}

func (sct *SqlDBConn) Commit() error {
	return sct.DB.Commit()
}

func (sct *SqlDBConn) TxEnd(txFunc func() error) error {
	var err error
	tx := sct.DB

	defer func() {
		if p := recover(); p != nil {
			log.Printf("found panic %v and rollback:\n", p)
			tx.Rollback()
			panic(p)
		} else if err != nil {
			log.Printf("found error %v\n", err)
			tx.Rollback()
		} else {
			log.Println("commit:")
			tx.Commit()
		}
	}()
	err = txFunc()
	return err
}

func (sdt *SqlDBConn) TxBegin() (gdbc.SqlGdbc, error) {
	return nil, nil
}
