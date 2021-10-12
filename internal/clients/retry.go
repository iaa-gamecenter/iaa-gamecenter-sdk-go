package clients

import (
	"context"
	"github.com/iaa-gamecenter/iaa-gamecenter-sdk-go/config"
	"github.com/iaa-gamecenter/iaa-gamecenter-sdk-go/internal/errors"
	"time"
)

func Retry(ctxTemp context.Context, uniqueName string, count uint8, reqConfig *config.RequestConfig, reqFunc func() (interface{}, error), originErrStr string) (interface{}, error) {
	iaaResp, err := reqFunc()
	if err != nil {
		//短路器打开或者重试次数大于maxCount）
		appendStr := err.Error()
		if originErrStr != "" {
			appendStr += "; " + originErrStr
		}
		if count >= reqConfig.RetryCount {
			return nil, errors.ErrOverMaxRetryCount.AppendData(appendStr)
		} else if count < reqConfig.RetryCount { //重试，重试次数小于RetryCount就自动重试
			time.Sleep(time.Duration(reqConfig.WaitTime) * time.Millisecond)
			iaaResp, err = Retry(ctxTemp, uniqueName, count+1, reqConfig, reqFunc, appendStr)
		}
	}
	return iaaResp, err
}
