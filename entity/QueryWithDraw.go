package entity

import "errors"

type (
	QueryWithDrawReq struct {
		OpenId    string
		OrderNo   string
		CpOrderNo string
	}
	QueryWithDrawResp struct {
		Code        int    `json:"code"`
		Message     string `json:"message"`
		ShowErr     int    `json:"showErr"`
		CurrentTime int    `json:"currentTime"`
		Data        struct {
			CpOrderNo string `json:"cp_order_no"`
			OrderNo   string `json:"order_no"`
			Status    string `json:"status"`
			Ext       string `json:"ext"`
		} `json:"data"`
	}
)

func (ut *QueryWithDrawReq) IsValid() error {
	if ut.OpenId == "" {
		return errors.New("open_id必填")
	}

	if ut.OrderNo == "" {
		return errors.New("order_no必填")
	}

	if ut.CpOrderNo == "" {
		return errors.New("cp_order_no必填")
	}

	return nil
}
