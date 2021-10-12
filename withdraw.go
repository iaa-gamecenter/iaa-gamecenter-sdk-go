package iaa_gamecenter_sdk_go

import (
	"fmt"
	"github.com/iaa-gamecenter/iaa-gamecenter-sdk-go/entity"
	"github.com/iaa-gamecenter/iaa-gamecenter-sdk-go/util"
	"net/url"
	"strconv"
)

/**
 * https://dev.iggrowth.cn/sdk/docs/index.html#/server/api_withdraw?id=show_withdraw
 * 用户绑定信息及支持的提现方式查询
 */
func (c *Client) ShowWithDraw(req *entity.ShowWithDrawReq) (resp *entity.ShowWithDrawResp, err error) {
	// 参数检查
	if err := req.IsValid(); err != nil {
		return nil, err
	}

	var data = url.Values{
		"app_id":    []string{c.InitConfig.AppInfo.AppId},
		"open_id":   []string{req.OpenId},
		"timestamp": []string{util.TodayUnixStr()},
	}
	data.Add("sign", util.GenerateSign(c.InitConfig.AppInfo.AppKey, data))
	urlStr := fmt.Sprintf("%s/x/gapp/withdraw/show", c.InitConfig.RequestUrl)
	resp = &entity.ShowWithDrawResp{}

	result, err := c.RetryBreakerPost(urlStr, data, "showWithDraw", resp)

	if err != nil {
		return nil, err
	}

	return result.(*entity.ShowWithDrawResp), nil

}

/**
 * @desc 提现下单接口
 * @note 游戏方确认订单及用户无异常情况后向游戏中心发起提现接口
 * @km https://dev.iggrowth.cn/sdk/docs/index.html#/server/api_withdraw?id=do_withdraw
 */

func (c *Client) CreateWithDraw(req *entity.CreateWithDrawReq) (resp *entity.CreateWithDrawResp, err error) {
	// 参数检查
	if err := req.IsValid(); err != nil {
		return nil, err
	}

	var data = url.Values{
		"app_id":      []string{c.InitConfig.AppInfo.AppId},
		"open_id":     []string{req.OpenId},
		"amount":      []string{strconv.Itoa(req.Amount)},
		"ship_delay":  []string{strconv.Itoa(req.ShipDelay)},
		"notify_url":  []string{req.NotifyUrl},
		"tk":          []string{req.Tk},
		"tuid":        []string{req.Tuid},
		"tp_way":      []string{req.TpWay},
		"cp_order_no": []string{req.CpOrderNo},
	}
	if req.Remark != "" {
		data.Add("remark", req.Remark)
	}
	data.Add("timestamp", util.TodayUnixStr())
	data.Add("sign", util.GenerateSign(c.InitConfig.AppInfo.AppKey, data))
	urlStr := fmt.Sprintf("%s/x/gapp/withdraw/create", c.InitConfig.RequestUrl)
	resp = &entity.CreateWithDrawResp{}

	result, err := c.RetryBreakerPost(urlStr, data, "createWithDraw", resp)

	if err != nil {
		return nil, err
	}

	return result.(*entity.CreateWithDrawResp), nil

}

/**
 * @desc 查询提现订单接口
 * @note 游戏方查询提现订单的接口（只提供两个月内的订单查询，当月和上月）
 * @km https://dev.iggrowth.cn/sdk/docs/index.html#/server/api_withdraw?id=%e6%9f%a5%e8%af%a2%e6%8f%90%e7%8e%b0%e8%ae%a2%e5%8d%95%e6%8e%a5%e5%8f%a3
 */
func (c *Client) QueryWithDraw(req *entity.QueryWithDrawReq) (resp *entity.QueryWithDrawResp, err error) {
	// 参数检查
	if err := req.IsValid(); err != nil {
		return nil, err
	}

	var data = url.Values{
		"app_id":      []string{c.InitConfig.AppInfo.AppId},
		"open_id":     []string{req.OpenId},
		"order_no":    []string{req.OrderNo},
		"cp_order_no": []string{req.CpOrderNo},
		"timestamp":   []string{util.TodayUnixStr()},
	}

	data.Add("sign", util.GenerateSign(c.InitConfig.AppInfo.AppKey, data))
	urlStr := fmt.Sprintf("%s/x/gapp/withdraw/query", c.InitConfig.RequestUrl)
	resp = &entity.QueryWithDrawResp{}

	result, err := c.RetryBreakerPost(urlStr, data, "showWithDraw", resp)

	if err != nil {
		return nil, err
	}

	return result.(*entity.QueryWithDrawResp), nil

}
