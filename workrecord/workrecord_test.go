package workrecord


import (
	"github.com/jurun/dingtalk"
	"testing"
)

func Test_workrecord_add(t *testing.T) {

	dingtalk.Option.AppKey = "dingyirbacim1xgtfrcq"
	dingtalk.Option.AppSecret="hwAk3apBb3kVGEawMzuYFPUvxn_VBMCeb9FtO3LYTfy280Qp3ZXZ3UHcD2zMR_3L"

	//add
	//id,err:=Add(AddRequest{Userid:"1203261440752994",Create_time:1496678400000, Title:"1112", Url:"http://www.baidu.com",FormItemList:FormItemList{Title:"2222", Content:"cccccccccc"}})
	//update
	id,err:=Update(UpdateRequest{Userid:"1203261440752994",Record_id:"record83a653f77976e81f78a79f89db351add"})

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(id)
}
