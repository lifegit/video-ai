package ai

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/mitchellh/mapstructure"
	"net/url"
	"strings"
)

// fork https://help.aliyun.com/document_detail/94072.html
type ai struct {
	AppKey string
	Client *sdk.Client
}

func New(RegionId, accessKeyId, accessKeySecret, appKey string) (*ai, error) {
	client, err := sdk.NewClientWithAccessKey(RegionId, accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}

	return &ai{Client: client, AppKey: appKey}, nil
}

func (a *ai) getCommonRequest() *requests.CommonRequest {
	req := requests.NewCommonRequest()
	req.Domain = "filetrans.cn-shanghai.aliyuncs.com"
	req.Version = "2018-08-17"
	req.Product = "nls-filetrans"
	return req
}

func (a *ai) request(req *requests.CommonRequest) (res map[string]interface{}, err error) {
	response, err := a.Client.ProcessCommonRequest(req)
	if err != nil {
		return
	}
	if response.GetHttpStatus() != 200 {
		return res, fmt.Errorf("录音文件识别请求失败，Http错误码: %d", response.GetHttpStatus())
	}
	content := response.GetHttpContentString()
	err = json.Unmarshal([]byte(content), &res)
	if err != nil {
		return
	}
	statusText, _ := res["StatusText"].(string)
	if statusText != "SUCCESS" {
		return res, fmt.Errorf("StatusText is %s, is not equ SUCCESS", statusText)
	}

	return
}

func (a *ai) Send(urlPath string) (res string, err error) {
	u, _ := url.Parse(urlPath)
	fileLink := strings.ReplaceAll(u.String(), "#", "%23")

	m := map[string]string{
		"version":      "4.0",   // 新接入请使用4.0版本，已接入（默认2.0）如需维持现状，请注释掉该参数设置。
		"enable_words": "false", // 设置是否输出词信息，默认为false。开启时需要设置version为4.0。
		"appkey":       a.AppKey,
		"file_link":    fileLink,
	}
	task, err := json.Marshal(m)
	if err != nil {
		return "", err
	}

	req := a.getCommonRequest()
	req.ApiName = "SubmitTask"
	req.Method = "POST"
	req.FormParams["Task"] = string(task)

	result, err := a.request(req)
	if err != nil {
		return "", err
	}
	// {"appkey":"heMiCs0f89fXiktZ","enable_words":"false","file_link":"~","version":"4.0"}
	return result["TaskId"].(string), nil
}

type Result struct {
	EndTime         int     `json:"EndTime"`
	SilenceDuration int     `json:"SilenceDuration"`
	BeginTime       int     `json:"BeginTime"`
	Text            string  `json:"Text"`
	ChannelId       int     `json:"ChannelId"`
	SpeechRate      int     `json:"SpeechRate"`
	EmotionValue    float64 `json:"EmotionValue"`
}

func (a *ai) GetResult(taskId string) (res []Result, err error) {
	req := a.getCommonRequest()
	req.ApiName = "GetTaskResult"
	req.Method = "GET"
	req.QueryParams["TaskId"] = taskId

	result, err := a.request(req)
	if err != nil {
		return res, err
	}

	// {"TaskId":"362607741c214dea8a65995dc51ac0f0","RequestId":"885800C4-F6BD-51B2-A204-85E0CD69ABB1","StatusText":"SUCCESS","BizDuration":163538,"SolveTime":1665230088222,"RequestTime":1665230066666,"StatusCode":21050000,"Result":{"Sentences":[{"EndTime":3700,"SilenceDuration":0,"BeginTime":0,"Text":"你好","ChannelId":0,"SpeechRate":210,"EmotionValue":7.4},{"EndTime":9820,"SilenceDuration":0,"BeginTime":3980,"Text":"你好","ChannelId":0,"SpeechRate":277,"EmotionValue":7.5}]}}
	v := result["Result"].(map[string]interface{})
	err = mapstructure.Decode(v["Sentences"], &res)

	return
}
