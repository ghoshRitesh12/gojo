package utils

import (
	"github.com/ghoshRitesh12/gojo/src/config"
	"github.com/go-resty/resty/v2"
)

// var domain_name = "anicrush"

const (
	BASE_URL      = "https://anicrush.to"
	HOME_URL      = "https://anicrush.to/home"
	AJAX_BASE_URL = "https://api.anicrush.to"
)

func NewAnicrushClient() *resty.Client {
	return config.NewClient().
		SetHeader("Referer", BASE_URL)
}
