package go_sequoia

type SequoiaInitParams struct {
	MerchantInfo `yaml:",inline" mapstructure:",squash"`

	DepositUrl        string `json:"depositUrl" mapstructure:"depositUrl" config:"depositUrl"  yaml:"depositUrl"`
	WithdrawUrl       string `json:"withdrawUrl" mapstructure:"withdrawUrl" config:"withdrawUrl"  yaml:"withdrawUrl"`
	DepositNotifyUrl  string `json:"depositNotifyUrl" mapstructure:"depositNotifyUrl" config:"depositNotifyUrl"  yaml:"depositNotifyUrl"`     // 回调地址
	WithdrawNotifyUrl string `json:"withdrawNotifyUrl" mapstructure:"withdrawNotifyUrl" config:"withdrawNotifyUrl"  yaml:"withdrawNotifyUrl"` // 回调地址
	ReturnUrl         string `json:"returnUrl" mapstructure:"returnUrl" config:"returnUrl"  yaml:"returnUrl"`                                 // 回调地址
}

type MerchantInfo struct {
	MerchantIdTJS    string `json:"merchantIdTJS" mapstructure:"merchantIdTJS" config:"merchantIdTJS"  yaml:"merchantIdTJS"`             // 商户id
	MerchantIdKGS    string `json:"merchantIdKGS" mapstructure:"merchantIdKGS" config:"merchantIdKGS"  yaml:"merchantIdKGS"`             // 商户id
	MerchantIdUZS    string `json:"merchantIdUZS" mapstructure:"merchantIdUZS" config:"merchantIdUZS"  yaml:"merchantIdUZS"`             // 商户id
	MerchantIdKZT    string `json:"merchantIdKZT" mapstructure:"merchantIdKZT" config:"merchantIdKZT"  yaml:"merchantIdKZT"`             // 商户id
	MerchantIdAZN    string `json:"merchantIdAZN" mapstructure:"merchantIdAZN" config:"merchantIdAZN"  yaml:"merchantIdAZN"`             // 商户id
	SecretKeyTJS     string `json:"secretKeyTJS" mapstructure:"secretKeyTJS" config:"secretKeyTJS"  yaml:"secretKeyTJS"`                 // 商户密钥
	SecretKeyKGS     string `json:"secretKeyKGS" mapstructure:"secretKeyKGS" config:"secretKeyKGS"  yaml:"secretKeyKGS"`                 // 商户密钥
	SecretKeyUZS     string `json:"secretKeyUZS" mapstructure:"secretKeyUZS" config:"secretKeyUZS"  yaml:"secretKeyUZS"`                 // 商户密钥
	SecretKeyKZT     string `json:"secretKeyKZT" mapstructure:"secretKeyKZT" config:"secretKeyKZT"  yaml:"secretKeyKZT"`                 // 商户密钥
	SecretKeyAZN     string `json:"secretKeyAZN" mapstructure:"secretKeyAZN" config:"secretKeyAZN"  yaml:"secretKeyAZN"`                 // 商户密钥
	WebhookSecretTJS string `json:"webhookSecretTJS" mapstructure:"webhookSecretTJS" config:"webhookSecretTJS"  yaml:"webhookSecretTJS"` // 商户密钥
	WebhookSecretKGS string `json:"webhookSecretKGS" mapstructure:"webhookSecretKGS" config:"webhookSecretKGS"  yaml:"webhookSecretKGS"` // 商户密钥
	WebhookSecretUZS string `json:"webhookSecretUZS" mapstructure:"webhookSecretUZS" config:"webhookSecretUZS"  yaml:"webhookSecretUZS"` // 商户密钥
	WebhookSecretKZT string `json:"webhookSecretKZT" mapstructure:"webhookSecretKZT" config:"webhookSecretKZT"  yaml:"webhookSecretKZT"` // 商户密钥
	WebhookSecretAZN string `json:"webhookSecretAZN" mapstructure:"webhookSecretAZN" config:"webhookSecretAZN"  yaml:"webhookSecretAZN"` // 商户密钥
}

