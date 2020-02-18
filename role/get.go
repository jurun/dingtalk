package role

import (
	"fmt"
	"github.com/jurun/dingtalk"
	"github.com/parnurzeal/gorequest"
)

type ListRequest struct {
	offset int `json:"offset"`
	size int `json:"size"`
}

type SimplelistRequest struct {
	roleId int `json:"role_id"`
	offset int `json:"offset"`
	size int `json:"size"`
}

type GetrolegroupResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Role_Group  Role_group `json:"role_group"`
}

type Role_group struct {
	GroupName string `json:"group_name"`
	Roles     []struct {
		RoleID   int    `json:"role_id"`
		RoleName string `json:"role_name"`
	} `json:"roles"`
}


type GetroleResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Role    Role `json:"role"`
}

type Role struct {
	GroupID int    `json:"groupId"`
	Name    string `json:"name"`
}


type ListResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Result  ListResult `json:"result"`
}

type ListResult struct {
		HasMore bool `json:"hasMore"`
		List    []struct {
			GroupID int    `json:"groupId"`
			Name    string `json:"name"`
			Roles   []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"roles"`
		} `json:"list"`
}


type SimplelistResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Result  SimpleResult `json:"result"`
}

type SimpleResult struct {
	HasMore bool `json:"hasMore"`
	List    []struct {
		Name   string `json:"name"`
		Userid string `json:"userid"`
	} `json:"list"`
}

//获取角色列表
func List(data ListRequest) (userlist ListResult, err error) {
	accessToken, err := dingtalk.AccessToken.GetToken()
	var _url=""
	if data.offset!=0&&data.size!=0 {
		_url = fmt.Sprintf("%s/topapi/role/list?access_token=%s&offset=%d&size=%d",
			dingtalk.ACCESS_URL, accessToken,data.offset,data.size)
	}else {
		_url = fmt.Sprintf("%s/topapi/role/list?access_token=%s",
			dingtalk.ACCESS_URL, accessToken)
	}
	var rsp ListResponse
	httpResp, _, errs := gorequest.New().Get(_url).EndStruct(&rsp)
	if len(errs) > 0 {
		return userlist, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return userlist, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return userlist, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return rsp.Result, nil
}


//获取角色下的员工列表
func Simplelist(data SimplelistRequest) (userlist SimpleResult, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()
	fmt.Println(accessToken)

	var _url=""
	if data.offset!=0&&data.size!=0 {
		_url = fmt.Sprintf("%s/topapi/role/simplelist?access_token=%s&role_id=%d&offset=%d&size=%d",
			dingtalk.ACCESS_URL, accessToken,data.roleId,data.offset,data.size)
	}else {
		_url = fmt.Sprintf("%s/topapi/role/simplelist?access_token=%s&role_id=%d",
			dingtalk.ACCESS_URL, accessToken,data.roleId)
	}


	var rsp SimplelistResponse
	httpResp, _, errs := gorequest.New().Get(_url).EndStruct(&rsp)
	if len(errs) > 0 {
		return userlist, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return userlist, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return userlist, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return rsp.Result, nil
}


//获取角色组
func Getrolegroup(groupid int) (role_group Role_group, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()
	fmt.Println(accessToken)

	_url := fmt.Sprintf("%s/topapi/role/getrolegroup?access_token=%s&group_id=%d",dingtalk.ACCESS_URL, accessToken,groupid)

	var rsp  GetrolegroupResponse
	httpResp, _, errs := gorequest.New().Get(_url).EndStruct(&rsp)
	if len(errs) > 0 {
		return role_group, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return role_group, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return role_group, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return rsp.Role_Group, nil
}


//获取角色详情
func Getrole(roleid int) (role Role, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()

	_url := fmt.Sprintf("%s/topapi/role/getrole?access_token=%s&roleId=%d",dingtalk.ACCESS_URL, accessToken,roleid)

	var rsp  GetroleResponse
	httpResp, _, errs := gorequest.New().Get(_url).EndStruct(&rsp)
	if len(errs) > 0 {
		return role, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return role, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return role, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return rsp.Role, nil
}
