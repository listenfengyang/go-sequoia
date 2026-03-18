package go_sequoia

import (
	"crypto/tls"
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/listenfengyang/go-sequoia/utils"
	"github.com/mitchellh/mapstructure"
)

// 下单
func (cli *Client) Deposit(req SequoiaDepositReq) (*SequoiaDepositRsp, error) {

	rawURL := cli.Params.DepositUrl

	var params map[string]string
	mapstructure.Decode(req, &params)
	params["callback_url"] = cli.Params.DepositNotifyUrl
	params["back_to_merchant_success_url"] = cli.Params.ReturnUrl
	params["back_to_merchant_url"] = cli.Params.ReturnUrl

	if req.Currency == "TJS" {
		delete(params, "card_number")
	} else if req.Currency != "KZT" {
		delete(params, "send_name")
		delete(params, "email")
	} else if req.Currency == "UZS" || req.Currency == "KGS" {
		delete(params, "wallet_provider")
	}

	// Generate signature
	key, err := GetSecretKey(req.Currency, *cli.Params)
	if err != nil {
		return nil, err
	}
	merchantId, err := GetMerchantId(req.Currency, *cli.Params)
	if err != nil {
		return nil, err
	}
	params["merchant_id"] = merchantId

	signStr, _ := utils.Sign(params, key)
	params["token"] = signStr
	fmt.Println(params)

	var result SequoiaDepositRsp

	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetFormData(params).
		SetBody(params).
		SetHeaders(getHeaders()).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp2))
	cli.logger.Infof("PSPResty#sequoia#deposit->%s", string(restLog))

	if err != nil {
		return nil, err
	}

	if resp2.StatusCode() != 200 {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("status code: %d", resp2.StatusCode())
	}

	if resp2.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v, body:%s", resp2.Error(), resp2.Body())
	}

	return &result, nil
}
