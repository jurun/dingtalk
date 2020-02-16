package dingtalk

import (
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"time"
)

const (
	ACCESS_TOKEN_KEY     = "dd:token"
	ACCESS_TOKEN_EXPIRES = int64(7000)
)

var Option option

type option struct {
	AppKey, AppSecret string
}

var AccessToken *_accessToken

func init() {
	if AccessToken == nil {
		AccessToken = &_accessToken{
			cache: new(defaultCache),
		}
	}
}

type accessTokenCache interface {
	Get(key string) (string, error)
	Set(key, token string, expiration time.Duration) error
}

type defaultCache struct {
}

func (this *defaultCache) Get(key string) (string, error) {
	return "", nil
}

func (this *defaultCache) Set(key, token string, expiration time.Duration) error {
	return nil
}

type _accessToken struct {
	cache accessTokenCache
}

func (this *_accessToken) SetCache(cache accessTokenCache) {
	this.cache = cache
}

func (this *_accessToken) GetToken() (string, error) {
	if Option.AppKey == "" {
		return "", errors.New("没有设置 appKey")
	}
	if Option.AppSecret == "" {
		return "", errors.New("没有设置 appSecret")
	}

	key := fmt.Sprintf("%s:%s", ACCESS_TOKEN_KEY, Option.AppKey)
	token, err := this.cache.Get(key)
	if err != nil {
		return "", err
	}
	if token != "" {
		fmt.Println("命中缓存")
		return token, nil
	}
	fmt.Println("缓存没有命中，重新获取")

	var rsp struct {
		Errcode     int    `json:"errcode"`
		Errmsg      string `json:"errmsg"`
		AccessToken string `json:"access_token"`
	}
	_url := fmt.Sprintf("https://oapi.dingtalk.com/gettoken?appkey=%s&appsecret=%s", Option.AppKey, Option.AppSecret)
	_rsp, _, errs := gorequest.New().Get(_url).EndStruct(&rsp)
	if _rsp.StatusCode != 200 {
		return "", fmt.Errorf("钉钉服务器异常,httpStatus: %d", _rsp.StatusCode)
	}

	if len(errs) > 0 {
		return "", errs[0]
	}

	if rsp.Errcode != 0 {
		return "", errors.New(rsp.Errmsg)
	}

	err = this.cache.Set(
		key,
		rsp.AccessToken,
		time.Duration(ACCESS_TOKEN_EXPIRES)*time.Second,
	)
	if err != nil {
		return "", err
	}
	return rsp.AccessToken, nil
}
