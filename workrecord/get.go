package workrecord

import (
	"fmt"
	"github.com/jurun/dingtalk"
	"github.com/parnurzeal/gorequest"
)

type SimplelistRequest struct {
	Userid string `json:"userid"`
	Offset int `json:"offset"`
	Limit int `json:"limit"`
	Status int `json:"status"`
}

type WorkrecordlistResponse struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Records Records `json:"records"`
}

type Records struct {
	HasMore bool `json:"has_more"`
	List    []struct {
		CreateTime int64 `json:"create_time"`
		Forms      []struct {
			Content string `json:"content"`
			Title   string `json:"title"`
		} `json:"forms"`
		RecordID string `json:"record_id"`
		Title    string `json:"title"`
		URL      string `json:"url"`
	} `json:"list"`
}


//获取部门用户
func Workrecordlist(data SimplelistRequest) (Records Records, err error) {
	accessToken, err := dingtalk.AccessToken.GetToken()
	//var _url=""
	//if data.Offset!=0&&data.Limit!=0 {
	//	_url = fmt.Sprintf("%s/topapi/workrecord/getbyuserid?access_token=%s",
	//		dingtalk.ACCESS_URL, accessToken )
	//}else {
	_url := fmt.Sprintf("%s/topapi/workrecord/getbyuserid?access_token=%s",
			dingtalk.ACCESS_URL, accessToken )
	//}
	var rsp WorkrecordlistResponse
	httpResp, _, errs := gorequest.New().Post(_url).Send(data).EndStruct(&rsp)
	if len(errs) > 0 {
		return Records, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return Records, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return Records, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return rsp.Records, nil
}