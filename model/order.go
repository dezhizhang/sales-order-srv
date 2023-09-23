package model

// 订单信息

type OrderInfo struct {
	BaseModel
	UserId     int32   `gorm:"type:int;index;" json:"user_id"`
	OrderNo    string  `gorm:"type:varchar(30)" json:"order_no"`    //订单号
	PayType    string  `gorm:"type:varchar(20)" json:"pay_type"`    //alipay(支付宝),wechat(微信)
	Status     string  `gorm:"type:varchar(20)" json:"status"`      //订单的状态
	TradeNo    string  `gorm:"type:varchar(100)" json:"trade_no"`   //订单交易号
	OrderMount float32 `gorm:"type:int" json:"order_mount"`         //订单的金额
	PayTime    int32   `gorm:"type:int" json:"pay_time"`            //订单的支付时间
	Address    string  `gorm:"type:varchar(100)" json:"address"`    //收件人地址
	SignName   string  `gorm:"type:varchar(20)" json:"sign_name"`   //收件人姓名
	SignMobile string  `gorm:"type:varchar(11)" json:"sign_mobile"` //收件人电话

}

// OrderProductInfo 订单商品信信息
type OrderProductInfo struct {
	BaseModel
	OrderId      int32   `gorm:"type:int" json:"order_id"`
	ProductId    int32   `gorm:"type:int" json:"product_id"`
	ProductName  string  `gorm:"type:varchar(100)" json:"product_name"`
	ProductPrice float32 `gorm:"type:int" json:"product_price"`
	Nums         int32   `gorm:"type:int" json:"nums"`
}
