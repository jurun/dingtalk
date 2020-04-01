package checkin
import (
	"fmt"
	"github.com/jurun/dingtalk"
	"testing"
)

func Test_data(t *testing.T) {

	dingtalk.Option.AppKey = "dingyirbacim1xgtfrcq"
	dingtalk.Option.AppSecret="hwAk3apBb3kVGEawMzuYFPUvxn_VBMCeb9FtO3LYTfy280Qp3ZXZ3UHcD2zMR_3L"

	//id,err:=Record(RecordRequest{Department_id:"1",Start_time:1582992000000,End_time:1585756800000,Offset:0,Size:10})

	id,err:=GetRecord(GetRecordRequest{Userid_list:"1203261440752994",Start_time:1585497600000,End_time:1585756800000,Cursor:0,Size:10})

	if err != nil {
		fmt.Println(err)
		//t.Error(err)
		return
	}

	t.Log(id)
}
