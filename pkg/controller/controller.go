package controller

import (
	"fmt"
	"github.com/dapoadeleke/chupa/pkg/service"
	"github.com/dapoadeleke/chupa/pkg/utility"
	"log"
	"sync"
)

type Controller struct {
	service service.Service
}

func (o *Controller) Process(urls []string, parallel int) {
	maxConcurrentRequestLimiterChannel := make(chan bool, parallel)
	var waitGroup sync.WaitGroup
	for _, url := range urls {
		url = utility.PrepareUrl(url)
		if !utility.ValidUrl(url) {
			log.Printf("url not valid: %s \n", url)
			continue
		}
		waitGroup.Add(1)
		url := url
		go func() {
			maxConcurrentRequestLimiterChannel <- true
			defer func() {
				<-maxConcurrentRequestLimiterChannel
			}()
			if str, err := o.service.MakeRequest(url); err != nil {
				log.Printf("an error occurred, %s \n", err.Error())
			} else {
				fmt.Print(str)
			}
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
}

func GetController() Controller {
	return Controller{
		service: service.GetService(),
	}
}
