package go_sequoia

import (
	"fmt"
)

func GetSecretKey(currency string, params SequoiaInitParams) (string, error) {
	switch currency {
	case "TJS":
		return params.SecretKeyTJS, nil
	case "KGS":
		return params.SecretKeyKGS, nil
	case "UZS":
		return params.SecretKeyUZS, nil
	case "KZT":
		return params.SecretKeyKZT, nil
	case "AZN":
		return params.SecretKeyAZN, nil
	default:
		return "", fmt.Errorf("currency not support")
	}
}

func GetMerchantId(currency string, params SequoiaInitParams) (string, error) {
	switch currency {
	case "TJS":
		return params.MerchantIdTJS, nil
	case "KGS":
		return params.MerchantIdKGS, nil
	case "UZS":
		return params.MerchantIdUZS, nil
	case "KZT":
		return params.MerchantIdKZT, nil
	case "AZN":
		return params.MerchantIdAZN, nil
	default:
		return "", fmt.Errorf("currency not support")
	}
}
