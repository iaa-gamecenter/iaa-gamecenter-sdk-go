package iaa_gamecenter_sdk_go

import (
	"context"
	"fmt"
	"github.com/iaa-gamecenter/iaa-gamecenter-sdk-go/entity"
	"github.com/iaa-gamecenter/iaa-gamecenter-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_AddCoin(t *testing.T) {
	var ctx context.Context
	client := NewClient(ctx,"dsada","dsadasd")
	var req = &entity.AddCoinReq{
		OpenId:  "u11YfBiAp2DPJHHAC5JMd",
		CoinNum: 1,
		TradeNo: "test"+util.TodayUnixStr(),
	}
	resp,err := client.AddCoin(req)
	fmt.Println(resp)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestClient_SubCoin(t *testing.T) {
	var ctx context.Context
	client := NewClient(ctx,"dsada","dsadasd")
	var req = &entity.SubCoinReq{
		OpenId:  "u11YfBiAp2DPJHHAC5JMd",
		CoinNum: 1,
		TradeNo: "test"+util.TodayUnixStr(),
	}
	resp,err :=  client.SubCoin(req)
	fmt.Println(resp)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestClient_BalanceCoin(t *testing.T) {
	var ctx context.Context
	client := NewClient(ctx,"dsada","dsadasd")
	var req = &entity.BalanceCoinReq{OpenId:"u11YfBiAp2DPJHHAC5JMd"}
	resp,err :=  client.BalanceCoin(req)
	fmt.Println(resp)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}