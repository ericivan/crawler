package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"ericivan/crawler/model"
	"fmt"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver : got item #%d : %v", itemCount, item)
		}
	}()

	return out
}
func save(item interface{}) {
	client, err := elastic.NewClient()

	if err != nil {
		panic(err)
	}

	body := model.Profile{
		Name:      "golang",
		Age:       10,
		Marriage:  "no",
		Education: "college",
		Income:    "20000",
		Height:    0,
		Job:       "golang",
		Id:        "10111",
	}
	put, err := client.Index().Index("workspace").Type("golang").BodyJson(body).Do(context.Background())

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v", put)

}
