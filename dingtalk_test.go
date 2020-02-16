package dingtalk

import "testing"

func Test_accessToken_GetToken(t *testing.T) {

    Option.AppKey = "dingyirbacim1xgtfrcq"
    Option.AppSecret="hwAk3apBb3kVGEawMzuYFPUvxn_VBMCeb9FtO3LYTfy280Qp3ZXZ3UHcD2zMR_3L"

    token,err:=AccessToken.GetToken()
    if err != nil {
        t.Error(err)
        return
    }

    t.Log(token)
}
