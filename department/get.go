package department

import (
	"fmt"
	"github.com/jurun/dingtalk"
	"github.com/parnurzeal/gorequest"
)

type GetResponse struct {
	AutoAddUser           bool   `json:"autoAddUser"`
	CreateDeptGroup       bool   `json:"createDeptGroup"`
	DeptHiding            bool   `json:"deptHiding"`
	DeptManagerUseridList string `json:"deptManagerUseridList"`
	DeptPermits           string `json:"deptPermits"`
	Errcode               int    `json:"errcode"`
	Errmsg                string `json:"errmsg"`
	ID                    int    `json:"id"`
	Name                  string `json:"name"`
	Order                 int    `json:"order"`
	OrgDeptOwner          string `json:"orgDeptOwner"`
	OuterDept             bool   `json:"outerDept"`
	OuterPermitDepts      string `json:"outerPermitDepts"`
	OuterPermitUsers      string `json:"outerPermitUsers"`
	Parentid              int    `json:"parentid"`
	SourceIdentifier      string `json:"sourceIdentifier"`
	UserPermits           string `json:"userPermits"`
}


type ListResponse struct {
	Department []*ListDepartment `json:"department"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type ListDepartment struct {
	AutoAddUser     bool   `json:"autoAddUser"`
	CreateDeptGroup bool   `json:"createDeptGroup"`
	Ext             string `json:"ext"`
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Parentid        int    `json:"parentid"`
	Child []*ListDepartment `json:"child"`
}

type List_idsResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	sub_dept_id_list []int `json:"sub_dept_id_list"`
}

type Parent_idsResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	ParentIds []int `json:"parentIds"`
}


//获取部门详情
func Get(depid string) (rsp GetResponse, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()

		_url := fmt.Sprintf("%s/department/get?access_token=%s&id=%s",
			dingtalk.ACCESS_URL, accessToken,depid)

	httpResp, _, errs := gorequest.New().Get(_url).EndStruct(&rsp)
	if len(errs) > 0 {
		return GetResponse{}, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return GetResponse{}, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return GetResponse{}, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return rsp, nil
}


//获取部门列表
func List(depid string) (department []*ListDepartment, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()
	fmt.Println(accessToken)

	var _url=""
	if depid=="" {
		_url = fmt.Sprintf("%s/department/list?access_token=%s",
			dingtalk.ACCESS_URL, accessToken)
	}else {
		_url = fmt.Sprintf("%s/department/list?access_token=%s&id=%s",
			dingtalk.ACCESS_URL, accessToken,depid)
	}

	var rsp ListResponse
	httpResp, _, errs := gorequest.New().Get(_url).EndStruct(&rsp)
	if len(errs) > 0 {
		return department, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return department, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return department, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return rsp.Department, nil
}


//获取子部门ID列表
func List_ids(depid string) (sub_dept_id_list []int, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()
	_url := fmt.Sprintf("%s/department/list_ids?access_token=%s&id=%s",
		dingtalk.ACCESS_URL, accessToken,depid)

	var rsp List_idsResponse
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

	return rsp.sub_dept_id_list, nil
}

//获取父部门路径
func List_parent_depts_by_dept(depid string) (parentids []int, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()
	_url := fmt.Sprintf("%s/department/list_parent_depts_by_dept?access_token=%s&id=%s",
		dingtalk.ACCESS_URL, accessToken,depid)

	var rsp Parent_idsResponse
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

	return rsp.ParentIds, nil
}
