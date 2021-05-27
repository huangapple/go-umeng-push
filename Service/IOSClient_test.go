package Service

import (
	"fmt"
	"github.com/huangapple/go-umeng-push/Constants"
	"testing"
)

func TestIOSClient_PushByDeviceTokens(t *testing.T) {
	client := NewIOSClient("xxx", "xxxxxx", Constants.PRODUCT)

	push, _ := client.PushByDeviceTokens("任务描述", "标题", "内容", "/pathxxx", []string{"fdsa", "fdsafdsadfd"})
	defer push.Close()
	fmt.Println(push.IsConnectSuccess())
	fmt.Println(push.IsErrorOccur())
	fmt.Println(push.GetErrorMessage())
	fmt.Println(push.GetErrorCode())

}
