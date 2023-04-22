# sales-order-srv
### go操作es
```go
func InitElastic() {
	//host := "127.0.0.1:9200"
	client, err := elastic.NewClient()
	if err != nil {
		log.Printf("连接失败%s", err)
	}

	q := elastic.NewMatchQuery("address", "street")
	do, err := client.Search().Index("user").Query(q).Do(context.Background())
	if err != nil {
		panic(err)
	}

	total := do.Hits.TotalHits.Value
	fmt.Printf("搜索结果数量：%d", total)

	for _, value := range do.Hits.Hits {
		fmt.Printf("value= %v", value)
	}
}
```
### es添加数据
```go
func Create() {
	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}

	user := model.User{
		Id:   "1234",
		Name: "刘德华",
		Age:  66,
		Sex:  "男",
	}

	put, err := client.Index().Index("user").Id("1").BodyJson(user).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s:%s:%s", put.Id, put.Index, put.Type)

}

```
