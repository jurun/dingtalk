package checkin

import (
	"fmt"
	"github.com/jurun/dingtalk"
	"github.com/parnurzeal/gorequest"
)

type RecordRequest struct {
	Department_id string `json:"department_id"`
	Start_time int `json:"start_time"`
	End_time int `json:"end_time"`
	Offset int `json:"offset"`
	Size int `json:"size"`
	Order string `json:"order"`
}

type RecordResponse struct {
	Data []Data `json:"data"`
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type Data struct {
	Avatar      string   `json:"avatar"`
	DetailPlace string   `json:"detailPlace"`
	ImageList   []string `json:"imageList"`
	Latitude    float64  `json:"latitude"`
	Longitude   float64  `json:"longitude"`
	Name        string   `json:"name"`
	Place       string   `json:"place"`
	Remark      string   `json:"remark"`
	Timestamp   int64    `json:"timestamp"`
	UserID      string   `json:"userId"`
}

type GetRecordResponse struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Result  Result `json:"result"`
}

type Result struct {
	NextCursor int64 `json:"next_cursor"`
	PageList   []struct {
		CheckinTime int64  `json:"checkin_time"`
		DetailPlace string `json:"detail_place"`
		ImageList   []string `json:"image_list"`
		Latitude    string `json:"latitude"`
		Longitude   string `json:"longitude"`
		Place       string `json:"place"`
		Remark      string `json:"remark"`
		Userid      string `json:"userid"`
		VisitUser   string `json:"visit_user"`
	} `json:"page_list"`
}

type GetRecordRequest struct {
	Userid_list string `json:"userid_list"`
	Start_time int `json:"start_time"`
	End_time int `json:"end_time"`
	Cursor int `json:"cursor"`
	Size int `json:"size"`
}

//获取部门用户签到记录
func Record(data RecordRequest) (rs []Data, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()
	fmt.Println(accessToken)

	var _url=""
	if data.Size!=0 {
		_url = fmt.Sprintf("%s/checkin/record?access_token=%s&department_id=%s&start_time=%d&end_time=%d&offset=%d&size=%d&order=%s",
			dingtalk.ACCESS_URL, accessToken,data.Department_id,data.Start_time,data.End_time,data.Offset,data.Size,data.Order)
	}else {
		_url = fmt.Sprintf("%s/checkin/record?access_token=%s&department_id=%s&start_time=%d&end_time=%d&order=%s",
			dingtalk.ACCESS_URL, accessToken,data.Department_id,data.Start_time,data.End_time,data.Order)
	}

	var rsp RecordResponse
	httpResp, _, errs := gorequest.New().Get(_url).EndStruct(&rsp)
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return nil, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return nil, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return rsp.Data, nil
}


//获取用户签到记录
func GetRecord(data GetRecordRequest) (rs Result, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()
	fmt.Println(accessToken)

	_url := fmt.Sprintf("%s/topapi/checkin/record/get?access_token=%s",
			dingtalk.ACCESS_URL, accessToken,)

	var rsp GetRecordResponse
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

	return rsp.Result, nil
}