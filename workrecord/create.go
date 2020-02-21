package workrecord

import (
	"fmt"
	"github.com/jurun/dingtalk"
	"github.com/parnurzeal/gorequest"
)

type UpdateRequest struct {
	Userid string `json:"userid"`
	Record_id string `json:"record_id"`
}

type AddRequest struct {
	Userid string `json:"userid"`
	Create_time int `json:"create_time"`
	Title string `json:"title"`
	Url string `json:"url"`
	FormItemList FormItemList `json:"formItemList"`
	PcUrl string `json:"pcUrl"`
	Source_name string `json:"source_name"`
	Biz_id string `json:"biz_id"`
	//PcOpenType int `json:"pc_open_type"`//这个加上 如果不填会报错  不知道为什么
}

type FormItemList struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

type AddResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Record_id string    `json:"record_id"`
}


//发起待办
func Add(data AddRequest) (record_id string, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()
	fmt.Println(accessToken)

	_url := fmt.Sprintf("%s/topapi/workrecord/add?access_token=%s",
		dingtalk.ACCESS_URL, accessToken)

	var rsp AddResponse
	httpResp, _, errs := gorequest.New().Post(_url).Send(data).EndStruct(&rsp)
	if len(errs) > 0 {
		return record_id, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return record_id, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return record_id, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return rsp.Record_id, nil
}


//更新待办
func Update(data UpdateRequest) (rs bool, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()
	fmt.Println(accessToken)

	_url := fmt.Sprintf("%s/topapi/workrecord/update?access_token=%s",
		dingtalk.ACCESS_URL, accessToken)

	var rsp AddResponse
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
