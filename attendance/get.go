package attendance

import (
	"fmt"
	"github.com/jurun/dingtalk"
	"github.com/parnurzeal/gorequest"
)

type ListRecordRequest struct {
	CheckDateFrom string `json:"checkDateFrom"`
	CheckDateTo   string `json:"checkDateTo"`
	UserIds       []int`json:"userIds"`
}

type ListRecordResponse struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	Recordresult []Recordresult `json:"recordresult"`
}

type Recordresult struct {
	BaseCheckTime  int64     `json:"baseCheckTime"`
	CheckType      string  `json:"checkType"`
	CorpID         string  `json:"corpId"`
	DeviceID       string  `json:"deviceId"`
	GroupID        int     `json:"groupId"`
	ID             int     `json:"id"`
	IsLegal        string  `json:"isLegal"`
	LocationMethod string  `json:"locationMethod"`
	LocationResult string  `json:"locationResult"`
	PlanCheckTime  int64     `json:"planCheckTime"`
	PlanID         int     `json:"planId"`
	ProcInstID     string  `json:"procInstId"`
	SourceType     string  `json:"sourceType"`
	TimeResult     string  `json:"timeResult"`
	UserAccuracy   int     `json:"userAccuracy"`
	UserAddress    string  `json:"userAddress"`
	UserCheckTime  int64     `json:"userCheckTime"`
	UserID         string  `json:"userId"`
	UserLatitude   int     `json:"userLatitude"`
	UserLongitude  float64 `json:"userLongitude"`
	WorkDate       int64     `json:"workDate"`
}

type ListRequest struct {
	WorkDateFrom string `json:"workDateFrom"`
	WorkDateTo   string `json:"workDateTo"`
	UserIdList []int`json:"userIdList"`
	Offset int `json:"offset"`
	Limit int `json:"limit"`
}

type ListResponse struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	HasMore      bool   `json:"hasMore"`
	Recordresult []Recordlistresult `json:"recordresult"`
}

type Recordlistresult struct {
	BaseCheckTime  int64    `json:"baseCheckTime"`
	CheckType      string `json:"checkType"`
	CorpID         string `json:"corpId"`
	GroupID        int    `json:"groupId"`
	ID             int    `json:"id"`
	LocationResult string `json:"locationResult"`
	PlanID         int    `json:"planId"`
	ProcInstID     string `json:"procInstId"`
	RecordID       int    `json:"recordId"`
	TimeResult     string `json:"timeResult"`
	UserCheckTime  int64    `json:"userCheckTime"`
	UserID         string `json:"userId"`
	WorkDate       int64    `json:"workDate"`
}

//获取打卡详情
func ListRecord(data ListRecordRequest) (reclist []Recordresult, err error) {
	accessToken, err := dingtalk.AccessToken.GetToken()
	_url := fmt.Sprintf("%s/attendance/listRecord?access_token=%s",
			dingtalk.ACCESS_URL, accessToken)
	var rsp ListRecordResponse
	httpResp, _, errs := gorequest.New().Post(_url).Send(data).EndStruct(&rsp)
	if len(errs) > 0 {
		return reclist, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return reclist, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return reclist, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return reclist, nil
}
//获取打卡详情结果
func List(data ListRequest) (reclist []Recordlistresult, err error) {
	accessToken, err := dingtalk.AccessToken.GetToken()
	_url := fmt.Sprintf("%s/attendance/list?access_token=%s",
		dingtalk.ACCESS_URL, accessToken)
	var rsp ListRecordResponse
	httpResp, _, errs := gorequest.New().Post(_url).Send(data).EndStruct(&rsp)
	if len(errs) > 0 {
		return reclist, errs[0]
	}


	if httpResp.StatusCode != 200 {
		return reclist, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return reclist, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return reclist, nil
}





