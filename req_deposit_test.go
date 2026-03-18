package go_sequoia

import (
	"testing"
	"time"
)

func TestDeposit(t *testing.T) {

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
		DepositUrl:       DEPOSIT_URL,
		DepositNotifyUrl: DEPOSIT_NOTIFY_URL,
		ReturnUrl:        RETURN_URL,
	})
	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		cli.logger.Errorf("err:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

// TJS, UZS, KGS, AZN, KZT
func GenDepositRequestDemo() SequoiaDepositReq {
	return SequoiaDepositReq{
		OrderId:        "2026023532263416",
		Amount:         "400",
		PaymentMethod:  "1",
		Currency:       "AZN",
		MerchantUserId: "321421",
		MerchantUserIp: "35.76.173.139",
		Date:           time.Now().Format(time.DateTime),
		WalletProvider: "Kaspi Bank", // DUSHANBE_CITY, UZCARD, 无, 无, Kaspi Bank
		SenderName:     "jane",
		CardNumber:     "2345678901234567", // 16dg
		Email:          "2535@gmail.com",
	}
}
