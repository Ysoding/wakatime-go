package main

import (
	"fmt"
	"os"

	wakatimev1 "github.com/Ysoding/wakatime-go/services/wakatime/v1"
)

func main() {
	key := os.Getenv("WAKATIME_API_KEY")
	client, _ := wakatimev1.NewClient(key, nil)

	// req := wakatimev1.NewAllTimeRequest()
	// req.Current = sdk.Bool(false)
	// req.UserID = sdk.String("Tesla")
	// req.Project = sdk.String("wakatime-go")

	// res, err := client.AllTime(req)
	res, err := client.AllTime(nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
