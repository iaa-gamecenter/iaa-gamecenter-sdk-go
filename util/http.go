package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// 发放 http go 请求并解析返回体
func HttpGet(client *http.Client, reqUrl string, values url.Values, headers map[string]string, result interface{}) error {

	reqUrl = fmt.Sprintf("%s?%s", reqUrl, values.Encode())
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return err
	}

	for key := range headers {
		req.Header.Set(key, headers[key])
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := ioutil.ReadAll(resp.Body)
	return json.Unmarshal(body, &result)
}

// 发出 http post 请求并解析返回体
func HttpPost(client *http.Client, reqUrl string, values url.Values, headers map[string]string, result interface{}) error {
	req, err := http.NewRequest(http.MethodPost, reqUrl, strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}

	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for key := range headers {
		req.Header.Set(key, headers[key])
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := ioutil.ReadAll(resp.Body)
	return json.Unmarshal(body, &result)
}
