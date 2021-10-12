package examples

import (
	"context"
	"fmt"
	"github.com/iaa-gamecenter/iaa-gamecenter-sdk-go"
	"github.com/iaa-gamecenter/iaa-gamecenter-sdk-go/entity"
	"github.com/iaa-gamecenter/iaa-gamecenter-sdk-go/util"
)

func main() {
	addCoin()
	subCoin()
	balanceCoin()
}
/**
 * 添加金币
 */
func addCoin() {
	var ctx context.Context
	//appId和appKey可通过开放平台自主获取
	client := iaa_gamecenter_sdk_go.NewClient(ctx,"dsada","dsadasd")
	//通过ticket来获取OpenId
	//CoinNum：增加金币数值
	//TradeNo：订单号，保持唯一，以便对账
	var req = &entity.AddCoinReq{
		OpenId:  "u11YfBiAp2DPJHHAC5JMd",
		CoinNum: 1,
		TradeNo: "test"+util.TodayUnixStr(),
	}
	resp,err := client.AddCoin(req)
	fmt.Println(resp,err)
}

/**
 * 扣减金币
 */
func subCoin() {
	var ctx context.Context
	//appId和appKey可通过开放平台自主获取
	client := iaa_gamecenter_sdk_go.NewClient(ctx,"dsada","dsadasd")
	//通过ticket来获取OpenId
    //CoinNum：扣减金币数值
    //TradeNo：订单号，保持唯一，以便对账
	var req = &entity.SubCoinReq{
		OpenId:  "u11YfBiAp2DPJHHAC5JMd",
		CoinNum: 1,
		TradeNo: "test"+util.TodayUnixStr(),
	}
	resp,err :=  client.SubCoin(req)
	fmt.Println(resp,err)
}


/**
 * 获取金币余额
 */
func balanceCoin() {
	var ctx context.Context
	//appId和appKey可通过开放平台自主获取
	client := iaa_gamecenter_sdk_go.NewClient(ctx,"dsada","dsadasd")
	//通过ticket来获取OpenId
	var req = &entity.BalanceCoinReq{OpenId:"u11YfBiAp2DPJHHAC5JMd"}
	resp,err :=  client.BalanceCoin(req)
	fmt.Println(resp,err)
}