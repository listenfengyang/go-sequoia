package go_sequoia

import (
	"errors"

	"github.com/listenfengyang/go-sequoia/utils"
)

// 出金-成功回调
func (cli *Client) WithdrawCallback(sign, signKey string, payload string, req SequoiaWithdrawCallbackReq, processor func(req SequoiaWithdrawCallbackReq) error) error {

	// Verify signature
	flag := utils.VerifyCallback(sign, payload, signKey)
	if !flag {
		//签名校验失败
		cli.logger.Errorf("sequoia withdraw back verify fail, req: %s", payload)
		return errors.New("sign verify error")
	}

	//开始处理
	return processor(req)
}
