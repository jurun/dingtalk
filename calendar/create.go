package calendar

import (
	"fmt"
	"github.com/jurun/dingtalk"
	"github.com/parnurzeal/gorequest"
)

//todo 可选
type AddRequest struct {
	CreateVo CreateVo `json:"create_vo"`
}

type CreateVo struct {
	BizID         string `json:"biz_id"`
	CalendarType  string `json:"calendar_type"`
	CreatorUserid string `json:"creator_userid"`
	EndTime EndTime `json:"end_time"`
	ReceiverUserids string `json:"receiver_userids"`
	Source  Source `json:"source"`
	StartTime StartTime `json:"start_time"`
	Summary string `json:"summary"`
	UUID    string `json:"uuid"`
}

type StartTime struct {
	UnixTimestamp string `json:"unix_timestamp"`
}

type EndTime struct {
	UnixTimestamp string `json:"unix_timestamp"`
}

type Source struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}


type AddResponse struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Result  Result `json:"result"`
}

type Result struct {
	DingtalkCalendarID string `json:"dingtalk_calendar_id"`
	InvalidUserids []string `json:"invalid_userids"`
}


//
//type InvalidUserids struct {
//	String []string `json:"string"`
//}


//增加日程
func Add(data AddRequest) (rs Result, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()
	fmt.Println(accessToken)

	_url := fmt.Sprintf("%s/topapi/calendar/create?access_token=%s",
		dingtalk.ACCESS_URL, accessToken)

	var rsp AddResponse
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

