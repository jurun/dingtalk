package calendar

import (
	"fmt"
	"github.com/jurun/dingtalk"
	"testing"
)


func Test_workrecord_add(t *testing.T) {

	dingtalk.Option.AppKey = "dingyirbacim1xgtfrcq"
	dingtalk.Option.AppSecret="hwAk3apBb3kVGEawMzuYFPUvxn_VBMCeb9FtO3LYTfy280Qp3ZXZ3UHcD2zMR_3L"

	//add
	//id,err:=Add(AddRequest{Userid:"1203261440752994",Create_time:1496678400000, Title:"1112", Url:"http://www.baidu.com",FormItemList:FormItemList{Title:"2222", Content:"cccccccccc"}})
	//update
	//id,err:=Update(UpdateRequest{Userid:"1203261440752994",Record_id:"record83a653f77976e81f78a79f89db351add"})

	id,err:=Add(AddRequest{CreateVo:CreateVo{BizID:"2",CalendarType:"meeting",
		CreatorUserid:"1203261440752994",EndTime:EndTime{UnixTimestamp:"1585751672"},
		ReceiverUserids:"1203261440752994",Source:Source{Title:"222",URL:"http://www.baidu.com"},
		StartTime:StartTime{UnixTimestamp:"1585665272930"},Summary:"1111",UUID:"2"}})

	if err != nil {
		fmt.Println(err)
		//t.Error(err)
		return
	}

	t.Log(id)
}
