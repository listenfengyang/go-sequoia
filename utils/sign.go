package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func Sign(params map[string]string, key string) (string, error) {
	var signStr string
	if orderid, ok := params["order_id"]; ok {
		signStr = fmt.Sprintf("%s%s", orderid, key)
	} else if id, ok := params["id"]; ok {
		signStr = fmt.Sprintf("%s%s", id, key)
	}

	fmt.Printf("[rawString]%s\n", signStr)

	hash := md5.Sum([]byte(signStr))
	signResult := hex.EncodeToString(hash[:])

	fmt.Printf("[rawString]%s\n", signResult)
	return signResult, nil
}

func Verify(signature string, params map[string]string, signKey string) (bool, error) {
	// Generate current signature
	currentSignature, err := Sign(params, signKey)
	if err != nil {
		return false, fmt.Errorf("signature generation failed: %w", err)
	}

	// Compare signatures
	return signature == currentSignature, nil
}

// 入金&出金回调-成功-验签
func VerifyCallback(sign, payloadJson, signKey string) bool {
	fmt.Printf("payload json: %s signKey: %s\n", payloadJson, signKey)
	mac := hmac.New(sha256.New, []byte(signKey))
	if _, err := mac.Write([]byte(payloadJson)); err != nil {
		return false
	}

	fmt.Printf("verify sign: %s\n", hex.EncodeToString(mac.Sum(nil)))
	fmt.Printf("header sign: %s\n", sign)
	return hex.EncodeToString(mac.Sum(nil)) == sign
}

func SignCallbackJSONRaw(payloadJson string, signKey string) (string, error) {
	mac := hmac.New(sha256.New, []byte(signKey))
	if _, err := mac.Write([]byte(payloadJson)); err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}

func VerifyCallbackJSON(payload interface{}, signature string, signKey string) bool {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("[VerifyCallbackJSON json.Marshal error]%v\n", err)
		return false
	}

	fmt.Printf("[payloadJson]%s\n", string(payloadBytes))

	mac := hmac.New(sha256.New, []byte(signKey))
	if _, err := mac.Write(payloadBytes); err != nil {
		fmt.Printf("[VerifyCallbackJSON mac.Write error]%v\n", err)
		return false
	}
	expected := hex.EncodeToString(mac.Sum(nil))

	fmt.Printf("[expectedSignature]%s\n", expected)
	fmt.Printf("[receivedSignature]%s\n", signature)

	return hmac.Equal([]byte(expected), []byte(signature))
}
