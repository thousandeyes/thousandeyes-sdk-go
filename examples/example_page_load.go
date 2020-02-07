package main

import (
	"fmt"
	thousandeyes "go-thousandeyes"
	"os"
	"time"
)

func main() {
	client := thousandeyes.NewClient(os.Getenv("TE_TOKEN"))
	exampleLoad := thousandeyes.PageLoad{
		TestName:     "test",
		Url:          "https://test.com",
		Interval:     300,
		HttpInterval: 300,
	}
	exampleLoad.AddAgent(48620)
	createdTest, err := client.CreatePageLoad(exampleLoad)
	if err != nil {
		panic(fmt.Sprintf("panicked creating test %s", err))
	}
	time.Sleep(time.Second * 2)
	update := thousandeyes.PageLoad{}
	update.AddAgent(47351)
	update.AddAgent(48620)
	t, err := client.UpdatePageLoad(createdTest.TestId, update)
	if err != nil {
		panic(fmt.Sprintf("panicked updating test %s", err))
	}
	fmt.Println(t)
}
