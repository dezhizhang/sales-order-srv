package driver

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"sales-order-srv/model"
)

func InitElastic() {
	//host := "127.0.0.1:9200"
	//client, err := elastic.NewClient()
	//if err != nil {
	//	log.Printf("连接失败%s", err)
	//}
	//
	//q := elastic.NewMatchQuery("address", "street")
	//do, err := client.Search().Index("user").Query(q).Do(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//
	//total := do.Hits.TotalHits.Value
	//fmt.Printf("搜索结果数量：%d", total)
	//
	//for _, value := range do.Hits.Hits {
	//	fmt.Printf("value= %v", value)
	//}

	client, err := elastic.NewClient()
	if err != nil {
		log.Printf("连接失败%s", err)
	}
	q := elastic.NewMatchQuery("name", "street")
	result, err := client.Search().Index("user").Query(q).Do(context.Background())
	if err != nil {
		panic(err)
	}
	value := result.Hits.TotalHits.Value
	fmt.Printf("搜索结果:%d", value)

	for _, value := range result.Hits.Hits {
		fmt.Printf("value:%v", value)
	}
}

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
	put, err := client.Index().Index("user").Id("1").BodyJson(&user).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf(":%s:%s:%s", put.Id, put.Index, put.Type)

}

const mapping = `{
	"mappings":{
		"properties":{
				"name":{
					"type":"text",
					"analyzer":"ik_max_word"
				},
				"id":{
					"type":"integer"
				}
			}
	}
}`

func Mappping() {

	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}
	createIndex, err := client.CreateIndex("user").BodyJson(mapping).Do(context.Background())
	if err != nil {
		panic(err)
	}
	if !createIndex.Acknowledged {

	}

}
