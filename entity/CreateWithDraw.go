package entity

import "errors"

type (
	//remark非必传
	CreateWithDrawReq struct {
		OpenId    string //用户在当前项目内的唯一标示
		Amount    int    //打款金额，单位：分
		ShipDelay int    //延迟打款时间，单位秒，至少传1，否则会报参数错误
		NotifyUrl string //回调url，具体解释在下面
		Tk        string //sdk获取的tk值（勿缓存）
		Tuid      string //sdk获取的tuid值（勿缓存）
		TpWay     string //提现方式，20000:微信 20001:支付宝
		CpOrderNo string //游戏方的订单id，需要保证唯一
		Remark    string //游戏方需要回传的信息
	}
	CreateWithDrawResp struct {
		Code        int    `json:"code"`
		Message     string `json:"message"`
		ShowErr     int    `json:"showErr"`
		CurrentTime int    `json:"currentTime"`
		Data        struct {
			CpOrderNo string `json:"cp_order_no"`
			OrderNo   string `json:"order_no"`
			Status    string `json:"status"`
		} `json:"data"`
	}
)

func (ut *CreateWithDrawReq) IsValid() error {
	if ut.OpenId == "" {
		return errors.New("open_id必填")
	}

	if ut.Amount == 0 {
		return errors.New("amount必填")
	}

	if ut.ShipDelay == 0 {
		return errors.New("ship_delay必填")
	}

	if ut.NotifyUrl == "" {
		return errors.New("notify_url必填")
	}

	if ut.Tk == "" {
		return errors.New("tk必填")
	}

	if ut.Tuid == "" {
		return errors.New("tuid必填")
	}

	if ut.TpWay == "" {
		return errors.New("tp_way必填")
	}

	if ut.CpOrderNo == "" {
		return errors.New("cp_order_no必填")
	}

	return nil
}
