package entity

import "errors"

type (
	DeviceInfoReq struct {
		Tuid   string
	}
	DeviceInfoResp struct {
		Code        int    `json:"code"`
		Message     string `json:"message"`
		ShowErr     int    `json:"showErr"`
		CurrentTime int    `json:"currentTime"`
		Data        struct {
			GrossProfit string `json:"gross_profit"`
		} `json:"data"`
	}
)
func (ut *DeviceInfoReq) IsValid() error {
	if ut.Tuid == ""{
		return errors.New("tuid必填")
	}


	return nil
}