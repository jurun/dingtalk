package message

import (
	"fmt"
	"github.com/jurun/dingtalk"
	"github.com/parnurzeal/gorequest"
)

//OA消息相关结构体
type AsyncsendOaRequest struct {
	Agent_id int `json:"agent_id"`
	Userid_list string `json:"userid_list"`
	Msg Oamsg `json:"msg"`
}

type Oamsg struct {
	Msgtype string `json:"msgtype"`
	Oa Oa `json:"oa"`
}

type Oa struct {
	Body Body `json:"body"`
	Head Head `json:"head"`
	MessageURL string `json:"message_url"`
}

type Head struct {
	Bgcolor string `json:"bgcolor"`
	Text    string `json:"text"`
}

type Body struct {
	Author    string `json:"author"`
	Content   string `json:"content"`
	FileCount string `json:"file_count"`
	Form      []Form `json:"form"`
	Image string `json:"image"`
	Rich  Rich `json:"rich"`
	Title string `json:"title"`
}

type Form struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Rich struct {
	Num  string `json:"num"`
	Unit string `json:"unit"`
}

//普通消息response

type Send_to_conversationOaRequest struct {
	Sender string `json:"sender"`
	Cid string `json:"cid"`
	Msg Oamsg `json:"msg"`
}

type Send_to_conversationCardRequest struct {
	Sender string `json:"sender"`
	Cid string `json:"cid"`
	Msg Cardmsg `json:"msg"`
}

type Send_to_conversationResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Receiver string    `json:"receiver"`
}

//卡片消息相关结构体
type AsyncsendResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Task_id int    `json:"task_id"`
}

type AsyncsendRequest struct {
	Agent_id int `json:"agent_id"`
	Userid_list string `json:"userid_list"`
	Msg Cardmsg `json:"msg"`
}


type Cardmsg struct {
	ActionCard ActionCard `json:"action_card"`
	Msgtype string `json:"msgtype"`
}

type ActionCard struct {
	BtnJSONList []BtnJSONList `json:"btn_json_list"`
	BtnOrientation string `json:"btn_orientation"`
	Markdown    string `json:"markdown"`
	SingleTitle string `json:"single_title"`
	SingleURL   string `json:"single_url"`
	Title       string `json:"title"`
}

type BtnJSONList struct {
	ActionURL string `json:"action_url"`
	Title     string `json:"title"`
}


type RecallRequest struct {
	Agent_id int `json:"agent_id"`
	Msg_task_id int `json:"msg_task_id"`
}


//群消息相关
type GetChatResponse struct {
	ChatInfo chatInfo `json:"chat_info"`
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type chatInfo struct {
	Chatid          string   `json:"chatid"`
	ConversationTag int64    `json:"conversationTag"`
	Name            string   `json:"name"`
	Owner           string   `json:"owner"`
	Useridlist      []string `json:"useridlist"`
}


type ChatCreateResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Chatid  string `json:"chatid"`
	ConversationTag int    `json:"conversationTag"`
}

type ChatCreateRequest struct {
	Name  string `json:"name"`
	Owner  string `json:"owner"`
	Useridlist []string    `json:"useridlist"`
}

type ChatSendOaRequest struct {
	Chatid string `json:"chatid"`
	Msg Oamsg `json:"msg"`
}

type ChatSendCardRequest struct {
	Chatid string `json:"chatid"`
	Msg Cardmsg `json:"msg"`
}

type ChatSendResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	MessageId string    `json:"messageId"`
}


type ChatUpdateRequest struct {
	Chatid  string `json:"chatid"`
	Name  string `json:"name"`
	Owner  string `json:"owner"`
	ShowHistoryType  int `json:"ShowHistoryType"`
	Add_useridlist []string    `json:"add_useridlist"`
}

type GetReadListRequest struct {
	MessageId  string `json:"messageId"`
	Cursor  int `json:"cursor"`
	Size int `json:"size"`
}

type GetReadListResponse struct {
	Errcode        int64    `json:"errcode"`
	Errmsg         string   `json:"errmsg"`
	NextCursor     int64    `json:"next_cursor"`
	ReadUserIDList []string `json:"readUserIdList"`
}


//发送工作消息
func SendworkMessage(data interface{}) (taskid int, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()
	fmt.Println(accessToken)

	_url := fmt.Sprintf("%s/topapi/message/corpconversation/asyncsend_v2?access_token=%s",
		dingtalk.ACCESS_URL, accessToken)

	var rsp AsyncsendResponse
	httpResp, _, errs := gorequest.New().Post(_url).Send(data).EndStruct(&rsp)
	if len(errs) > 0 {
		return taskid, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return taskid, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return taskid, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return rsp.Task_id, nil
}


//工作通知消息撤回
func Recall(data RecallRequest) (rs bool, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()

	_url := fmt.Sprintf("%s/topapi/message/corpconversation/recall?access_token=%s",
		dingtalk.ACCESS_URL, accessToken)

	var rsp AsyncsendResponse
	httpResp, _, errs := gorequest.New().Post(_url).Send(data).EndStruct(&rsp)
	if len(errs) > 0 {
		return false, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return false, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return false, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return true, nil
}


//发送普通消息
func Send_to_conversation(data interface{}) (receiver string, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()
	fmt.Println(accessToken)

	_url := fmt.Sprintf("%s/message/send_to_conversation?access_token=%s",
		dingtalk.ACCESS_URL, accessToken)

	var rsp Send_to_conversationResponse
	httpResp, _, errs := gorequest.New().Post(_url).Send(data).EndStruct(&rsp)
	if len(errs) > 0 {
		return receiver, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return receiver, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return receiver, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return rsp.Receiver, nil
}

//群创建会话
func ChatCreate(data ChatCreateRequest) (rs ChatCreateResponse, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()
	fmt.Println(accessToken)

	_url := fmt.Sprintf("%s/chat/create?access_token=%s",
		dingtalk.ACCESS_URL, accessToken)

	var rsp ChatCreateResponse
	httpResp, _, errs := gorequest.New().Post(_url).Send(data).EndStruct(&rsp)
	if len(errs) > 0 {
		return rs, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return rs, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return rs, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return rsp, nil
}


//发送群消息
func Chat_Send(data interface{}) (messageId string, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()
	fmt.Println(accessToken)

	_url := fmt.Sprintf("%s/chat/send?access_token=%s",
		dingtalk.ACCESS_URL, accessToken)

	var rsp ChatSendResponse
	httpResp, _, errs := gorequest.New().Post(_url).Send(data).EndStruct(&rsp)
	if len(errs) > 0 {
		return messageId, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return messageId, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return messageId, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return rsp.MessageId, nil
}


//群修改会话
func ChatUpdate(data ChatUpdateRequest) (rs bool, err error) {

	accessToken, err := dingtalk.AccessToken.GetToken()
	fmt.Println(accessToken)

	_url := fmt.Sprintf("%s/chat/update?access_token=%s",
		dingtalk.ACCESS_URL, accessToken)

	var rsp ChatCreateResponse
	httpResp, _, errs := gorequest.New().Post(_url).Send(data).EndStruct(&rsp)
	if len(errs) > 0 {
		return false, errs[0]
	}

	if httpResp.StatusCode != 200 {
		return false, fmt.Errorf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)
	}

	if rsp.Errcode != 0 {
		return false, fmt.Errorf("接口调用失败，errcode: %d，errmsg: %s", rsp.Errcode, rsp.Errmsg)
	}

	return true, nil
}

