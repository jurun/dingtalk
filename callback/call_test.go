package callback

import (
	"github.com/jurun/dingtalk"
	"testing"
)


func Test_data(t *testing.T) {

	dingtalk.Option.AppKey = "dingyirbacim1xgtfrcq"
	dingtalk.Option.AppSecret="hwAk3apBb3kVGEawMzuYFPUvxn_VBMCeb9FtO3LYTfy280Qp3ZXZ3UHcD2zMR_3L"

	st:=make([]string,0)
	st=append(st,"user_add_org")
	//id,err:= Register_call_back(RegistercallbackRequest{AesKey:"xxxxxxxxlvdhntotr3x9qhlbytb18zyz5zxxxxxxxxx",Token:"2",CallBackTag:st,URL:"http://www.baidu.com"})
	//id,err:= Update_call_back(RegistercallbackRequest{AesKey:"xxxxxxxxlvdhntotr3x9qhlbytb18zyz5zxxxxxxxxx",Token:"2",CallBackTag:st,URL:"http://www.baidu.com"})
	//id,err:= Delete_call_back()
	//id,err:= Get_call_back()
	id,err:= Get_call_back_failed_result()

	if err != nil {
		t.Error(err,"33333333")
		return
	}

	t.Log(id,"4444444444")
}
