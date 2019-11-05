//
// Задача:
//
// Создать тип, описывающий контакт в телефонной книге.
// Создать псевдоним типа телефонной книги (массив контактов)
// и реализовать для него интерфейс​ Sort{}
//

package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type Phones []int

type Contact struct {
	Name string
	Phone Phones
}

type PhoneBook struct {
	Contacts []Contact
}

func (pb *PhoneBook) Add(contact Contact) {
	pb.Contacts=append(pb.Contacts,contact)
}

func (pb PhoneBook) ToJSON() ([]byte, error) {
	return json.Marshal(pb)
}

func (pb PhoneBook) ToBeautifulJSON() ([]byte, error) {
	return json.MarshalIndent(pb,"","\t")
}

func (pb *PhoneBook) FromJSON (js string) error {
	return json.Unmarshal([]byte(js),pb)
}

func (pb PhoneBook) Len() int {
	return len(pb.Contacts)
}

func (pb PhoneBook) Less(i, j int) bool {
	return pb.Contacts[i].Name < pb.Contacts[j].Name
}

func (pb *PhoneBook) Swap(i, j int) {
	pb.Contacts[i],pb.Contacts[j] = pb.Contacts[j],pb.Contacts[i]
}

func main() {
	var myPhoneBook PhoneBook
	//myPhoneBook.Add(Contact{"Миша", Phones{78293467382} } )
	//myPhoneBook.Add(Contact{"Никита", Phones{89167253764, 89635437382}})
	//myPhoneBook.Add(Contact{"Алёна", Phones{89123456789}})
	//myPhoneBook.Add(Contact{"!МЧС", Phones{112}})
	err:=myPhoneBook.FromJSON(`{"Contacts":[{"Name":"Миша","Phone":[78293467382]},{"Name":"Никита","Phone":[89167253764,89635437382]},{"Name":"Алёна","Phone":[89123456789]},{"Name":"!МЧС","Phone":[112]}]}`)
	if (err!=nil) {
		fmt.Println("Не получилось выполнить преобразование из JSON в телефонную книгу :-(")
		return
	}

	sort.Sort(&myPhoneBook)

	fmt.Println("Отсортированная телефонная книга по алфавиту:")

	js,err:=myPhoneBook.ToBeautifulJSON()
	if err==nil {
		fmt.Printf(string(js))
	} else {
		fmt.Println("Ошибка преобразования телефонной книги в JSON:",err)
	}
}