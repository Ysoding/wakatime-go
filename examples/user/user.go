package main

import (
	"fmt"
	"os"

	"github.com/Ysoding/wakatime-go/sdk"
	wakatimev1 "github.com/Ysoding/wakatime-go/services/wakatime/v1"
)

func main() {
	key := os.Getenv("WAKATIME_API_KEY")
	client, _ := wakatimev1.NewClient(key, nil)

	req := wakatimev1.NewUserRequest()
	req.Current = sdk.Bool(false)
	req.User = sdk.String("Tesla")

	res, err := client.User(req)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
