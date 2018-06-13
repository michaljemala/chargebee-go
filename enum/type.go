package enum

type Type string

const (
	TypeCard                  Type = "card"
	TypePaypalExpressCheckout Type = "paypal_express_checkout"
	TypeAmazonPayments        Type = "amazon_payments"
	TypeDirectDebit           Type = "direct_debit"
	TypeGeneric               Type = "generic"
	TypeAlipay                Type = "alipay"
	TypeUnionpay              Type = "unionpay"
	TypeApplePay              Type = "apple_pay"
	TypeWechatPay             Type = "wechat_pay"
)
