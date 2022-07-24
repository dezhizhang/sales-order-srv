package model

type Order struct {
	User int32 `gorm:"type:int;index"`
}
