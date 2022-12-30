package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
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
	count, err := strconv.Atoi(os.Getenv("COUNT"))
	if err != nil {
		log.Println(err)
		count = 3
	}
	if count < 1 {
		count = 1
	}
	num, err := strconv.Atoi(os.Getenv("NUM"))
	if err != nil {
		log.Println(err)
		num = 1
	}
	if num < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		wg.Add(1)
		scoupIter := i
		go func() {

			resp, err := http.Get(fmt.Sprintf("http://server:8080/api/v1/datalist/?num=%d", num))
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
