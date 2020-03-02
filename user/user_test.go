package user

import (
	"github.com/jurun/dingtalk"
	"testing"
)


func Test_user_get(t *testing.T) {

	dingtalk.Option.AppKey = "dingyirbacim1xgtfrcq"
	dingtalk.Option.AppSecret="hwAk3apBb3kVGEawMzuYFPUvxn_VBMCeb9FtO3LYTfy280Qp3ZXZ3UHcD2zMR_3L"


	//_,id,err:=GetUseridByUnionid("GraZjx1CJ1cpzrdWc4PSegiEiE")
	id,err:=Listbypage(ListbypageRequest{DeptId:1,Offset:0,Size:20})
	//id,err:=Simplelist(SimplelistRequest{deptId:1,offset:1,size:1})
	//id,err:=GetDeptMember("1")
	//id,err:=GetbyMobile("15901084927")
	//id,err:=Get("1203261440752994")
	//id,err:=GetAdmin()

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(id)
}
