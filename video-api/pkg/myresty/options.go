package myresty

import (
	"time"

	"github.com/go-resty/resty/v2"
)

type option func(*resty.Client)

func WithUrl(url string) option {
	return func(r *resty.Client) {
		r.SetBaseURL(url)
	}
}

func WithTimeout(timeout time.Duration) option {
	return func(r *resty.Client) {
		r.SetTimeout(timeout)
	}
}
