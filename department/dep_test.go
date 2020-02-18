package department

import (
	"github.com/jurun/dingtalk"
	"testing"
)


func Test_dep_get(t *testing.T) {

	dingtalk.Option.AppKey = "dingyirbacim1xgtfrcq"
	dingtalk.Option.AppSecret="hwAk3apBb3kVGEawMzuYFPUvxn_VBMCeb9FtO3LYTfy280Qp3ZXZ3UHcD2zMR_3L"


	//id,err:=Get("1")
	id,err:=List("")
	//id,err:=List_ids("1")

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(id)
}
