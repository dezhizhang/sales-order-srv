package main

import (
	"fmt"
	"sales-order-srv/driver"
)

func main() {
	driver.InitDB()
	fmt.Println("hello")
}
