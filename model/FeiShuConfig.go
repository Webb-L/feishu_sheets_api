package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type FeiShuConfig struct {
	SpreadsheetToken string
	AppId            string
	AppSecret        string
	Token            string
}

type TenantAccessToken struct {
	Code   int    `json:"code"`
	Expire int    `json:"expire"`
	Msg    string `json:"msg"`
	Token  string `json:"tenant_access_token"`
}

// GetTenantAccessToken @title GetTenantAccessToken
// @description 获取TenantAccessToken
func (feiShuConfig *FeiShuConfig) GetTenantAccessToken() {
	payload := strings.NewReader(fmt.Sprintf(`{
	"app_id": "%s",
	"app_secret": "%s"
}`, feiShuConfig.AppId, feiShuConfig.AppSecret))
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal", payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	tenantAccessToken := &TenantAccessToken{}

	json.Unmarshal(body, tenantAccessToken)
	feiShuConfig.Token = tenantAccessToken.Token
}

// CheckIsCreate @title CheckIsCreate
// @description 检测对象是否创建
// @return bool
func (feiShuConfig *FeiShuConfig) CheckIsCreate() bool {
	return feiShuConfig.AppId != "" && feiShuConfig.AppSecret != ""
}
