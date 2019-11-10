//
// Задача:
//
// Дополните функцию ​ hello() ​ http-сервера так, чтобы принять и отобразить на странице один
// GET-параметр, например ​ name​ . При этом в браузере запрос будет выглядеть так:
// http://localhost/hello?name=World​
// Значение этого параметра надо получить в функции и отобразить при выводе html-кода.
//

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var wwwFilesRoot string = "/home/gauss/go/src/geekbrains_go_syntax/homework-06/http/"

// в шаблоне страницы /hello у нас описан описан параметр {{.Name}}
type HelloPageData struct {
	Name string
}

// эта функция обработает, когда сделают http запрос /hello
func helloPageHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")

	// ловим $_GET параметр "name" в структуру, которой будем парсить html template
	data := HelloPageData{Name:req.URL.Query().Get("name")}

	// подгружаем html template из заготовленного файла
	tmpl, err := template.ParseFiles(wwwFilesRoot+"template/hello.tmpl")
	if err != nil {
		log.Fatal("Could not parse template/hello.html")
	}

	// парсим html template, используя структуру
	tmpl.Execute(res, data)
}

func main() {
	// статичные файлы для http сервера: /index.html
	fs:= http.FileServer(http.Dir(wwwFilesRoot+"static"))
	http.Handle("/",fs)

	// динамические страницы для http сервера: /hello
	http.HandleFunc("/hello", helloPageHandler)

	// стартуем http сервер ...
	log.Println("Listening on :8000")
	err := http.ListenAndServe(":8000",nil)
	if err!=nil {
		fmt.Println(err)
	}
}
