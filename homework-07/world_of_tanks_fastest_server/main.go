//
// Задача:
//
// * Доработайте фрагмент из раздела «Буферизованные каналы», превратив его в полноценное
// приложение. Найдите три реальных зеркала (либо три разных сайта), замерьте их отклик, а
// затем протестируйте работу приложения, убедившись, что функция ​ mirroredQuery()
// действительно выполняет свою роль.

package main

import (
	"fmt"
	"github.com/sparrc/go-ping"
	"time"
)

type server struct {
	hostname string
	ip       string
	location string
}

var wot_servers = []server{
	server{"RU1", "92.223.6.30", "Россия, Москва"},
	server{"RU2", "92.223.5.15", "Россия, Москва"},
	server{"RU3", "92.223.21.240", "Германия, Франкфурт"},
	server{"RU4", "92.223.38.61", "Россия, Екатеринбург"},
	server{"RU5", "92.223.4.23", "Россия, Москва"},
	server{"RU6", "92.223.5.48", "Россия, Москва"},
	server{"RU7", "92.223.6.14", "Россия, Москва"},
	server{"RU8", "92.223.14.110", "Россия, Красноярск"},
	server{"RU9", "92.223.36.40", "Россия, Хабаровск"},
	server{"RU10", "88.204.200.209", "Казахстан, Павлодар"},
}

func main() {
	fmt.Println("Ищем ближайший игровой сервер World of Tanks в России...")
	fmt.Println("Программу необходимо запускать с правами администратора в windows или root в linux.")
	theNearestWOTServer := mirroredQuery()
	fmt.Println("Ближайший сервер:", theNearestWOTServer)
	time.Sleep(10 * time.Second)
}

func mirroredQuery() string {
	responses := make(chan string, len(wot_servers))

	for _, srv := range wot_servers {
		go func(out chan<- string, s server) {
			rtt := pingAvg5(s.ip)
			res := "ping " + s.hostname + " [" + s.location + "]: " + rtt.String()
			fmt.Println(res)
			out <- res
		}(responses, srv)
	}

	return <-responses
}

func pingAvg5(dst string) time.Duration {
	pinger, _ := ping.NewPinger(dst)
	pinger.Count = 5
	pinger.SetPrivileged(true)
	pinger.Run()
	stats := pinger.Statistics()
	return stats.AvgRtt
}
