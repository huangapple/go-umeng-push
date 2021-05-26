package Service

import (
	"fmt"
	"github.com/huangapple/go-umeng-push/Constants"
)

func ExampleIOSClient_Broadcast() {
	client := NewIOSClient("your app key", "your secret", Constants.TEST)
	alert := AlertParams{
		"title", "subTitle", "Body",
	}
	aps := ApsParams{Alert: alert}
	payload := Payload{
		Aps: aps,
	}
	push, _ := client.Broadcast(&payload)
	defer push.Close()
	fmt.Println(push.IsConnectSuccess())
	fmt.Println(push.IsErrorOccur())
	fmt.Println(push.GetErrorMessage())
	fmt.Println(push.GetErrorCode())
	fmt.Println(push.GetTaskId())
	// Output:
	//true
	//true
	//appkey不存在
	//2021
}

func ExampleIOSClient_Push() {
	client := NewIOSClient("your app key", "your secret", Constants.TEST)
	alert := AlertParams{
		"title", "subTitle", "Body",
	}
	aps := ApsParams{Alert: alert}
	payload := Payload{
		Aps: aps,
	}

	customized := new(Customized)

	push, _ := client.Push(&payload, Constants.LISTS_PUSH, customized, &Option{
		Description: "",
		MiPush:      "",
		MiActivity:  "",
	})
	defer push.Close()
	fmt.Println(push.IsConnectSuccess())
	fmt.Println(push.IsErrorOccur())
	fmt.Println(push.GetErrorMessage())
	fmt.Println(push.GetErrorCode())

}
func ExampleIOSClient_PushStatus() {
	client := NewIOSClient("your app key", "your secret", Constants.TEST)
	taskId := "your taskId"
	push, _ := client.PushStatus(taskId)
	defer push.Close()
	fmt.Println(push.IsConnectSuccess())
	fmt.Println(push.IsErrorOccur())
	fmt.Println(push.GetErrorMessage())
	fmt.Println(push.GetErrorCode())
	fmt.Println(push.GetOpenCount())
	fmt.Println(push.GetTotalCount())
	fmt.Println(push.GetSentCount())
	// Output:
	//true
	//true
	//appkey不存在
	//2021
	//0
	//0
	//0
}
