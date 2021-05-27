package Service

import (
	"fmt"
	"github.com/huangapple/go-umeng-push/Constants"
)

func ExampleAndroid_PushByDeviceTokens() {
	anClient := NewAndroidClient("xxxx", "xxxxx", Constants.PRODUCT)

	push, _ := anClient.PushByDeviceTokens("任务描述", "标题", "内容", "/pathxxx", []string{"fdsafs", "fdsafdas"})

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
