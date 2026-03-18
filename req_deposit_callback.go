package go_sequoia

import (
	"encoding/json"
	"errors"

	"github.com/listenfengyang/go-sequoia/utils"
)

// 充值-成功回调（原始版本，沿用 MD5 验签）
func (cli *Client) DepositCallback(headerSign string, req SequoiaDepositCallbackReq, processor func(SequoiaDepositCallbackReq) error) error {
	//验证签名
	data, _ := json.Marshal(req)
	payload := string(data)
	flag := utils.VerifyCallback(headerSign, payload, cli.Params.MerchantInfo.WebhookSecretTJS)
	if !flag {
		//签名校验失败
		reqJson, _ := json.Marshal(req)
		cli.logger.Errorf("sequoia deposit back verify fail, req: %s", string(reqJson))
		return errors.New("sign verify error")
	}

	//开始处理
	return processor(req)
}
