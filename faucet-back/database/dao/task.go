package dao

import (
	"faucet-app/business/model"
	"faucet-app/database"
	"log"
	"strings"
	"time"
)

func TaskCheckLatest(address string, taskid int64) *model.Task {
	address = strings.ToLower(address)

	task := new(model.Task)
	sqlStr := `SELECT * FROM task WHERE address=? AND taskid=?
	ORDER BY timestamp DESC
	LIMIT 1;`

	database.MyDB.Get(task, sqlStr, address, taskid)

	return task
}

func TaskCheck(address string, taskid int64) (tasks []model.Task) {
	address = strings.ToLower(address)

	sqlStr := `SELECT * FROM task WHERE address=? AND taskid=?
	ORDER BY timestamp DESC;`

	database.MyDB.Select(&tasks, sqlStr, address, taskid)
	return tasks
}

func TaskInsert(task model.Task) error {
	task.Address = strings.ToLower(task.Address)
	task.Timestamp = time.Now().Unix()

	sqlStr := `INSERT INTO task(address, timestamp, taskid, hash) values(?,?,?,?)`

	_, err := database.MyDB.Exec(sqlStr, task.Address, task.Timestamp, task.Taskid, task.Hash)
	if err != nil {
		log.Println("database insert error ===============")
		log.Println(err)
	}
	return nil
}
