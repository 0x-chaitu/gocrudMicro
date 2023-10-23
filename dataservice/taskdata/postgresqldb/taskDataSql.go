package sqldb

import (
	"log"

	"github.com/0x-chaitu/gocrudMicro/dataservice"
	"github.com/0x-chaitu/gocrudMicro/model"
	"github.com/0x-chaitu/gocrudMicro/tool/gdbc"
	"github.com/pkg/errors"
)

type TaskDataSql struct {
	DB gdbc.SqlGdbc
}

func (tds *TaskDataSql) Insert(task *model.Task) (*model.Task, error) {

	stmt, err := tds.DB.Prepare("INSERT userinfo SET username=?,department=?,created=?")
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	defer stmt.Close()
	res, err := stmt.Exec(task.Name, task.Id, task.Created)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	task.Id = int(id)
	log.Printf("user inserted: %v", task)
	return task, nil
}

func (tds *TaskDataSql) EnableTx(tx dataservice.TxDataInterface) {
	tds.DB = tx.GetTx()
}
