package go_sequoia

import (
	"testing"
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
	// amount=443.00&back_to_merchant_success_url=https%3A%2F%2Fcpta.supermarkets.com&back_to_merchant_url=https%3A%2F%2Fcpta.supermarkets.com&callback_url=https%3A%2F%2Fapi-test.logtec.dev%2Ffapi%2Fpayment%2Fpsp%2Fpublic%2Fsequoia%2Fdeposit%2Fback&card_number=&currency=AZN&date=2026-03-27+13%3A53%3A41&email=jane.y1%40yopmail.com&merchant_id=b06eaaec-c729-4666-bf78-fb21cec32d3e&merchant_user_id=860007610&merchant_user_ip=10.16.67.67&order_id=202603271353410627&payment_method=1&sender_name=%E8%B5%AB%E6%95%8F%C2%B7%E7%8F%8D%E7%8F%8D%C2%B7%E6%A0%BC%E5%85%B0%E6%9D%B0+&token=37f4a8d77acd528d849f92077862f7b1&wallet_provider=

	return SequoiaDepositReq{
		OrderId:        "20260235322634168",
		Amount:         "443.00",
		PaymentMethod:  "1",
		Currency:       "AZN",
		MerchantUserId: "321421",
		MerchantUserIp: "35.76.173.139",
		SenderName:     "jane",
		CardNumber:     "2345678901234567", // 16dg
		Email:          "2535@gmail.com",
	}
}
