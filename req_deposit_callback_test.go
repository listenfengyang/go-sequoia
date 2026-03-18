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
		},
		DepositUrl:       DEPOSIT_URL,
		WithdrawUrl:      WITHDRAW_URL,
		DepositNotifyUrl: DEPOSIT_NOTIFY_URL,
	})

	headerSign := "aa4d0328b880f641a41a3b5736bf4cb468bac2a07bd2a5a3431839d5941a4664"

	err := cli.DepositCallback(headerSign, GenCallbackRequestDemo(), func(SequoiaDepositCallbackReq) error { return nil })
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", err)
}

func GenCallbackRequestDemo() SequoiaDepositCallbackReq {
	return SequoiaDepositCallbackReq{
		OrderId:        "2026023532263465",
		Date:           "13.03.2026 08:05",
		Amount:         "300",
		PaymentType:    1,
		Status:         "success",
		Currency:       "TJS",
		ShowIntruction: nil,
		IsRepayment:    false,
	}
}
