package role

import (
	"github.com/jurun/dingtalk"
	"testing"
)


func Test_user_get(t *testing.T) {

	dingtalk.Option.AppKey = "dingyirbacim1xgtfrcq"
	dingtalk.Option.AppSecret="hwAk3apBb3kVGEawMzuYFPUvxn_VBMCeb9FtO3LYTfy280Qp3ZXZ3UHcD2zMR_3L"


	//id,err:=Getrole(461831703)
	id,err:=Getrolegroup(461831700)
	//id,err:=Simplelist(SimplelistRequest{roleId:461831704,offset:1,size:1})
	//id,err:=List(ListRequest{offset:1,size:1})


	if err != nil {
		t.Error(err)
		return
	}

	t.Log(id)
}