//============================================================

// Sequoia入金
type SequoiaDepositReq struct {
	OrderId                  string `json:"order_id" mapstructure:"order_id" config:"order_id"  yaml:"order_id"`                                                                                 // 订单号
	Amount                   string `json:"amount" mapstructure:"amount" config:"amount"  yaml:"amount"`                                                                                         // 金额
	Token                    string `json:"token" mapstructure:"token" config:"token"  yaml:"token"`                                                                                             // 令牌
	PaymentMethod            string `json:"payment_method" mapstructure:"payment_method" config:"payment_method"  yaml:"payment_method"`                                                         // 支付方式
	Currency                 string `json:"currency" mapstructure:"currency" config:"currency"  yaml:"currency"`                                                                                 // 币种
	BackToMerchantSuccessUrl string `json:"back_to_merchant_success_url" mapstructure:"back_to_merchant_success_url" config:"back_to_merchant_success_url"  yaml:"back_to_merchant_success_url"` // 回调地址
	BackToMerchantUrl        string `json:"back_to_merchant_url" mapstructure:"back_to_merchant_url" config:"back_to_merchant_url"  yaml:"back_to_merchant_url"`                                 // 回调地址
	MerchantUserId           string `json:"merchant_user_id" mapstructure:"merchant_user_id" config:"merchant_user_id"  yaml:"merchant_user_id"`                                                 // 商户用户id
	MerchantUserIp           string `json:"merchant_user_ip" mapstructure:"merchant_user_ip" config:"merchant_user_ip"  yaml:"merchant_user_ip"`                                                 // 商户用户ip
	Date                     string `json:"date" mapstructure:"date" config:"date"  yaml:"date"`                                                                                                 // 日期
	WalletProvider           string `json:"wallet_provider" mapstructure:"wallet_provider" config:"wallet_provider"  yaml:"wallet_provider"`
	CardNumber               string `json:"card_number" mapstructure:"card_number" config:"card_number"  yaml:"card_number"`
	SenderName               string `json:"sender_name" mapstructure:"sender_name" config:"sender_name"  yaml:"sender_name"`
	Email                    string `json:"email" mapstructure:"email" config:"email"  yaml:"email"`

	// MerchantId               string  `json:"merchantId" mapstructure:"merchantId" config:"merchantId"  yaml:"merchantId"`                                                         // 商户id
	// CallbackUrl string `json:"callback_url" mapstructure:"callback_url" config:"callback_url"  yaml:"callback_url"` // 回调地址
}

type SequoiaDepositRsp struct {
	Status  string         `json:"status" mapstructure:"status"`
	Message string         `json:"message" mapstructure:"message"`
	Code    string         `json:"code" mapstructure:"code"`
	Data    DepositRspData `json:"data" mapstructure:"data"`
}

type DepositRspData struct {
	RedirectUrl string `json:"redirect_url" mapstructure:"redirect_url"` // 重定向url
}

// Sequoia出金
type SequoiaWithdrawReq struct {
	Id             string `json:"id" mapstructure:"id"`                           // 出金订单id
	Date           string `json:"date" mapstructure:"date"`                       // 日期
	CardNumber     string `json:"card_number" mapstructure:"card_number"`         // 银行账户号
	Amount         string `json:"amount" mapstructure:"amount"`                   // 金额
	Currency       string `json:"currency" mapstructure:"currency"`               // 币种
	CallbackUrl    string `json:"callback_url" mapstructure:"callback_url"`       // 回调地址
	WalletProvider string `json:"wallet_provider" mapstructure:"wallet_provider"` // 钱包供应商
	MerchantId     string `json:"merchant_id" mapstructure:"merchant_id"`         // 商户id
	Token          string `json:"token" mapstructure:"token"`                     // 令牌
	PayOutMethod   string `json:"pay_out_method" mapstructure:"pay_out_method"`   // 出金方式
}

