package go_sequoia

import (
	"testing"
)

func TestWithdraw(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &SequoiaInitParams{
		MerchantInfo: MerchantInfo{
			MerchantIdTJS: MERCHANT_ID_TJS, SecretKeyTJS: SECRET_KEY_TJS,
			MerchantIdUZS: MERCHANT_ID_UZS, SecretKeyUZS: SECRET_KEY_UZS,
			MerchantIdKGS: MERCHANT_ID_KGS, SecretKeyKGS: SECRET_KEY_KGS,
			MerchantIdAZN: MERCHANT_ID_AZN, SecretKeyAZN: SECRET_KEY_AZN,
			MerchantIdKZT: MERCHANT_ID_KZT, SecretKeyKZT: SECRET_KEY_KZT,
		},
		WithdrawUrl:       WITHDRAW_URL,
		WithdrawNotifyUrl: WITHDRAW_NOTIFY_URL,
	})
	//发请求
	resp, err := cli.WithdrawReq(GenWithdrawRequestDemo())
	if err != nil {
		cli.logger.Errorf("err:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

// TJS, UZS, KGS, AZN, KZT
func GenWithdrawRequestDemo() SequoiaWithdrawReq {
	return SequoiaWithdrawReq{
		Id:         "202602351322634130",
		CardNumber: "2345678901234567",
		Amount:     "10",
		Currency:   "TJS",
	}
}
