package attendance
import (
	"github.com/jurun/dingtalk"
	"testing"
)

func Test_att_get(t *testing.T) {
	dingtalk.Option.AppKey = "dingyirbacim1xgtfrcq"
	dingtalk.Option.AppSecret="hwAk3apBb3kVGEawMzuYFPUvxn_VBMCeb9FtO3LYTfy280Qp3ZXZ3UHcD2zMR_3L"
	us:=make([]int,0)
	us=append(us,1203261440752994)

	//id,err:=ListRecord(ListRecordRequest{UserIds:us,CheckDateFrom:"2020-02-25 22:10:53",CheckDateTo:"2020-02-26 22:10:53"})
	id,err:=List(ListRequest{UserIdList:us,WorkDateFrom:"2020-02-25 22:10:53",
		Offset:0,Limit:1,
		WorkDateTo:"2020-02-26 22:10:53"})

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(id)
}
