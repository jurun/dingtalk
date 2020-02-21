package message

import (
	"fmt"
	"github.com/jurun/dingtalk"
	"github.com/parnurzeal/gorequest"
)
type GetsendprogressRequest struct {
	Agent_id int `json:"agent_id"`
	Task_id int `json:"task_id"`
}

type GetsendprogressResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Progress Progress `json:"progress"`
}

type Progress struct {
	ProgressInPercent int `json:"progress_in_percent"`
	Status            int `json:"status"`
}

type GetsendresultResponse struct {
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	SendResult SendResult `json:"send_result"`
}

type SendResult struct {
	FailedUserIDList    interface{} `json:"failed_user_id_list"`
	ForbiddenUserIDList interface{} `json:"forbidden_user_id_list"`
	InvalidDeptIDList   interface{} `json:"invalid_dept_id_list"`
	InvalidUserIDList   interface{} `json:"invalid_user_id_list"`
	ReadUserIDList      interface{} `json:"read_user_id_list"`
	UnreadUserIDList    interface{} `json:"unread_user_id_list"`
}

//查询工作通知消息的发送结果
func Getsendresult(data GetsendprogressRequest) (rs SendResult, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()

	_url := fmt.Sprintf("%s/topapi/message/corpconversation/getsendresult?access_token=%s",
		dingtalk.ACCESS_URL, accessToken)

	var rsp GetsendresultResponse
	httpResp, _, errs := gorequest.New().Post(_url).Send(data).EndStruct(&rsp)
	if len(errs) > 0 {
		return rs, errs[0]
	}
	fmt.Println(httpResp)

	if httpResp.StatusCode != 200 {
		return rs, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return rs, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return rsp.SendResult, nil
}



//查询工作通知消息的发送进度
func Getsendprogress(data GetsendprogressRequest) (rs Progress, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()

	_url := fmt.Sprintf("%s/topapi/message/corpconversation/getsendprogress?access_token=%s",
		dingtalk.ACCESS_URL, accessToken)

	var rsp GetsendprogressResponse
	httpResp, _, errs := gorequest.New().Post(_url).Send(data).EndStruct(&rsp)
	if len(errs) > 0 {
		return rs, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return rs, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return rs, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return rsp.Progress, nil
}
