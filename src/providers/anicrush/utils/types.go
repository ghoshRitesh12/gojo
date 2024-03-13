package utils

type SubParserResp[T comparable] struct {
	Status bool `json:"status"`
	Data   T    `json:"data"`
}
