package user

import (
	"fmt"
	"github.com/jurun/dingtalk"
	"github.com/parnurzeal/gorequest"
)

type UserinfoResponse struct {
	Active          bool   `json:"active"`
	Avatar          string `json:"avatar"`
	Department      []int  `json:"department"`
	Errcode         int    `json:"errcode"`
	Errmsg          string `json:"errmsg"`
	HiredDate       int64    `json:"hiredDate"`
	IsAdmin         bool   `json:"isAdmin"`
	IsBoss          bool   `json:"isBoss"`
	IsHide          bool   `json:"isHide"`
	IsLeaderInDepts string `json:"isLeaderInDepts"`
	IsSenior        bool   `json:"isSenior"`
	Jobnumber       string `json:"jobnumber"`
	Mobile          string   `json:"mobile"`
	Name            string `json:"name"`
	OrderInDepts    string `json:"orderInDepts"`
	Position        string `json:"position"`
	Roles           []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
	} `json:"roles"`
	Unionid string `json:"unionid"`
	Userid  string `json:"userid"`
}

type ListbypageRequest struct {
	DeptId int `json:"deptId"`
	Offset int `json:"offset"`
	Size int `json:"size"`
}

type ListbypageUserlist struct {
	Active     bool   `json:"active"`
	Avatar     string `json:"avatar"`
	Department []int  `json:"department"`
	Email      string `json:"email"`
	//Extattr    struct {
	//	年龄 string `json:"年龄"`
	//	爱好 string `json:"爱好"`
	//} `json:"extattr"`
	IsAdmin   bool   `json:"isAdmin"`
	IsBoss    bool   `json:"isBoss"`
	IsHide    bool   `json:"isHide"`
	IsLeader  bool   `json:"isLeader"`
	Jobnumber string `json:"jobnumber"`
	Mobile    string `json:"mobile"`
	Name      string `json:"name"`
	Order     int    `json:"order"`
	Position  string `json:"position"`
	Remark    string `json:"remark"`
	Tel       string `json:"tel"`
	Unionid   string `json:"unionid"`
	Userid    string `json:"userid"`
	WorkPlace string `json:"workPlace"`
}

type ListbypageResponse struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	HasMore  bool   `json:"hasMore"`
	Userlist []ListbypageUserlist `json:"userlist"`
}


type SimplelistRequest struct {
	deptId int `json:"deptId"`
	offset int `json:"offset"`
	size int `json:"size"`
}


type SimplelistResponse struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	HasMore  bool   `json:"hasMore"`
	Userlist []SimplelistUserlist `json:"userlist"`
}


type SimplelistUserlist struct {
	Name   string `json:"name"`
	Userid string `json:"userid"`
}

type Get_adminResponse struct {
	AdminList []Adminlist `json:"admin_list"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type Adminlist struct {
	SysLevel int    `json:"sys_level"`
	Userid   string `json:"userid"`
}

type GetDeptMemberResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Userids []string `json:"userIds"`
}


type GetUseridByUnionidResponse struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	ContactType int    `json:"contactType"`//联系类型，0表示企业内部员工，1表示企业外部联系人
	Userid string `json:"userid"`
}


type GetbyMobileResponse struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	Userid string `json:"userid"`
}

//根据unionid获取userid
func GetUseridByUnionid(unionid string) (ContactType int,Userid string, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()
	fmt.Println(accessToken)
    _url := fmt.Sprintf("%s/user/getUseridByUnionid?access_token=%s&unionid=%s",
			dingtalk.ACCESS_URL, accessToken,unionid)

	var rsp GetUseridByUnionidResponse
	httpResp, _, errs := gorequest.New().Get(_url).EndStruct(&rsp)
	if len(errs) > 0 {
		return 0,"", errs[0]
	}

	if httpResp.StatusCode != 200 {
		return 0,"", fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return 0,"", fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return rsp.ContactType,rsp.Userid, nil
}


//获取部门用户详情
func Listbypage(data ListbypageRequest) (userlist []ListbypageUserlist, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()
	fmt.Println(accessToken)

	var _url=""
	if data.Size!=0 {
		_url = fmt.Sprintf("%s/user/listbypage?access_token=%s&department_id=%d&offset=%d&size=%d",
			dingtalk.ACCESS_URL, accessToken,data.DeptId,data.Offset,data.Size)
	}else {
		_url = fmt.Sprintf("%s/user/listbypage?access_token=%s&department_id=%d",
			dingtalk.ACCESS_URL, accessToken,data.DeptId)
	}


	var rsp ListbypageResponse
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

	return rsp.Userlist, nil
}


//获取部门用户
func Simplelist(data SimplelistRequest) (userlist []SimplelistUserlist, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()

	var _url=""
	if data.offset!=0&&data.size!=0 {
		_url = fmt.Sprintf("%s/user/simplelist?access_token=%s&department_id=%d&offset=%d&size=%d",
			dingtalk.ACCESS_URL, accessToken,data.deptId,data.offset,data.size)
	}else {
		_url = fmt.Sprintf("%s/user/simplelist?access_token=%s&department_id=%d",
			dingtalk.ACCESS_URL, accessToken,data.deptId)
	}


	var rsp SimplelistResponse
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

	return rsp.Userlist, nil
}

//获取部门用户userid列表
func GetDeptMember(deptId string) (userids []string, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()

	_url := fmt.Sprintf("%s/user/getDeptMember?access_token=%s&deptId=%s",dingtalk.ACCESS_URL, accessToken,deptId)

	var rsp GetDeptMemberResponse
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

	return rsp.Userids, nil
}


//根据手机号获取userid
func GetbyMobile(mobile string) (userid string, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()

	_url := fmt.Sprintf("%s/user/get_by_mobile?access_token=%s&mobile=%s",dingtalk.ACCESS_URL, accessToken,mobile)

	var rsp GetbyMobileResponse
	httpResp, _, errs := gorequest.New().Get(_url).EndStruct(&rsp)
	if len(errs) > 0 {
		return "", errs[0]
	}

	if httpResp.StatusCode != 200 {
		return "", fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return "", fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return rsp.Userid, nil
}

//获取用户详情
func Get(userid string) (rsp UserinfoResponse, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()

	_url := fmt.Sprintf("%s/user/get?access_token=%s&userid=%s",dingtalk.ACCESS_URL, accessToken,userid)

	httpResp, _, errs := gorequest.New().Get(_url).EndStruct(&rsp)
	if len(errs) > 0 {
		return UserinfoResponse{}, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return UserinfoResponse{}, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return UserinfoResponse{}, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return rsp, nil
}

//获取管理员列表
func GetAdmin() (adminlist []Adminlist, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()

	_url := fmt.Sprintf("%s/user/get_admin?access_token=%s",dingtalk.ACCESS_URL, accessToken)

	var rsp Get_adminResponse
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

	return rsp.AdminList, nil
}

