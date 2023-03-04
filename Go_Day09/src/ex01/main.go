package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ch1 := make(chan string)
	done := make(chan struct{})
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("ON!")
	go func() {
		<-sig
		close(done)
	}()
	go sendURL(ch1, done)
	ch2 := crawlWeb(ch1, done)
	for i := range ch2 {
		fmt.Println(i)
	}
	fmt.Println("OFF!")
}

func sendURL(ch1 chan string, done chan struct{}) {
	for j := 1; j <= 13; j++ {
		str := fmt.Sprintf("https://simpsons-2x2.net/14-2x2/1-season/%d-episode", j)
		select {
		case ch1 <- str:
		case <-done:
			close(ch1)
			os.Exit(1)
		}
		time.Sleep(time.Second / 5)
	}
	close(ch1)
}

func crawlWeb(ch <-chan string, done chan struct{}) chan string {
	ch2 := make(chan string)
	maxGo := 8
	guard := make(chan struct{}, maxGo)
	go func() {
		wg := sync.WaitGroup{}
		for i := range ch {
			guard <- struct{}{}
			wg.Add(1)
			url := i
			go func() {
				getBody(&wg, url, ch2, done)
				<-guard
			}()
		}
		wg.Wait()
		close(ch2)
	}()

	return ch2
}

func getBody(s *sync.WaitGroup, url string, ch2 chan string, done chan struct{}) {
	defer s.Done()
	r, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	select {
	case ch2 <- string(body):
	case <-done:
		close(ch2)
		os.Exit(1)
	}
}
