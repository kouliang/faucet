package verify

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"faucet-app/business/model"
	"faucet-app/setting"
	"fmt"
	"io"
	"net/http"
)

type VerifyResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Status        int `json:"status"`
		TwitterNumber int `json:"twitterNumber"`
	} `json:"data"`
}

func Verify(task model.ParamClaim) (bool, error) {

	requestParam := fmt.Sprintf("address=%s&code=%s&codeNo=1&day=1&merchantOn=MPNY2TJBNFPNDM&status=%d", task.Address, task.Extra, task.Taskid)
	request := VerifyRequest("/faucet/findChannelMember?", requestParam)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var response VerifyResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return false, err
	}

	if task.Taskid == 1 {
		return response.Data.TwitterNumber > 0, nil
	} else {
		return response.Data.Status == 1, nil
	}
}

func VerifyRequest(path, param string) *http.Request {

	fulUrl := setting.Conf.VerifyHost + path + param

	req, err := http.NewRequest("GET", fulUrl, nil)
	if err != nil {
		panic(err)
	}

	// 添加请求头信息
	sign := sign(param, "")
	req.Header.Set("sign", sign)

	return req
}

func sign(param string, privateStr string) string {

	block, _ := pem.Decode([]byte(privateStr))
	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	hash := sha256.Sum256([]byte(param))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(signature)
}
