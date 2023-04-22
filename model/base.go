package model

import "time"

type BaseModel struct {
	Id        string    `gorm:"primaryKey;type:varchar(200)" json:"id"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updateAt" json:"updatedAt"`
	DeletedAt time.Time `gorm:"column:deletedAt" json:"deletedAt"`
	IsDeleted bool      `json:"isDeleted"`
}
