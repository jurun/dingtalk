package callback

import (
	"fmt"
	"github.com/jurun/dingtalk"
	"github.com/parnurzeal/gorequest"
)

type RegistercallbackRequest struct {
	AesKey      string `json:"aes_key"`
	CallBackTag []string `json:"call_back_tag"`
	Token       string `json:"token"`
	URL         string `json:"url"`
}

type RegistercallbackResponse struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}


type Get_call_back_failed_resultResponse struct {
	Errcode    int64  `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	FailedList []struct {
		BpmsInstanceChange struct {
			BpmsCallBackData struct{} `json:"bpmsCallBackData"`
			Corpid           string   `json:"corpid"`
		} `json:"bpms_instance_change"`
		BpmsTaskChange struct {
			BpmsCallBackData struct{} `json:"bpmsCallBackData"`
			Corpid           string   `json:"corpid"`
		} `json:"bpms_task_change"`
		LabelConfAdd struct {
			Corpid          string   `json:"corpid"`
			RoleLabelChange struct{} `json:"roleLabelChange"`
		} `json:"label_conf_add"`
		UserAddOrg struct {
			CallbackData struct{} `json:"callbackData"`
			Corpid       string   `json:"corpid"`
		} `json:"user_add_org"`
	} `json:"failed_list"`
	HasMore bool `json:"has_more"`
}


//注册业务事件回调接口
func Register_call_back(data RegistercallbackRequest) (rs bool, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()
	fmt.Println(accessToken)

	_url := fmt.Sprintf("%s/call_back/register_call_back?access_token=%s",
		dingtalk.ACCESS_URL, accessToken)

	var rsp RegistercallbackResponse
	httpResp, _, errs := gorequest.New().Post(_url).Send(data).EndStruct(&rsp)
	if len(errs) > 0 {
		return false, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return false, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return false, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return true, nil
}

//更新业务事件回调接口
func Update_call_back(data RegistercallbackRequest) (rs bool, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()
	fmt.Println(accessToken)

	_url := fmt.Sprintf("%s/call_back/update_call_back?access_token=%s",
		dingtalk.ACCESS_URL, accessToken)

	var rsp RegistercallbackResponse
	httpResp, _, errs := gorequest.New().Post(_url).Send(data).EndStruct(&rsp)
	if len(errs) > 0 {
		return false, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return false, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return false, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return true, nil
}

//删除业务事件回调接口
func Delete_call_back() (rs bool, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()
	fmt.Println(accessToken)

	_url := fmt.Sprintf("%s/call_back/delete_call_back?access_token=%s",
		dingtalk.ACCESS_URL, accessToken)

	var rsp RegistercallbackResponse
	httpResp, _, errs := gorequest.New().Get(_url).EndStruct(&rsp)
	if len(errs) > 0 {
		return false, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return false, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return false, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return true, nil
}

//查询事件回调接口
func Get_call_back() (rs bool, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()
	fmt.Println(accessToken)

	_url := fmt.Sprintf("%s/call_back/get_call_back?access_token=%s",
		dingtalk.ACCESS_URL, accessToken)

	var rsp RegistercallbackResponse
	httpResp, _, errs := gorequest.New().Get(_url).EndStruct(&rsp)
	if len(errs) > 0 {
		return false, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return false, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return false, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return true, nil
}

//获取回调失败的结果
func Get_call_back_failed_result() (rs Get_call_back_failed_resultResponse, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()
	fmt.Println(accessToken)

	_url := fmt.Sprintf("%s/call_back/get_call_back_failed_result?access_token=%s",
		dingtalk.ACCESS_URL, accessToken)

	var rsp Get_call_back_failed_resultResponse
	httpResp, _, errs := gorequest.New().Get(_url).EndStruct(&rsp)
	if len(errs) > 0 {
		return rs, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return rs, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return rs, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return rs, nil
}