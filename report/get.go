package report

import (
	"fmt"
	"github.com/jurun/dingtalk"
	"github.com/parnurzeal/gorequest"
)

type CommentListRequest struct {
	Report_id string `json:"report_id"`
	Offset int `json:"offset"`
	Size int `json:"size"`
}

type CommentListResponse struct {
	Errcode int64 `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Result  Cresult `json:"result"`
}

type Cresult struct {
	Comments []struct {
		Content    string `json:"content"`
		CreateTime string `json:"create_time"`
		Userid     string `json:"userid"`
	} `json:"comments"`
	HasMore    bool  `json:"has_more"`
	NextCursor int64 `json:"next_cursor"`
}


type GetunreadcountResponse struct {
	Count int `json:"count"`
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type Listrequest struct {
	Userid string `json:"userid"`
	Start_time int `json:"start_time"`
	End_time int `json:"end_time"`
	Cursor int `json:"cursor"`
	Size int `json:"size"`
}

type ListResponse struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Result  Result `json:"result"`
}

type Result struct {
	DataList []struct {
		Contents []struct {
			Key   string `json:"key"`
			Sort  string `json:"sort"`
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"contents"`
		CreateTime   int64  `json:"create_time"`
		CreatorID    string `json:"creator_id"`
		CreatorName  string `json:"creator_name"`
		DeptName     string `json:"dept_name"`
		Remark       string `json:"remark"`
		ReportID     string `json:"report_id"`
		TemplateName string `json:"template_name"`
	} `json:"data_list"`
	HasMore    bool  `json:"has_more"`
	NextCursor int64 `json:"next_cursor"`
	Size       int64 `json:"size"`
}


type StatisticsRespones struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Result  Sresult `json:"result"`
	Success bool `json:"success"`
}

type Sresult struct {
	CommentNum     int64 `json:"comment_num"`
	CommentUserNum int64 `json:"comment_user_num"`
	LikeNum        int64 `json:"like_num"`
	ReadNum        int64 `json:"read_num"`
}

//获取日志评论详情
func CommentList(data CommentListRequest) (rs Cresult, err error) {
	accessToken, err := dingtalk.AccessToken.GetToken()

	_url := fmt.Sprintf("%s/topapi/report/comment/list?access_token=%s",
		dingtalk.ACCESS_URL, accessToken,)

	var rsp CommentListResponse
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



//获取用户日志未读数
func List(data Listrequest) (rs Result, err error) {
	accessToken, err := dingtalk.AccessToken.GetToken()

	_url := fmt.Sprintf("%s/topapi/report/list?access_token=%s",
		dingtalk.ACCESS_URL, accessToken,)

	var rsp ListResponse
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


//获取用户日志未读数
func Getunreadcount(uid string) (count int, err error) {
	data:=make(map[string]string)
	data["userid"]=uid

	accessToken, err := dingtalk.AccessToken.GetToken()
	fmt.Println(accessToken)

	_url := fmt.Sprintf("%s/topapi/report/getunreadcount?access_token=%s",
		dingtalk.ACCESS_URL, accessToken,)

	var rsp GetunreadcountResponse
	httpResp, _, errs := gorequest.New().Post(_url).Send(data).EndStruct(&rsp)
	if len(errs) > 0 {
		return count, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return count, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return count, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return rsp.Count, nil
}

//获取日志统计数据
func Statistics(report_id string) (rs Sresult, err error) {
	data:=make(map[string]string)
	data["report_id"]=report_id

	accessToken, err := dingtalk.AccessToken.GetToken()
	fmt.Println(accessToken)

	_url := fmt.Sprintf("%s/topapi/report/statistics?access_token=%s",
		dingtalk.ACCESS_URL, accessToken,)

	var rsp StatisticsRespones
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