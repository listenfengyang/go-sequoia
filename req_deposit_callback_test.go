package go_sequoia

import (
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func TestCallback(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &SequoiaInitParams{
		MerchantInfo: MerchantInfo{
			MerchantIdTJS:    MERCHANT_ID_TJS,
			SecretKeyTJS:     SECRET_KEY_TJS,
			WebhookSecretTJS: WEBHOOK_SECRET_TJS,
			MerchantIdAZN:    MERCHANT_ID_AZN,
			SecretKeyAZN:     SECRET_KEY_AZN,
			WebhookSecretAZN: WEBHOOK_SECRET_AZN,
		},
		DepositUrl:       DEPOSIT_URL,
		WithdrawUrl:      WITHDRAW_URL,
		DepositNotifyUrl: DEPOSIT_NOTIFY_URL,
	})

	headerSign := "e885c4d217fbf582638ceff6ad2a6189cf61e83a6dfb779a105f67b853a26801"

	err := cli.DepositCallback(headerSign, cli.Params.MerchantInfo.WebhookSecretAZN, GenCallbackRequestDemo(), func(SequoiaDepositCallbackReq) error { return nil })
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", err)
}

// {\"order_id\":\"202603300404420530\",\"date\":\"30.03.2026 07:06\",\"amount\":\"442\",\"payment_type\":1,\"status\":\"success\",\"currency\":\"AZN\",\"show_instruction\":null,\"is_repayment\":false}
func GenCallbackRequestDemo() SequoiaDepositCallbackReq {
	return SequoiaDepositCallbackReq{
		OrderId:        "202603300404420530",
		Date:           "30.03.2026 07:06",
		Amount:         "442",
		PaymentType:    1,
		Status:         "success",
		Currency:       "AZN",
		ShowIntruction: nil,
		IsRepayment:    false,
	}
}
