//
// Задача:
//
// Сформулируйте предложения по улучшению сетевого чата.
// Реализуйте одно из своих предложений по улучшению сетевого чата.
//
// Предложение: Добавить Никнеймы, сохранив возможность присоединяться к чату через простой telnet
//

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type client chan<- string

type message struct {
	from string
	text string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan message)
)

// Описание кода из методички:
// Главная горутина прослушивает и принимает входящие сетевые подключения от клиентов. Для
// каждого из них создается новая горутина ​handleConn​.
func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("ChatServer: Listening on 0.0.0.0:9999")
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

// Описание кода из методички:
// Горутина broadcaster ​хранит информацию о всех клиентах и прослушивает каналы событий и
// сообщений, используя мультиплексирование с помощью ​select
func broadcaster() {
	clients := make(map[client]bool)
	nicknames := make(map[string]string)
	for {
		select {
		case msg := <-messages:
			if strings.HasPrefix(msg.text, "/") { // если нам пришла команда
				words := strings.Split(msg.text, " ")
				switch words[0] {
				case "/name": // команда "Изменить свой никнейм"
					if len(words) > 1 {
						nick := strings.Join(words[1:], " ")
						nicknames[msg.from] = nick
						for cli := range clients {
							cli <- "ChatServer: " + msg.from + " is now known as '" + nick + "'"
						}
					}
				case "/unname": // команда "Забыть свой никнейм", так же применяем её при выходе пользователя, чтобы забыть его никнейм
					nick := nicknames[msg.from]
					for cli := range clients {
						cli <- "ChatServer: '" + nick + "' is now known as " + msg.from
					}
					delete(nicknames, msg.from)
				}
			} else { // если нам пришло сообщение
				for cli := range clients {
					if nick, found := nicknames[msg.from]; found {
						cli <- nick + ": " + msg.text
					} else {
						cli <- msg.from + ": " + msg.text
					}
				}
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

// Описание кода из методички:
// Горутина handleConn создает новый канал исходящих сообщений для своего клиента и объявляет
// широковещателю о поступлении этого клиента по каналу entering . Затем она считывает каждую
// строку текста от клиента, отправляет их широковещателю по глобальному каналу входящих
// сообщений, предваряя каждое сообщение указанием отправителя. Когда от клиента получена вся
// информация, handleConn​ объявляет об убытии клиента по каналу ​leaving​ и закрывает подключение.
func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "Hello, " + who + " You can set your nickname with command:"
	ch <- "/name Your Nickname"
	messages <- message{"ChatServer", who + " has arrived"}
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- message{who, input.Text()}
	}
	leaving <- ch
	messages <- message{who, "/unname"}
	messages <- message{"ChatServer", who + " has left"}
	conn.Close()
}

// Описание кода из методички:
// Кроме того, ​handleConn создает горутину clientWriter для каждого клиента. Она получает
// широковещательные сообщения по исходящему каналу клиента и записывает их в его сетевое
// подключение. Цикл завершается, когда широковещатель закрывает канал, получив уведомление leaving​
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
