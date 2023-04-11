package service

import (
	"fmt"
	"github.com/dapoadeleke/chupa/pkg/utility"
	"io/ioutil"
	"net/http"
)

type Service interface {
	MakeRequest(url string) (string, error)
}

type Impl struct{}

func (o *Impl) MakeRequest(url string) (str string, e error) {
	body, e := get(url)
	md5Str := utility.Hash(body)
	str = fmt.Sprintf("%s %x\n", url, md5Str)
	return
}

var get = func(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetService() Service {
	return &Impl{}
}
