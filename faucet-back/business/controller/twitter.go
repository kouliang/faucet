package controller

import (
	"encoding/json"
	"faucet-app/business/model"
	"faucet-app/business/verify"
	"faucet-app/web3/account"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Twitter(ctx *gin.Context) {
	state := ctx.Query("state")
	code := ctx.Query("code")

	newPath := fmt.Sprintf("/faucet?source=twitter&state=%s&code=%s", state, code)

	ctx.Redirect(http.StatusMovedPermanently, newPath)
}

func TwitterUploadCode(ctx *gin.Context) {
	var result model.ResInfo

	//1.参数解析
	p := new(model.ParamClaim)
	ctx.ShouldBindJSON(p)

	//2.验证参数有效性
	if p.Taskid != 1 || !account.IsAvailableAddress(p.Address) {
		result.Code = 101
		result.Msg = "parameter error"
		ctx.JSON(200, result)
		return
	}
	result.Taskid = p.Taskid

	requestParam := fmt.Sprintf("address=%s&code=%s&codeNo=1&merchantOn=MPNY2TJBNFPNDM&state=%s", p.Address, p.Extra, p.Address)
	request := verify.VerifyRequest("/faucet/twitterCodeToToken?", requestParam)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		result.Code = 102
		result.Msg = "verify service error"
		result.Extra = err.Error()
		result.ResState = checkTodayState(p.Address)
		ctx.JSON(200, result)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var response ShareUrlResponse
	err = json.Unmarshal(body, &response)
	if response.Code != 200 {
		result.Code = 103
		result.Msg = "verify data error"
		result.ResState = checkTodayState(p.Address)
		if err != nil {
			result.Extra = err.Error()
		} else {
			result.Extra = response.Msg
		}

		ctx.JSON(200, result)
		return
	}

	result.Code = 100
	result.Msg = "success"
	result.ResState = checkTodayState(p.Address)
	if response.Data.CodeUrl != "" {
		result.Extra = response.Data.CodeUrl
	} else if response.Data.URL != "" {
		result.Extra = response.Data.URL
	}
	ctx.JSON(200, result)
}
