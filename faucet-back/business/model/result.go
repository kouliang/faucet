package model

type ResState struct {
	Hash1 string `json:"hash1"`
	Hash2 string `json:"hash2"`
	Hash3 string `json:"hash3"`
}

type ResInfo struct {
	Code   int64  `json:"code"`
	Taskid int64  `json:"taskid"`
	Msg    string `json:"msg"`
	Extra  string `json:"extra"`
	ResState
}
