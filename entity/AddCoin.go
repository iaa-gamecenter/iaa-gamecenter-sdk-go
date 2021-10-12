package entity

import "errors"

type (
	AddCoinReq struct {
		OpenId string
		CoinNum int
		TradeNo string
	}
	AddCoinResp struct {
		Code        int    `json:"code"`
		Message     string `json:"message"`
		ShowErr     int    `json:"showErr"`
		CurrentTime int    `json:"currentTime"`
		Data        struct {
		} `json:"data"`
	}
)

func (ut *AddCoinReq) IsValid() error {
	if ut.OpenId == ""{
		return errors.New("open_id必填")
	}

	if ut.CoinNum == 0{
		return errors.New("coin_num必填")
	}

	if ut.TradeNo == ""{
		return errors.New("trade_no必填")
	}

	return nil
}
