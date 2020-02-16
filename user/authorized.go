package user

import (
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/jurun/dingtalk"
)

// AuthorizedResponse 企业授权响应数据
type AuthorizedResponse struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	Userid   string `json:"userid"`    // 员工在当前企业内的唯一标识，也称staffId
	SysLevel int    `json:"sys_level"` // 级别，1：主管理员，2：子管理员，100：老板，0：其他（如普通员工）
	IsSys    bool   `json:"is_sys"`    // 是否是管理员，true：是，false：不是
}

// Authorized 用户免登陆授权
func Authorized(code string) (userid string, isSys bool, sysLevel int, err error) {
	accessToken, err := dingtalk.AccessToken.GetToken()
	if err != nil {
		return "", false, 0, err
	}

	_url := fmt.Sprintf("https://oapi.dingtalk.com/user/getuserinfo?access_token=%s&code=%s", accessToken, code)

	var rsp AuthorizedResponse
	_rsp, _, errs := gorequest.New().Get(_url).EndStruct(&rsp)
	if _rsp.StatusCode != 200 {
		return "", false, 0, fmt.Errorf("钉钉服务器异常, httpStatus: %d", _rsp.StatusCode)
	}
	if len(errs) > 0 {
		return "", false, 0, errs[0]
	}
	if rsp.Errcode != 0 {
		return "", false, 0, errors.New(rsp.Errmsg)
	}

	return rsp.Userid, rsp.IsSys, rsp.SysLevel, nil
}
