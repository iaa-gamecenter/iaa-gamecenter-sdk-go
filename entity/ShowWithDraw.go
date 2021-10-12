package entity

import "errors"

type (
	ShowWithDrawReq struct {
		OpenId string
	}
	ShowWithDrawResp struct {
		Code        int    `json:"code"`
		Message     string `json:"message"`
		ShowErr     int    `json:"showErr"`
		CurrentTime int    `json:"currentTime"`
		Data        struct {
			AppId          string   `json:"app_id"`
			DayNum         int      `json:"day_num"`
			DayAmount      int      `json:"day_amount"`
			PayMethod      []string `json:"pay_method"`
			PayTitle       string   `json:"pay_title"`
			IsBindWx       int      `json:"is_bind_wx"`
			WxNickname     string   `json:"wx_nickname"`
			IsBindAlipay   int      `json:"is_bind_alipay"`
			AlipayNickName string   `json:"alipay_nick_name"`
		} `json:"data"`
	}
)

func (ut *ShowWithDrawReq) IsValid() error {
	if ut.OpenId == "" {
		return errors.New("open_id必填")
	}

	return nil
}
