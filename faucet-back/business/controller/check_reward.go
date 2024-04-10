package controller

import (
	"faucet-app/business/model"
	"faucet-app/database/dao"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckHandler(ctx *gin.Context) {
	//1.参数解析
	p := new(model.ParamCheck)
	if ctx.ShouldBindJSON(p) != nil {
		ctx.JSON(200, gin.H{
			"code": "101",
			"msg":  "Parameter error",
		})
		return
	}

	ctx.JSON(200, checkTodayState(p.Address))
}

func CheckAll(ctx *gin.Context) {
	//1.参数解析
	p := new(model.ParamClaim)
	if ctx.ShouldBindJSON(p) != nil {
		ctx.JSON(200, gin.H{
			"code": "101",
			"msg":  "Parameter error",
		})
		return
	}

	tasks := dao.TaskCheck(p.Address, p.Taskid)
	ctx.JSON(200, tasks)
}

func checkTodayState(address string) model.ResState {
	task1 := dao.TaskCheckLatest(address, 1)
	task2 := dao.TaskCheckLatest(address, 2)
	task3 := dao.TaskCheckLatest(address, 3)

	var state model.ResState
	if isSameDay(task1.Timestamp) {
		state.Hash1 = task1.Hash
	}
	if isSameDay(task2.Timestamp) {
		state.Hash2 = task2.Hash
	}
	if isSameDay(task3.Timestamp) {
		state.Hash3 = task3.Hash
	}

	return state
}

func checkTodayHash(address string, taskid int64) string {
	task := dao.TaskCheckLatest(address, taskid)

	if isSameDay(task.Timestamp) {
		return task.Hash
	} else {
		return ""
	}
}

// 判断时间戳是否属于当天
func isSameDay(timestamp int64) bool {
	now := time.Now()
	t := time.Unix(timestamp, 0)

	return t.Year() == now.Year() && t.Month() == now.Month() && t.Day() == now.Day()
}
