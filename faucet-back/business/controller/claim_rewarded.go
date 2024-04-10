package controller

import (
	"log"

	"github.com/gin-gonic/gin"

	"faucet-app/business/model"
	"faucet-app/database/dao"

	"faucet-app/business/verify"
	"faucet-app/web3"
	"faucet-app/web3/account"
)

func ClaimHandler(ctx *gin.Context) {

	var result model.ResInfo

	//1.参数解析
	p := new(model.ParamClaim)
	ctx.ShouldBindJSON(p)

	//2.验证参数有效性
	if p.Taskid < 1 || p.Taskid > 3 || !account.IsAvailableAddress(p.Address) {
		result.Code = 101
		result.Msg = "parameter error"
		ctx.JSON(200, result)
		return
	}
	result.Taskid = p.Taskid

	//3.验证当天有没有重复领取
	todayHash := checkTodayHash(p.Address, p.Taskid)
	if todayHash != "" {
		result.Code = 102
		result.Msg = "you've already claimed it today"
		result.ResState = checkTodayState(p.Address)

		ctx.JSON(200, result)
		return
	}

	//4.验证领取资格
	complete, err := verify.Verify(*p)
	if !complete {
		result.Code = 102
		result.Msg = "incomplete task"
		result.ResState = checkTodayState(p.Address)
		if err != nil {
			result.Extra = err.Error()
		}

		ctx.JSON(200, result)
		return
	}

	//5.领取奖励
	log.Println("SendErc20Token===" + p.Address)
	txHash, _ := web3.SendErc20Token(p.Address, "1000000000000000000000")
	if txHash == nil {
		result.Code = 103
		result.Msg = "transaction error"
		result.ResState = checkTodayState(p.Address)

		ctx.JSON(200, result)
		return
	}
	log.Println("Erc20Hash===", txHash)
	//5.1插入数据库
	task := model.Task{
		Address: p.Address,
		Taskid:  p.Taskid,
		Hash:    *txHash,
	}
	if err := dao.TaskInsert(task); err != nil {
		result.Code = 104
		result.Msg = "insert database error"
		result.ResState = checkTodayState(p.Address)

		ctx.JSON(200, result)
		return
	}
	log.Println("Erc20Hash insert database")

	//6.如果没有发放过燃料币，赠送一次
	task4 := dao.TaskCheckLatest(p.Address, 4)
	if task4.Hash == "" {
		ethHash, _ := web3.SendETH(p.Address, "1000000000000000000")
		if ethHash != nil {
			newTask := model.Task{
				Address: p.Address,
				Taskid:  4,
				Hash:    *txHash,
			}
			dao.TaskInsert(newTask)
		}
	}

	result.Code = 100
	result.Msg = "success"
	result.ResState = checkTodayState(p.Address)
	ctx.JSON(200, result)
}
