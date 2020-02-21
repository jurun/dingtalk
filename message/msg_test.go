package message

import (
	"github.com/jurun/dingtalk"
	"testing"
)

func Test_msg_get(t *testing.T) {

	dingtalk.Option.AppKey = "dingyirbacim1xgtfrcq"
	dingtalk.Option.AppSecret="hwAk3apBb3kVGEawMzuYFPUvxn_VBMCeb9FtO3LYTfy280Qp3ZXZ3UHcD2zMR_3L"

	//撤回
	id,err:=Recall(RecallRequest{Agent_id:469378390,Msg_task_id:124598710607})

	//id,err:=Getsendresult(GetsendprogressRequest{Agent_id:469378390,Task_id:124598710607})
	//id,err:=Getsendprogress(GetsendprogressRequest{Agent_id:469378390,Task_id:124598710607})


	if err != nil {
		t.Error(err)
		return
	}

	t.Log(id)
}


func Test_msg_send(t *testing.T) {

	dingtalk.Option.AppKey = "dingyirbacim1xgtfrcq"
	dingtalk.Option.AppSecret="hwAk3apBb3kVGEawMzuYFPUvxn_VBMCeb9FtO3LYTfy280Qp3ZXZ3UHcD2zMR_3L"

	//OA
	var form []Form
	form=append(form,Form{Key:"姓名",Value:"张三"})
	form=append(form,Form{Key:"年龄",Value:"20"})
	id,err:=SendworkMessage(AsyncsendOaRequest{Agent_id:469378390,Userid_list:"1203261440752994",
		Msg:Oamsg{Msgtype:"oa",Oa:Oa{MessageURL:"http://www.baidu.com",Head:Head{
			Bgcolor:"FFBBBBBB",Text:"头部标题"},Body:Body{Title:"body标题",Form:form,Rich:Rich{Num:"15.6",Unit:"元"}}}}})

	//整体跳转
	//id,err:=SendworkMessage(AsyncsendRequest{Agent_id:469378390,Userid_list:"1203261440752994",
	//	Msg:Cardmsg{Msgtype:"action_card",ActionCard:ActionCard{Markdown:"1",Title:"11",SingleTitle:"sss",SingleURL:"http://www.baidu.com"}}})

	//独立跳转
	//var btnlist []BtnJSONList
	//btnlist=append(btnlist,BtnJSONList{ActionURL:"http://www.126.com",Title:"111"})
	//btnlist=append(btnlist,BtnJSONList{ActionURL:"http://www.sina.com",Title:"222"})
	//id,err:=SendworkMessage(AsyncsendRequest{Agent_id:469378390,Userid_list:"1203261440752994",
	//	Msg:Cardmsg{Msgtype:"action_card",ActionCard:ActionCard{
	//		Markdown:"1",
	//		Title:"11",
	//		BtnJSONList:btnlist,
	//		BtnOrientation:"1"}}})

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(id)
}
