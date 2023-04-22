package model

type Cart struct {
	BaseModel
	User    int32 `gorm:"type:int;index;"`
	Product int32 `gorm:"type:int;index"`
	Nums    int32 `gorm:"type"`
	Checked bool  `json:"checked"`
}
