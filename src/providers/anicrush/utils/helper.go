package utils

import (
	"log"
	"sync"

	"github.com/go-resty/resty/v2"
)

func Parallelize(functions ...func()) {
	wg := sync.WaitGroup{}
	wg.Add(len(functions))

	defer wg.Wait()

	for _, fn := range functions {
		go func(fun func()) {
			defer wg.Done()
			fun()
		}(fn)
	}
}

func CaptureSubParserErrs(res *resty.Response, err error, url string) bool {
	log.Println("sub_parser:", url)

	if err != nil {
		if res.StatusCode() == 0 {
			log.Printf("Request failed with status: %d\n", res.StatusCode())
			return true
		}
		log.Println(err.Error())
		return true
	}

	if res.IsError() {
		log.Printf("Request failed with status: %d\n", res.StatusCode())
		return true
	}

	return false
}
