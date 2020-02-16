package department

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/jurun/dingtalk"
)

type CreateRequest struct {
	Parentid         string `json:"parentid"`
	Order            string `json:"order"`
	Createdeptgroup  bool   `json:"createDeptGroup"`
	Sourceidentifier string `json:"sourceIdentifier"`
	Depthiding       bool   `json:"deptHiding"`
	Name             string `json:"name"`
	Outerpermitusers string `json:"outerPermitUsers"`
	Userpermits      string `json:"userPermits"`
	Outerpermitdepts string `json:"outerPermitDepts"`
	Deptpermits      string `json:"deptPermits"`
	Outerdept        bool   `json:"outerDept"`
	Ext              string `json:"ext"`
}

type CreateResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	ID      int    `json:"id"`
}

func Create(data CreateRequest) (depId int, err error) {
	// @todo 验证 data

	accessToken, err := dingtalk.AccessToken.GetToken()

	_url := fmt.Sprintf("https://oapi.dingtalk.com/department/create?access_token=%s", accessToken)

	var rsp CreateResponse
	httpResp, _, errs := gorequest.New().Get(_url).EndStruct(&rsp)
	if len(errs) > 0 {
		return 0, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return 0, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return 0, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return rsp.ID, nil
}