type SequoiaWithdrawRsp struct {
	Status  string       `json:"status" mapstructure:"status"`
	Message string       `json:"message" mapstructure:"message"`
	Code    string       `json:"code" mapstructure:"code"`
	Data    WithdrawData `json:"data" mapstructure:"data"`
}

type WithdrawData struct {
	InternalId int64 `json:"internal_id" mapstructure:"internal_id"` // 内部订单号
}

// 入金回调
type SequoiaDepositCallbackReq struct {
	OrderId        string `json:"order_id" form:"order_id" mapstructure:"order_id"`                         // 订单号
	Date           string `json:"date" form:"date" mapstructure:"date"`                                     // 日期
	Amount         string `json:"amount" form:"amount" mapstructure:"amount"`                               // 金额（字符串形式，例如"300"）
	PaymentType    int32  `json:"payment_type" form:"payment_type" mapstructure:"payment_type"`             // 支付方式
	Status         string `json:"status" form:"status" mapstructure:"status"`                               // 订单状态：success, expired
	Currency       string `json:"currency" form:"currency" mapstructure:"currency"`                         // 币种
	ShowIntruction *bool  `json:"show_instruction" form:"show_instruction" mapstructure:"show_instruction"` // 是否展示_instructions，允许为 null
	IsRepayment    bool   `json:"is_repayment" form:"is_repayment" mapstructure:"is_repayment"`             // 是否为还款
}

type DepositCallbackData struct {
	Amount         string `json:"amount" form:"amount" mapstructure:"amount"`                               // 金额（字符串形式，例如"300"）
	CardNumber     string `json:"card_number" form:"card_number" mapstructure:"card_number"`                // 银行账户号
	Currency       string `json:"currency" form:"currency" mapstructure:"currency"`                         // 币种
	Date           string `json:"date" form:"date" mapstructure:"date"`                                     // 日期
	OrderId        string `json:"order_id" form:"order_id" mapstructure:"order_id"`                         // 订单号
	PaymentType    int32  `json:"payment_type" form:"payment_type" mapstructure:"payment_type"`             // 支付方式
	Status         string `json:"status" form:"status" mapstructure:"status"`                               // 订单状态：success, expired
	ShowIntruction *bool  `json:"show_instruction" form:"show_instruction" mapstructure:"show_instruction"` // 是否展示_instructions，允许为 null
	// success callback with changed sum 独有的参数 =》新金额（字符串形式）
	NewAmount string `json:"new_amount" form:"new_amount" mapstructure:"new_amount"`
}

// 出金回调
type SequoiaWithdrawCallbackReq struct {
	Amount     float64 `json:"amount" form:"amount" mapstructure:"amount"`                // 金额（小数点后取2位）
	CardNumber string  `json:"card_number" form:"card_number" mapstructure:"card_number"` // 银行账户号
	Currency   string  `json:"currency" form:"currency" mapstructure:"currency"`          // 币种
	Date       string  `json:"date" form:"date" mapstructure:"date"`                      // 日期
	Id         string  `json:"id" form:"id" mapstructure:"id"`                            // 出金订单id
	Status     string  `json:"status" form:"status" mapstructure:"status"`                // 订单状态：success, fail, pending
}

type WithdrawCallbackData struct {
	Amount     float64 `json:"amount" form:"amount" mapstructure:"amount"`                // 金额（小数点后取2位）
	CardNumber string  `json:"card_number" form:"card_number" mapstructure:"card_number"` // 银行账户号
	Currency   string  `json:"currency" form:"currency" mapstructure:"currency"`          // 币种
	Date       string  `json:"date" form:"date" mapstructure:"date"`                      // 日期
	Id         string  `json:"id" form:"id" mapstructure:"id"`                            // 出金订单id
	Status     string  `json:"status" form:"status" mapstructure:"status"`                // 订单状态：success, fail, pending
}
