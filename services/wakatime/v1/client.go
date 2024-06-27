package v1

import "github.com/Ysoding/wakatime-go/sdk"

const defaultEndpoint = "wakatime.com/api"

type Client struct {
	apikey string
	sdk.Client
}

func NewClient(apiKey string, config *sdk.Config) (*Client, error) {
	client := &Client{apikey: apiKey}
	if config == nil {
		config = sdk.NewConfig().WithEndoint(defaultEndpoint)
	}

	client.Init().WithSecret("TOKEN", apiKey).WithConfig(config)
	return client, nil
}
