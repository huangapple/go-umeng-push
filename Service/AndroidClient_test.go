package Service

import (
	"fmt"
	"github.com/huangapple/go-umeng-push/Constants"
)

func ExampleAndroid_Push() {
	anClient := NewAndroidClient("5fe5679044bb94418a6496fd", "comdfsdcuaxtphtphgxjzemweictkdqr", Constants.PRODUCT)

	anBody := Body{
		Ticker: "title",
		Title:  "subTitle",
		Text:   "Body",
	}
	anPayload := AnPayload{
		DisplayType: "notification",
		Body:        anBody,
	}
	anPolicy := Policy{}
	anOption := Option{}

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
