package Service

import (
	"fmt"
	"github.com/huangapple/go-umeng-push/Constants"
	"testing"
	"time"
)

func TestIOSClient_Push(t *testing.T) {
	client := NewIOSClient("xxx", "xxxxxx", Constants.PRODUCT)
	alert := AlertParams{
		"标题", "", "通知内容",
	}
	aps := ApsParams{
		Alert:          alert,
		Sound:          "default",
		MutableContent: 1,
		Link:           "/path/video/de",
	}
	payload := Payload{
		Aps: aps,
	}

	customized := &Customized{
		DeviceTokens: []string{"ArlDNnd5bPmDJH_Iz1IkKN3i2E6qCWopSQWkbn6JLJLU"},
		Description:  "任务描述",
	}
	policy := &Policy{
		ExpireTime: time.Now().AddDate(0, 0, 3).Format("2006-01-02 15:04:05"),
	}

	push, _ := client.Push(&payload, policy, Constants.LISTS_PUSH, customized)
	defer push.Close()
	fmt.Println(push.IsConnectSuccess())
	fmt.Println(push.IsErrorOccur())
	fmt.Println(push.GetErrorMessage())
	fmt.Println(push.GetErrorCode())

}

func TestIOSClient_PushByDeviceTokens(t *testing.T) {
	client := NewIOSClient("xxx", "xxxxxx", Constants.PRODUCT)

	push, _ := client.PushByDeviceTokens("任务描述", "标题", "内容", "/pathxxx", []string{"fdsa", "fdsafdsadfd"})
	defer push.Close()
	fmt.Println(push.IsConnectSuccess())
	fmt.Println(push.IsErrorOccur())
	fmt.Println(push.GetErrorMessage())
	fmt.Println(push.GetErrorCode())

}
