package main

import (
	"fmt"
	"os"

	wakatimev1 "github.com/Ysoding/wakatime-go/services/wakatime/v1"
)

func main() {
	key := os.Getenv("WAKATIME_API_KEY")
	client, _ := wakatimev1.NewClient(key, nil)

	res, err := client.Orgs(nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
