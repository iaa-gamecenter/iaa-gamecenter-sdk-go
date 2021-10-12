package iaa_gamecenter_sdk_go

import (
	"fmt"
	"github.com/iaa-gamecenter/iaa-gamecenter-sdk-go/entity"
	"github.com/iaa-gamecenter/iaa-gamecenter-sdk-go/util"
	"net/url"
	"strconv"
)
/**
1、找商务申请金币cointype，sceneid以及设置金币风控。
2、扣除趣头条金币时，须有二次弹框确认，说明消耗的是趣头条金币。
3、游戏测试前须找商务申请金币cointype，sceneid，报备金币增减项及数量
4、10000金币=1元rmb
 */


/**
 * @desc 增加金币
 * @note 无
 * @km https://dev.iggrowth.cn/sdk/docs/index.html#/server/api_coin?id=%e5%a2%9e%e5%8a%a0%e9%87%91%e5%b8%81
 */
func (c *Client) AddCoin(req *entity.AddCoinReq) (resp *entity.AddCoinResp, err error)  {
	// 参数检查
	if err := req.IsValid(); err != nil {
		return nil, err
	}

	var data = url.Values{
		"app_id":[]string{c.InitConfig.AppInfo.AppId},
		"open_id":[]string{req.OpenId},
		"coin_num":[]string{strconv.Itoa(req.CoinNum)},
		"trade_no":[]string{req.TradeNo},
		"time":[]string{util.TodayUnixStr()},
	}
	data.Add("sign",util.GenerateSign(c.InitConfig.AppInfo.AppKey,data))
	urlStr := fmt.Sprintf("%s/x/open/coin/add", c.InitConfig.RequestUrl)
	resp = &entity.AddCoinResp{}

	result,err := c.RetryBreakerPost(urlStr,data,"addCoin",resp)

	if err != nil{
		return nil,err
	}

	return result.(*entity.AddCoinResp),nil

}


/**
 * @desc 扣减金币
 * @note 无
 * @km https://dev.iggrowth.cn/sdk/docs/index.html#/server/api_coin?id=%e6%89%a3%e5%87%8f%e9%87%91%e5%b8%81
 */
func (c *Client) SubCoin(req *entity.SubCoinReq) (resp *entity.SubCoinResp, err error)  {
	// 参数检查
	if err := req.IsValid(); err != nil {
		return nil, err
	}

	var data = url.Values{
		"app_id":[]string{c.InitConfig.AppInfo.AppId},
		"open_id":[]string{req.OpenId},
		"coin_num":[]string{strconv.Itoa(req.CoinNum)},
		"trade_no":[]string{req.TradeNo},
		"time":[]string{util.TodayUnixStr()},
	}
	data.Add("sign",util.GenerateSign(c.InitConfig.AppInfo.AppKey,data))
	urlStr := fmt.Sprintf("%s/x/open/coin/sub", c.InitConfig.RequestUrl)
	resp = &entity.SubCoinResp{}

	result,err := c.RetryBreakerPost(urlStr,data,"subCoin",resp)

	if err != nil{
		return nil,err
	}

	return result.(*entity.SubCoinResp),nil

}

/**
 * @desc 查询金币
 * @note 无
 * @km https://dev.iggrowth.cn/sdk/docs/index.html#/server/api_coin?id=%e6%9f%a5%e8%af%a2%e9%87%91%e5%b8%81
 */
func (c *Client) BalanceCoin(req *entity.BalanceCoinReq) (resp *entity.BalanceCoinResp, err error)  {
	// 参数检查
	if err := req.IsValid(); err != nil {
		return nil, err
	}

	var data = url.Values{
		"app_id":[]string{c.InitConfig.AppInfo.AppId},
		"open_id":[]string{req.OpenId},
		"time":[]string{util.TodayUnixStr()},
	}
	data.Add("sign",util.GenerateSign(c.InitConfig.AppInfo.AppKey,data))
	urlStr := fmt.Sprintf("%s/x/open/coin/balance", c.InitConfig.RequestUrl)
	resp = &entity.BalanceCoinResp{}

	result,err := c.RetryBreakerPost(urlStr,data,"balanceCoin",resp)

	if err != nil{
		return nil,err
	}

	return result.(*entity.BalanceCoinResp),nil

}