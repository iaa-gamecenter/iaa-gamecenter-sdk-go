package iaa_gamecenter_sdk_go

import (
	"context"
	"fmt"
	"github.com/iaa-gamecenter/iaa-gamecenter-sdk-go/config"
	"github.com/iaa-gamecenter/iaa-gamecenter-sdk-go/entity"
	"github.com/iaa-gamecenter/iaa-gamecenter-sdk-go/internal/clients"
	"github.com/iaa-gamecenter/iaa-gamecenter-sdk-go/util"
	"net/url"
	"strconv"
)

type Client struct {
	Stx        context.Context
	InitConfig *config.InitConfig
}

func NewClient(ctx context.Context, appId, appKey string, options ...config.InitOption) *Client {
	var initConfig *config.InitConfig
	initConfig = config.PrdConfig()
	initConfig.AppInfo = &config.AppInfoConfig{
		AppId:  appId,
		AppKey: appKey,
	}
	for _, o := range options {
		o.Apply(initConfig)
	}
	return &Client{Stx: ctx, InitConfig: initConfig}
}




/**
 * @desc 获取用户登录信息
 * @note 获取用户信息通过服务端接口用 ticket 获取。 OpenID 与 UnionID 的定义及使用建议，见本页最下方 OpenID 与 UnionID
 * @km https://dev.iggrowth.cn/sdk/docs/index.html#/server/api_account?id=%e8%b4%a6%e5%8f%b7%e6%8e%a5%e5%85%a5
 */
func (c *Client) GetUserTicket(req *entity.UserTicketReq) (resp *entity.UserTicketResp, err error ){

	// 参数检查
	if err := req.IsValid(); err != nil {
		return nil, err
	}

	timeStr := strconv.FormatInt(util.TodayUnix(),10)
	var data = url.Values{
		"app_id":[]string{c.InitConfig.AppInfo.AppId},
		"platform":[]string{req.Platform},
		"ticket":[]string{req.Ticket},
		"time":[]string{timeStr},
	}

	data.Add("sign",util.GenerateSign(c.InitConfig.AppInfo.AppKey,data))

	urlStr := fmt.Sprintf("%s/x/open/user/ticket", c.InitConfig.RequestUrl)

	resp = &entity.UserTicketResp{}


	//重试&熔断器算法
	result,err := c.RetryBreakerGet(urlStr,resp,data,"openUserTicket")
	if err != nil{
		return nil,err
	}

	return result.(*entity.UserTicketResp),nil

}

/**
 * @desc 设备信息查询【判断该用户是否是负毛利用户】
 * @note 设备信息一般当天内不会改变，建议游戏服务端在查询后做好一定时间的缓存，不用过于频繁地请求
 * @km https://dev.iggrowth.cn/sdk/docs/index.html#/server/api_device_info
 */
func (c *Client) QueryDeviceInfo(req *entity.DeviceInfoReq) (resp *entity.DeviceInfoResp, err error ) {
	// 参数检查
	if err := req.IsValid(); err != nil {
		return nil, err
	}

	timeStr := strconv.FormatInt(util.TodayUnix(),10)
	var data = url.Values{
		"app_id":[]string{c.InitConfig.AppInfo.AppId},
		"tuid":[]string{req.Tuid},
		"timestamp":[]string{timeStr},
	}

	data.Add("sign",util.GenerateSign(c.InitConfig.AppInfo.AppKey,data))
	urlStr := fmt.Sprintf("%s/x/open/deviceinfo/query", c.InitConfig.RequestUrl)

	resp = &entity.DeviceInfoResp{}


	//重试&熔断器算法
	result,err := c.RetryBreakerGet(urlStr,resp,data,"queryDeviceInfo")
	if err != nil{
		return nil,err
	}

	return result.(*entity.DeviceInfoResp),nil
}

/**
 * 重试+熔断器策略
 */
func (c *Client) RetryBreakerGet(urlStr string,resp interface{},data url.Values,uniqueName string) (interface{},error)  {
	client := c.InitConfig.GetIaaClient()

	return  clients.Retry(c.Stx,uniqueName,0,c.InitConfig.Request, func() (i interface{}, err error) {
		err = util.HttpGet(client,urlStr,data, map[string]string{},resp)

		if err != nil{
			return nil,err
		}

		return resp,nil
	},"")


}

/**
  * post 重试+熔断策略
 */
func (c *Client) RetryBreakerPost(urlStr string,data url.Values,uniqueName string,resp interface{}) (interface{},error) {
	client := c.InitConfig.GetIaaClient()

	return  clients.Retry(c.Stx,uniqueName,0,c.InitConfig.Request,
		func() (i interface{}, err error) {
			err = util.HttpPost(client,urlStr,data, map[string]string{"Content-Type": "application/x-www-form-urlencoded"},resp)

			if err != nil{
				return nil,err
			}

			return resp,nil
		},"")
}
