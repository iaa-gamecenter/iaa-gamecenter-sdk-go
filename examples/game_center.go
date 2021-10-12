package examples

import (
	"context"
	"fmt"
	iaa_gamecenter_sdk_go "github.com/iaa-gamecenter/iaa-gamecenter-sdk-go"
	"github.com/iaa-gamecenter/iaa-gamecenter-sdk-go/entity"
)

func main() {
	GetUserTicket()
	QueryDeviceInfo()

}

/**
 * 获取用户ticket
 */
func GetUserTicket() {
	var ctx context.Context
	//appId和appKey可通过开放平台自主获取
	client := iaa_gamecenter_sdk_go.NewClient(ctx, "dsada","dsadasd")

	//Platform：游戏app登录返回信息中获取platform字段，不要对该字段做任何校验，直接透传即可
	//Ticket：ticket值由游戏app调用登录接口返回获取
	var req = &entity.UserTicketReq{
		Platform: "gapp",
		Ticket:   "t11YfNhmecNWb3dxu14TH",
	}
	resp, err := client.GetUserTicket(req)
	fmt.Println(resp, err)

}

/**
 * 单测查询负毛利用户
 */
func QueryDeviceInfo() {
	var ctx context.Context
	//appId和appKey可通过开放平台自主获取
	client := iaa_gamecenter_sdk_go.NewClient(ctx, "dsada","dsadasd")

	//Tuid:通过sdk获取的
	var req = &entity.DeviceInfoReq{
		Tuid: "u11YfBiAp2DPJHHAC5JMd",
	}
	resp, err := client.QueryDeviceInfo(req)
	fmt.Println(resp, err)

}