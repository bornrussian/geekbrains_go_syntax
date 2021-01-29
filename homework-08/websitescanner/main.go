//
// Задачи:
//
// Исследуйте работу последовательного и параллельного сканера веб-сайтов, задав большое
// (не менее 10) количество URL. Какие выводы можно сделать?
//
// 			Выводы - параллельное сканирование быстрее
//
//
//
//
// Какие практические варианты применения сканера веб-сайтов вы можете предложить?
//
// 			Если сканировать один и тот же сайт в несколько потоков, то можно организовать ему стресс-тест
//



package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {


	var sites []string

	if len(os.Args) > 1 {
		sites = os.Args[1:]
	} else {
		sites = []string{
			"https://google.ru",
			"https://youtube.ru",
			"https://habr.com",
			"https://drive2.ru",
			"https://github.com",
			"https://gmail.com",
			"https://mail.ru",
			"https://yandex.ru",
			"https://lenta.ru",
			"https://geekbrains.ru",
		}
	}


	fmt.Printf("\ngoroutine sites fetching:\n")
	start_go := time.Now()
	ch := make(chan string)
	for _, url :=range sites {
		go fetch_goroutine(url, ch)
	}
	for range sites {
		fmt.Print(<-ch) // receive from channel
	}
	start_go_length := time.Since(start_go).Seconds()
	fmt.Printf("%.2fs elapsed\n", start_go_length)


	fmt.Printf("\nstep by step sites fetching:\n")
	start_linear := time.Now()
	for _, url := range sites {
		fetch_linear(url)
	}
	start_linear_length := time.Since(start_linear).Seconds()
	fmt.Printf("%.2fs elapsed\n", start_linear_length)

	if start_go_length < start_linear_length {
		fmt.Printf("\nSo goroutine site fetching (%.2fs) is faster than linear (%.2fs).\n",start_go_length,start_linear_length)
	} else {
		fmt.Printf("\nSo linear site fetching (%.2fs) is faster than goroutine (%.2fs).\n",start_linear_length,start_go_length)
	}
}

func fetch_linear(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, url)
}

func fetch_goroutine(url string, ch chan<- string) {
	start:= time.Now()
	resp, err := http.Get(url)
	if err!=nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // to /dev/null
	resp.Body.Close()
	if err!=nil {
		ch <- fmt.Sprintf("error while reading %s: %v", url, err)
		return
	}
	ch <- fmt.Sprintf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, url)
}
