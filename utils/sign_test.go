package utils

import (
	"testing"
)

func TestSign(t *testing.T) {

}

// 验证与 https://emn178.github.io/online-tools/sha256.html 生成结果一致
func TestSignCallbackJSONRaw_TJS(t *testing.T) {
	const payload = `{"order_id":"2026023532263465","date":"13.03.2026 08:05","amount":"300","payment_type":1,"status":"success","currency":"TJS","show_instruction":null,"is_repayment":false}`
	const key = ""
	const expected = ""

	signature, err := SignCallbackJSONRaw(payload, key)
	if err != nil {
		t.Fatalf("SignCallbackJSONRaw error: %v", err)
	}

	if signature != expected {
		t.Fatalf("signature mismatch, got %s, want %s", signature, expected)
	}
}
