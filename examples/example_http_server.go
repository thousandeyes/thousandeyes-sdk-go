package main

import (
	"fmt"
	thousandeyes "go-thousandeyes"
	"os"
	"time"
)

func main() {
	client := thousandeyes.NewClient(os.Getenv("TE_TOKEN"))
	exampleTest := thousandeyes.HttpServer{
		TestName: "test",
		Url:      "https://dashboards.coqa.cloudreach.com",
		Interval: 300,
	}
	exampleTest.AddAgent(48620)
	createdTest, err := client.CreateHttpServer(exampleTest)
	if err != nil {
		panic(fmt.Sprintf("panicked creating test %s", err))
	}
	time.Sleep(time.Second * 2)
	update := thousandeyes.HttpServer{}
	update.AddAgent(47351)
	update.AddAgent(48620)
	t, err := client.UpdateHttpServer(createdTest.TestId, update)
	if err != nil {
		panic(fmt.Sprintf("panicked updating test %s", err))
	}
	fmt.Println(t)
}
