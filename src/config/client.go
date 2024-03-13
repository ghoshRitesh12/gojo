package config

import (
	"github.com/go-resty/resty/v2"
)

const (
	ACCEPT_HEADER          string = "application/json, text/plain, */*"
	USER_AGENT_HEADER      string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.4692.71 Safari/417.36"
	ACCEPT_ENCODING_HEADER string = "gzip, deflate, br, zstd"
	ACCEPT_LANGUAGE        string = "en-US,en;q=0.9,bn;q=0.8"
)

func NewClient() *resty.Client {
	baseClient := resty.New().
		SetHeader("Accept", ACCEPT_HEADER).
		SetHeader("User-Agent", USER_AGENT_HEADER).
		SetHeader("Accept-Language", ACCEPT_LANGUAGE)

	return baseClient
}
