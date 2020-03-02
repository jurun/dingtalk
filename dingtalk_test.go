package dingtalk

import (
	"fmt"
	"github.com/xinzf/xunray/storage"
	"testing"
	"time"
)

func Test_accessToken_GetToken(t *testing.T) {

	Option.AppKey = "dingyirbacim1xgtfrcq"
	Option.AppSecret = "hwAk3apBb3kVGEawMzuYFPUvxn_VBMCeb9FtO3LYTfy280Qp3ZXZ3UHcD2zMR_3L"
	Option.Cache = new(DingTalkCache)

	token, err := AccessToken.GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(token)
}

const ACCESS_TOKEN_KEY = "dd:token"

type DingTalkCache struct {
}

func (this *DingTalkCache) Get(appKey string) (string, error) {
	key := fmt.Sprintf("%s:%s", ACCESS_TOKEN_KEY, appKey)
	return storage.Redis.Client().Get(key).Result()
}

func (this *DingTalkCache) Set(appKey, token string, expiration time.Duration) error {
	key := fmt.Sprintf("%s:%s", ACCESS_TOKEN_KEY, appKey)
	return storage.Redis.Client().Set(key, token, expiration).Err()
}
