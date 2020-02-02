package main

import (
	"fmt"
	thousandeyes "go-thousandeyes"
	"os"
)

func main() {
	client := thousandeyes.NewClient(os.Getenv("TE_TOKEN"))
	exampleTest := thousandeyes.HttpTest{
		TestName: "test",
		Url:      "https://cloudreach.com",
		Interval: 300,
	}
	exampleTest.AddAgent(48620)
	t, err := client.CreateHttpTest(exampleTest)
	if err != nil {
		panic(fmt.Sprintf("panicked creating test %s", err))
	}
	fmt.Println(t)
}
