package util

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"sort"
)

/**
 * 生成签名
 */
func GenerateSign(appKey string,params url.Values) (signStr string) {

    signStr = signature(params,appKey)

	return
}

/**
 * 签名算法
 */
func signature(values url.Values,appKey string) (r string)  {
	values.Del("sign")
	values.Add("app_key", appKey)
	keys := make([]string, 0, len(values))
	for k, _ := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, v := range keys {
		r += v + values.Get(v)
	}
	hashed := md5.Sum([]byte(r))
	r = fmt.Sprintf("%x", hashed)
	values.Del("app_key")
	return
}