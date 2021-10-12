package iaa_gamecenter_sdk_go

import (
	"context"
	"fmt"
	"github.com/iaa-gamecenter/iaa-gamecenter-sdk-go/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * 单测通过用户Ticket获取openId
 */
func TestGetUserTicket(t *testing.T)  {
    var ctx context.Context
	client := NewClient(ctx,"dsada","dsadasd")

	var req  = &entity.UserTicketReq{
		Platform:"gapp",
		Ticket:"t11YfNhmecNWb3dxu14TH",
	}
    //o8ycF6it0wUJo7NfBlmU6mQ4QbfY
	resp,err := client.GetUserTicket(req)
	fmt.Println(resp)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

}

/**
 * 单测查询负毛利用户
 */
func TestClient_QueryDeviceInfo(t *testing.T) {
	var ctx context.Context
	client := NewClient(ctx,"dsada","dsadasd")

	var req  = &entity.DeviceInfoReq{
		Tuid:"u11YfBiAp2DPJHHAC5JMd",
	}
	resp,err := client.QueryDeviceInfo(req)
	fmt.Println(resp)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

}