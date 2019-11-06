//
//  Задача:
//
//  Самостоятельно изучите пакет ​ encoding/csv​ .
//  Напишите пример, иллюстрирующий его применение.
//


package main

import (
"encoding/csv"
"fmt"
"log"
"strings"
)

func main() {
	in := `first_name;last_name;username
"Rob";"Pike";rob
# lines beginning with a # character are ignored
Ken;Thompson;ken
"Robert";"Griesemer";"gri"
`
	r := csv.NewReader(strings.NewReader(in))
	r.Comma = ';'
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=============")
	fmt.Println("Source text:")
	fmt.Println("=============")
	fmt.Print(in)

	fmt.Println("==================")
	fmt.Println("Parsed  CSV array:")
	fmt.Println("==================")
	fmt.Print(records)
}