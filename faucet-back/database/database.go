package database

import (
	"faucet-app/setting"

	_ "github.com/glebarez/sqlite"
	"github.com/jmoiron/sqlx"
)

var MyDB *sqlx.DB

func Init(cfg *setting.DatabaseConfig) (err error) {
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	cfg.User,
	// 	cfg.Password,
	// 	cfg.Host,
	// 	cfg.Port,
	// 	cfg.DbName,
	// )
	dsn := cfg.DBPath

	// MyDB, err = sqlx.Connect("mysql", dsn)
	// MyDB.SetMaxOpenConns(cfg.MaxOpenConns)
	// MyDB.SetMaxIdleConns(cfg.MaxIdleConns)

	MyDB, err = sqlx.Connect("sqlite", dsn)
	if err != nil {
		return
	}
	err = initTable()
	if err != nil {
		return
	}

	return
}

func initTable() error {
	sqlc := `CREATE TABLE IF NOT EXISTS task (
		address VARCHAR(255),
		timestamp INTEGER,
		taskid INTEGER,
		hash VARCHAR(255),
		PRIMARY KEY (address, timestamp, taskid)
	  );`

	_, err := MyDB.Exec(sqlc)
	return err
}

func Close() {
	MyDB.Close()
}
