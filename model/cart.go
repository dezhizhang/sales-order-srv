package model

type Cart struct {
	BaseModel
	UserId    int32 `gorm:"type:int;index" json:"userId"`
	ProductId int32 `gorm:"type:int;index" json:"productId"`
	Nums      int32 `gorm:"type:int" json:"nums"`
	Checked   bool  `json:"checked"`
}
