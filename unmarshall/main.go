package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func PrintUser(u User) {
	fmt.Printf("Name: %s\n", u.Name)
	fmt.Printf("Type: %s\n", u.Type)
	fmt.Printf("Age: %d\n", u.Age)
	fmt.Printf("vk: %s\n Facebook: %s\n", u.Social.Vkontakte, u.Social.Facebook)
}

func changeUserAge(u *User, age uint8) {
	u.Age = age
}

var myStruct struct {
	name string
	age  int
}

type Users struct {
	Users []User `json:"users"` // Типом значения поля "Users" является СЕГМЕНТ(Slice) значений типа "User", описанного ниже
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    uint8  `json:"age"`
	Social Social `json:"social"` // Типом значения поля "Social" является тип "Social", по сути это поле - структура в структуре
}

type Social struct {
	Vkontakte string `json:"vkontakte"`
	Facebook  string `json:"facebook"`
}

func main() {
	// 1. Создаём файловый дескриптор (создаётся новый поток ввода (из файла), насколько я понимаю,
	// ядро ОС возвращает файловый дескриптор, связанный с этим потоком)
	jsonFile, err := os.Open("C:\\Users\\vlad\\go\\src\\specialistGolangLevel2\\unmarshall\\users.json")
	if err != nil {
		log.Fatal(err)
	}
	// 2. Десериализуем содержимое users.json в экземпляр Go
	// 2.1 Инициализируем экземпляр типа "Users"
	var users Users
	fmt.Println(users)
	// 2.2 Вычитываем содержимое переменной jsonFile (точнее говоря, значение переменной - указатель на то, что там нужно)
	// в виде последовательности байт
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	// 2.3 Теперь задача - перенести всё из byteValue в users - это и есть десериализация
	fmt.Println(byteValue) // Тип byteValue - [], там лежит slice, содержащий в себе набор значений типа byte(uint8)
	json.Unmarshal(byteValue, &users)
	fmt.Println(users)
	//fmt.Println("In myUser:", myUser, reflect.TypeOf(myUser))
	//PrintUser()
	changeUserAge(&users.Users[0], 23)
	for _, user := range users.Users {
		PrintUser(user)
	}
}
