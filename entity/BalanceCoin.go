package entity

import "errors"

type (
	BalanceCoinReq struct {
		OpenId  string
	}
	BalanceCoinResp struct {
		Code        int    `json:"code"`
		Message     string `json:"message"`
		ShowErr     int    `json:"showErr"`
		CurrentTime int    `json:"currentTime"`
		Data        struct {
			CoinBalance int `json:"coin_balance"`
		} `json:"data"`
	}
)
func (ut *BalanceCoinReq) IsValid() error {
	if ut.OpenId == ""{
		return errors.New("open_id必填")
	}

	return nil
}