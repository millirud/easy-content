package myresty

import "github.com/go-resty/resty/v2"

func New(opts ...option) *resty.Client {
	client := resty.New()

	for _, option := range opts {
		option(client)
	}

	return client
}
