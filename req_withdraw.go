package go_sequoia

import (
	"crypto/tls"
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/listenfengyang/go-sequoia/utils"
	"github.com/mitchellh/mapstructure"
)

func (cli *Client) WithdrawReq(req SequoiaWithdrawReq) (*SequoiaWithdrawRsp, error) {

	var params map[string]string
	mapstructure.Decode(req, &params)
	params["date"] = time.Now().Format(time.DateTime)
	params["callback_url"] = cli.Params.WithdrawNotifyUrl

	if req.Currency == "TJS" {

	} else if req.Currency == "KZT" {
		params["wallet_provider"] = "Kaspi Bank"
	} else if req.Currency == "UZS" {
		params["pay_out_method"] = "UZCARD"
	} else if req.Currency == "KGS" {

	}

	merchantId, err := GetMerchantId(req.Currency, *cli.Params)
	if err != nil {
		return nil, err
	}
	params["merchant_id"] = merchantId

	// Generate signature
	key, err := GetSecretKey(req.Currency, *cli.Params)
	if err != nil {
		return nil, err
	}
	signStr, _ := utils.Sign(params, key)
	params["token"] = signStr
	var result SequoiaWithdrawRsp
	fmt.Println(params)

	rawURL := cli.Params.WithdrawUrl
	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(params).
		SetHeaders(getHeaders()).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp2))
	cli.logger.Infof("PSPResty#sequoia#withdraw->%s", string(restLog))

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
