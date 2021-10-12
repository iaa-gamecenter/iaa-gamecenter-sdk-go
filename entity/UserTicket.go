package entity

import "errors"

type (
	UserTicketReq struct {
		Platform string  //游戏app登录返回信息中获取platform字段，不要对该字段做任何校验，直接透传即可
		Ticket   string //获取用户信息临时标识，24小时有效(游戏app登录返回信息中获取)
	}

	UserTicketResp struct {
		Code        int    `json:"code"`
		Message     string `json:"message"`
		ShowErr     int    `json:"showErr"`
		CurrentTime int    `json:"currentTime"`
		Data        struct {
			OpenID       string `json:"open_id"`
			Nickname     string `json:"nickname"`
			Avatar       string `json:"avatar"`
			UnionID      string `json:"union_id"`
			WlxPlatform  string `json:"wlx_platform"`
			Ext          string `json:"ext,omitempty"`
			Market       string `json:"market,omitempty"`
			MarketOpenId string `json:"market_open_id,omitempty"`
		} `json:"data"`
	}
)

func (ut *UserTicketReq) IsValid() error {
	if ut.Ticket == ""{
		return errors.New("ticket必填")
	}

	if ut.Platform == ""{
		return errors.New("platform必填")
	}

    return nil
}