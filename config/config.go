package config

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

var (
	//client map
	clientMap sync.Map

	iaaClient = "iaa_client"

)

type (
	InitConfig struct {
		RequestUrl string
		AppInfo *AppInfoConfig
		Request *RequestConfig
	}

	//请求方法配置信息
	RequestConfig struct {
		RetryCount     uint8          //请求异常时候，重试的次数，默认为1
		RequestTimeout uint64         //请求的超时时间，单位ms，默认为500ms
		WaitTime       uint64         //在下次重试之前等待的时间，单位ms，默认为5ms
		RetryCodes     []string       //除网络错误外和业务Code为-1外，其他业务编码下也要重试
		Breaker        *BreakerConfig //请求特定接口时候，制定特定熔断器配置，如果不配置，以全局的为准。
	}
	// 熔断器配置信息
	BreakerConfig struct {
		MaxRequests    uint32  //当熔断器为half-open时候，容许的最大通过数。如果为0默认为1。
		Interval       uint64  //在熔断器关闭状态时，多久清除一次。为0则不清除。单位为s，默认60s
		BreakerTimeout uint64  //在熔断器为open时候，多久可以变成half-open，如果为0默认为60s
		TripRequest    uint32  //熔断器的熔断条件中，当前区间统计的最低请求数量，默认为10。
		FailureRatio   float64 //熔断器的熔断条件中，当前区间统计的最低错误比率，默认为0.6。
	}
	AppInfoConfig struct {
		AppId  string
		AppKey string
	}
)

func PrdConfig() *InitConfig {
	return &InitConfig{
		RequestUrl: "https://api.iggrowth.cn",
		Request: DefaultRequestConfig(nil),
	}
}

type InitOption interface {
	Apply(*InitConfig)
}

func DefaultRequestConfig(reqConfig *RequestConfig) *RequestConfig {
	var retryCount uint8 = 1
	var requestTimeout uint64 = 500
	var waitTime uint64 = 5
	var breaker = BreakerConfig{Interval: 60, TripRequest: 10, FailureRatio: 0.6}
	if reqConfig != nil {
		retryCount = reqConfig.RetryCount
		requestTimeout = reqConfig.RequestTimeout
		waitTime = reqConfig.WaitTime
		breaker = *reqConfig.Breaker
	}
	return &RequestConfig{RetryCount: retryCount, RequestTimeout: requestTimeout, WaitTime: waitTime, Breaker: &breaker}
}

func (b *BreakerConfig) GetKey() string {
	return fmt.Sprintf("%d_%d_%d_%d_%d", b.BreakerTimeout, b.Interval, b.MaxRequests, b.TripRequest, b.FailureRatio)
}

func (c *InitConfig) GetIaaClient() *http.Client {
	v, ok := clientMap.Load(iaaClient)
	if v == nil || !ok {
		client := &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   1000 * time.Millisecond,
					KeepAlive: 60 * time.Second,
					DualStack: true,
				}).DialContext,
				MaxIdleConns:          1500,
				MaxIdleConnsPerHost:   500,
				IdleConnTimeout:       30 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
			Timeout: time.Duration(c.Request.RequestTimeout) * time.Millisecond,
		}
		clientMap.Store(iaaClient, client)
		return client
	}

	return v.(*http.Client)
}
