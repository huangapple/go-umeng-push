package Service

import (
	"fmt"
	"github.com/huangapple/go-umeng-push/Constants"
	"time"
)

func ExampleAndroid_Push() {
	anClient := NewAndroidClient("xxxx", "xxxxx", Constants.PRODUCT)

	anBody := Body{
		Ticker:    "title",
		Title:     "subTitle",
		Text:      "Body",
		AfterOpen: "go_activity",
		Activity:  "/path/xxxxx",
		PlaySound: true,
	}
	anPayload := AnPayload{
		DisplayType: "notification",
		Body:        anBody,
		Extra: map[string]interface{}{
			"link": "/path/xxxxx",
		},
	}
	anPolicy := Policy{
		ExpireTime: time.Now().AddDate(0, 0, 3).Format("2006-01-02 15:04:05"),
	}
	anOption := Option{
		Description: "这是任务描述",
		MiPush:      true,
		MiActivity:  "/path/xxxx",
	}

	anCustomized := AnCustomized{
		PushType:     Constants.LISTS_PUSH,
		DeviceTokens: []string{"ArlDNnd5bPmDJH_Iz1IkKN3i2E6qCWopSQWkbn6JLJLU"}, // 当type=unicast时, 必填, 表示指定的单个设备  当type=listcast时, 必填, 要求不超过500个, 以英文逗号分隔

	}

	push, _ := anClient.Push(&anPayload, &anPolicy, &anCustomized, &anOption)

	defer push.Close()

	fmt.Println(push.IsConnectSuccess())
	fmt.Println(push.IsErrorOccur())
	fmt.Println(push.GetErrorCode())
	fmt.Println(push.GetErrorMessage())

	// Output:
	// true
	// true
	// 2021
	// appkey不存在
}
