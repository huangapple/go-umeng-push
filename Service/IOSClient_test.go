package Service

import (
	"fmt"
	"github.com/huangapple/go-umeng-push/Constants"
	"testing"
)

func TestIOSClient_Push(t *testing.T) {
	client := NewIOSClient("5fe567c9adb42d58268e0f73", "crskl8vywizfgstmfjgmjnviyumktvzg", Constants.PRODUCT)
	alert := AlertParams{
		"title", "subTitle", "Body",
	}
	aps := ApsParams{Alert: alert}
	payload := Payload{
		Aps: aps,
	}

	customized := &Customized{
		DeviceTokens: []string{"ArlDNnd5bPmDJH_Iz1IkKN3i2E6qCWopSQWkbn6JLJLU"},
	}

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
