package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	MakeRequest()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}

func MakeRequest() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		scoupIter := i
		go func() {

			resp, err := http.Get("http://server:8080/api/v1/datalist/?num=10")
			if err != nil {
				log.Fatalln(err)
			}
			log.Println(resp)

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}
			log.Printf("%d | %s", scoupIter, string(body))
			wg.Done()
		}()
	}

	wg.Wait()
}
