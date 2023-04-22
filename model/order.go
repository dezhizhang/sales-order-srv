package model

// 订单信息

type OrderInfo struct {
	BaseModel
	UserId     int32   `gorm:"type:int" json:"userId"`
	OrderNo    string  `gorm:"type:varchar(30)"`  // 订单号
	PayType    string  `gorm:"type:varchar(20)"`  //alipay(支付宝),wechat(微信)
	Status     string  `gorm:"type:varchar(20)"`  // 订单的状态
	TradeNo    string  `gorm:"type:varchar(100)"` // 订单交易号
	OrderMount float32 `gorm:"type:int"`          //订单的金额
	PayTime    int32   `gorm:"type:int"`          // 订单的支付时间
	Address    string  `gorm:"type:varchar(100)"` // 收件人地址
	SignName   string  `gorm:"type:varchar(20)"`  // 收件人姓名
	SignMobile string  `gorm:"type:varchar(11)"`  //收件人电话

}

type OrderProductInfo struct {
	BaseModel
	OrderId      int32   `gorm:"type:int" json:"orderId"`
	ProductId    int32   `gorm:"type:int" json:"productId"`
	ProductName  string  `gorm:"type:varchar(100)" json:"productName"`
	ProductPrice float32 `json:"productPrice"`
	Nums         int32   `gorm:"type:int" json:"nums"`
}
