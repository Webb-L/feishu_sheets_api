package model

import (
	"feishu_sheets_api/utils"
)

type FeiShuData struct {
	password string
}

// BuildPassWord @title BuildPassWord
// @description 生成密码。
// @return string
func (feiShuData *FeiShuData) BuildPassWord() string {
	feiShuData.password = utils.RandStr(128)
	return feiShuData.password
}

// CheckPassWord @title CheckPassWord
// @description 检查密码。
// @param password string
// @return bool
func (feiShuData *FeiShuData) CheckPassWord(password string) bool {
	return feiShuData.password != "" && feiShuData.password == password
}
