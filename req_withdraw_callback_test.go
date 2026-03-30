package go_sequoia

import (
	"testing"
)

func TestWithdrawCallback(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &SequoiaInitParams{
		MerchantInfo:      MerchantInfo{MerchantIdTJS: MERCHANT_ID_TJS, SecretKeyTJS: SECRET_KEY_TJS, WebhookSecretTJS: WEBHOOK_SECRET_TJS},
		WithdrawUrl:       WITHDRAW_URL,
		WithdrawNotifyUrl: WITHDRAW_NOTIFY_URL,
	})

	headerSign := "6e68bfdcb2ba59443cc47f55da9424ea76d461764b9c5a087b74690acba6805c"

	payloadJson, req := GenWdRequestDemo()

	err := cli.WithdrawCallback(headerSign, cli.Params.MerchantInfo.WebhookSecretTJS, payloadJson, req, func(req SequoiaWithdrawCallbackReq) error { return nil })
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}
}

// {\"id\":\"202602351322634130\",\"date\":\"16.03.2026 10:57\",\"amount\":100,\"card_number\":\"2345********4567\",\"status\":\"success\",\"currency\":\"TJS\"}
func GenWdRequestDemo() (payloadJson string, req SequoiaWithdrawCallbackReq) {
	return "{\"id\":\"202602351322634130\",\"date\":\"16.03.2026 10:57\",\"amount\":100,\"card_number\":\"2345********4567\",\"status\":\"success\",\"currency\":\"TJS\"}",
		SequoiaWithdrawCallbackReq{
			Id:         "202602351322634130",
			Date:       "16.03.2026 10:57",
			Amount:     100,
			CardNumber: "2345678901234567",
			Status:     "success",
			Currency:   "TJS",
		}
}
