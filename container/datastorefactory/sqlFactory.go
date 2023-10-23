package datastorefactory

import (
	"database/sql"
	"log"

	config "github.com/0x-chaitu/gocrudMicro/appconfig"
	"github.com/0x-chaitu/gocrudMicro/container"
	"github.com/0x-chaitu/gocrudMicro/tool/gdbc/databasehandler"
	"github.com/pkg/errors"
)

type sqlFactory struct{}

func (sf *sqlFactory) Build(c container.Container, dsc *config.DataStoreConfig) (DataStoreInterface, error) {
	log.Println("sqlfactory")
	key := dsc.Code

	if value, found := c.Get(key); found {
		//type assertion: convert interface to concrete type
		sdb := value.(*sql.DB)
		sdt := databasehandler.SqlDbTx{DB: sdb}
		log.Printf("found db in container %v:\n", sdt)
		return &sdt, nil
	}

	db, err := sql.Open(dsc.DriverName, dsc.UrlAdress)
	if err != nil {
		return nil, errors.Wrap(err, "found err in opening db")
	}
	dt := databasehandler.SqlDbTx{DB: db}
	c.Put(key, db)
	return &dt, nil
}
