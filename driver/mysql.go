package driver

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sales-order-srv/model"
)

var DB *gorm.DB

func InitDB() (err error) {
	dsn := "root:701XTAY1993@tcp(127.0.0.1:3306)/sales?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		return err
	}
	DB = db
	return nil
}

func init() {
	InitDB()

	err := DB.AutoMigrate(model.Cart{}, model.OrderProductInfo{}, model.OrderInfo{})
	if err != nil {
		fmt.Printf("初始化表失败%s", err.Error())
	}
}
