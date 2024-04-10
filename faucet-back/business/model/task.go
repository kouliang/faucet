package model

type Task struct {
	Address   string `db:"address" json:"address"`
	Timestamp int64  `db:"timestamp" json:"timestamp"`
	Taskid    int64  `db:"taskid" json:"taskid"`
	Hash      string `db:"hash" json:"hash"`
}
