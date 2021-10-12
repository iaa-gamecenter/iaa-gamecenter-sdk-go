package iaa_gamecenter_sdk_go

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/iaa-gamecenter/iaa-gamecenter-sdk-go/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_CreateWithDraw(t *testing.T) {
	var ctx context.Context
	client := NewClient(ctx,"dsada","dsadasd")


	var req = &entity.CreateWithDrawReq{
		OpenId:    "u11YfeuCWacpBqXkQ9aZD",
		Amount:    10,
		ShipDelay: 1,
		NotifyUrl: "http://www.test.net",
		Tk:        "ACEVm-ndIO2MRM9JcGZGLWcyXx9NZf7baz5pZ3dkZHl5",
		Tuid:      "FZvp3SDtjETPSXBmRi1nMg",
		TpWay:     "20000",
		CpOrderNo: "20201129-1527534111-FZvp3SDtjETPSXBmRi1nMg1",
		Remark:    "test",
	}

	resp,err := client.CreateWithDraw(req)
	str,_ := json.Marshal(resp)
	fmt.Print(string(str))
	fmt.Println(resp.Data.CpOrderNo,resp.Data.OrderNo)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestClient_ShowWithDraw(t *testing.T) {
	var ctx context.Context
	client := NewClient(ctx,"dsada","dsadasd")

	var req = &entity.ShowWithDrawReq{OpenId:"u11YfBiAp2DPJHHAC5JMd"}
	resp,err := client.ShowWithDraw(req)
	str,_ := json.Marshal(resp)
	fmt.Println(string(str))
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestClient_QueryWithDraw(t *testing.T) {
	var ctx context.Context
	client := NewClient(ctx,"dsada","dsadasd")

	var req = &entity.QueryWithDrawReq{
		OpenId:    "u11YfBiAp2DPJHHAC5JMd",
		OrderNo:   "2i3gj2oig23g2g",
		CpOrderNo: "20201127-185953601-11Y9vxVodV4Uprs7THQT",
	}
	
	resp,err := client.QueryWithDraw(req)
	str,_ := json.Marshal(resp)
	fmt.Println(string(str))
	assert.Nil(t, err)
	assert.NotNil(t, resp)

}