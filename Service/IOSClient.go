package Service

import (
	"github.com/huangapple/go-umeng-push/Constants"
	"github.com/huangapple/go-umeng-push/Responses/Status"
	"github.com/huangapple/go-umeng-push/Responses/TaskPush"
	"github.com/huangapple/go-umeng-push/Responses/UniCast"
	"strings"
	"time"
)

type IOSClient struct {
	abstractNotification
}

func NewIOSClient(appKey, appSecret, envMode string) *IOSClient {
	notification := newNotification(appKey, appSecret, envMode)
	ios := IOSClient{*notification}
	return &ios
}

type AlertParams struct {
	Title    string `json:"title,omitempty"`
	SubTitle string `json:"subTitle,omitempty"`
	Body     string `json:"body,omitempty"`
}
type ApsParams struct {
	Alert            AlertParams `json:"alert"`                       // 当content-available=1时(静默推送)，可选; 否则必填。
	Badge            string      `json:"badge,omitempty"`             // 可选
	Sound            string      `json:"sound,omitempty"`             // 可选
	ContentAvailable int         `json:"content-available,omitempty"` // 可选，代表静默推送
	Category         string      `json:"category,omitempty"`          // 可选，注意: ios8才支持该字段。

	MutableContent int `json:"mutable-content,omitempty"`
}
type PolicyParams struct {
	StartTime      string `json:"start_time,omitempty"`       // 可选，定时发送时间，若不填写表示立即发送。
	ExpireTime     string `json:"expire_time,omitempty"`      // 可选，消息过期时间，其值不可小于发送时间或者
	OutBizNo       string `json:"out_biz_no,omitempty"`       // 可选，开发者对消息的唯一标识，服务器会根据这个标识避免重复发送。
	ApnsCollapseId string `json:"apns_collapse_id,omitempty"` // 可选，多条带有相同apns_collapse_id的消息，iOS设备仅展示
}
type Payload struct {
	Aps  ApsParams `json:"aps"`
	Link string    `json:"link,omitempty"`
}

type Customized struct {
	DeviceTokens []string `json:"device_tokens"` /// 当type=unicast时, 必填, 表示指定的单个设备 当type=listcast时, 必填, 要求不超过500个, 以英文逗号分隔
	AliasType    string   `json:"alias_type"`    // 当type=customizedcast时, 必填
	Alias        string   `json:"alias"`         // 当type=customizedcast时, 选填(此参数和file_id二选一)
	FileId       string   `json:"file_id"`       // 当type=filecast时，必填，file内容为多条device_token，以回车符分割
	Description  string   `json:"description"`   // 可选，发送消息描述，建议填写。
	Filter       string   `json:"file_id"`       // 当type=groupcast时，必填，用户筛选条件，如用户标签、渠道等，参考附录G。@see https://developer.umeng.com/docs/66632/detail/68343#h2--g-14
}

//廣播
func (c *IOSClient) Broadcast(p *Payload) (*TaskPush.TaskPush, error) {
	var result TaskPush.TaskPush
	var err error
	params, err := c.getParams(p, nil, Constants.BROADCAST, &Customized{})

	if err != nil {
		return &result, err
	}

	response, err := c.abstractNotification.sent(Constants.HOST_URL+Constants.PUSH_URI, params)

	if err != nil {
		return &result, err
	}
	return TaskPush.New(response)
}

// 單一裝置推播
//func (c *IOSClient) Push(p Payload, deviceToken string) (result *UniCast.UniCast, err error) {
//	params, err := c.getParams(p, Constants.UNICAST, Customized{DeviceTokens: deviceToken})
//	if err != nil {
//		return result, err
//	}
//	response, err := c.abstractNotification.sent(Constants.HOST_URL+Constants.PUSH_URI, params)
//	if err != nil {
//		return result, err
//	}
//
//	return UniCast.New(response)
//
//}

func (c *IOSClient) Push(payload *Payload, policy *Policy, pushType string, customized *Customized) (response *UniCast.UniCast, err error) {
	params, err := c.getParams(payload, policy, pushType, customized)
	if err != nil {
		return response, err
	}
	httpResponse, err := c.sent(Constants.HOST_URL+Constants.PUSH_URI, params)

	return UniCast.New(httpResponse)
}

func (c *IOSClient) PushByDeviceTokens(description, title, content, path string, deviceTokens []string) (response *UniCast.UniCast, err error) {
	alert := AlertParams{
		Title: title,
		Body:  content,
	}
	aps := ApsParams{
		Alert: alert,
		Sound: "default",
		//MutableContent: 1,
	}
	payload := &Payload{
		Aps:  aps,
		Link: path,
	}

	customized := &Customized{
		DeviceTokens: deviceTokens,
		Description:  description,
	}
	policy := &Policy{
		ExpireTime: time.Now().AddDate(0, 0, 3).Format("2006-01-02 15:04:05"),
	}

	return c.Push(payload, policy, Constants.LISTS_PUSH, customized)
}

func (c *IOSClient) getParams(p *Payload, policy *Policy, pushType string, customized *Customized) (map[string]interface{}, error) {

	params := map[string]interface{}{
		"appkey":        c.abstractNotification.appKey,
		"timestamp":     time.Now().Unix() * 1000,
		"type":          pushType,
		"device_tokens": strings.Join(customized.DeviceTokens, ","),
		"description":   customized.Description,
		"payload":       p,
	}
	if policy != nil {
		params["policy"] = policy
	}

	return params, nil

}

//群組推播
func (c *IOSClient) PushStatus(taskId string) (result *StatusResponse.IOSStatusResponse, err error) {
	response, err := c.statusQuery(taskId)
	if err != nil {
		return result, err

	}
	return StatusResponse.NewIOSStatusResponse(response)

}
