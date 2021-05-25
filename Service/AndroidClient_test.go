package Service

import (
	"fmt"
	"github.com/huangapple/go-umeng-push/Constants"
)

func ExampleAndroid_Push() {
	anClient := NewAndroidClient("5fe5679044bb94418a6496fd", "828cce9779fbf7b9011d4af70d165cae", Constants.TEST)

	anBody := Body{
		Ticker: "title",
		Title:  "subTitle",
		Text:   "Body",
	}
	anPayload := AnPayload{
		DisplayType: "message",
		Body:        anBody,
	}
	anPolicy := Policy{}
	anOption := Option{}

	anCustomized := AnCustomized{
		PushType:     Constants.LISTS_PUSH,
		DeviceTokens: []string{""}, // 当type=unicast时, 必填, 表示指定的单个设备  当type=listcast时, 必填, 要求不超过500个, 以英文逗号分隔

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

func ExampleAbstractNotification_SetApp() {
	anClient := NewAndroidClient("your app key ", "your secret", Constants.TEST)
	anClient.SetApp("set your new app", "your new secret")
	fmt.Println(anClient)
	// Output:
	//&{{set your new app your new secret test {<nil> <nil> <nil> 0}}}
}

func ExampleAbstractNotification_Upload() {
	anClient := NewAndroidClient("your app key ", "your secret", Constants.TEST)

	deviceToken := []string{
		"devcetoke1",
		"devcetoke2",
	}

	push, _ := anClient.Upload(deviceToken)
	defer push.Close()
	fmt.Println(push.IsConnectSuccess())
	fmt.Println(push.IsErrorOccur())
	fmt.Println(push.GetErrorCode())
	fmt.Println(push.GetErrorMessage())
	// Output:
	//true
	//true
	//2021
	//appkey不存在
}
