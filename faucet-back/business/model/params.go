package model

type ParamCheck struct {
	Address string `json:"address" binding:"required"`
}

type ParamClaim struct {
	Address string `json:"address" binding:"required"`
	Taskid  int64  `json:"taskid" binding:"required"`

	Extra string `json:"extra"` //taskid=1时需要此数据
}
