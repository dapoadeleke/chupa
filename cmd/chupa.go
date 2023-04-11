package main

import (
	"flag"
	"github.com/dapoadeleke/chupa/pkg/controller"
)

func main() {
	ctrl := controller.GetController()
	parallel := flag.Int("parallel", 10, "number of parallel requests")
	flag.Parse()
	urls := flag.Args()
	ctrl.Process(urls, *parallel)
}
