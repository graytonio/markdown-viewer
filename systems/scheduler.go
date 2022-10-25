package systems

import (
	"log"
	"reflect"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

type Service func(*sync.WaitGroup)

var services = []Service{}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func AddService(serv Service) {
	services = append(services, serv)
}

func StartSystem() {
	for idx, system := range services {
		wg.Add(1)
		log.Printf("Starting System %d: %v\n", idx, GetFunctionName(system))
		go system(&wg)
	}
	wg.Wait()
}
