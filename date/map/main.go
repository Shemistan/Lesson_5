package main

import "fmt"

func main() {
	//panic
	//var userlistPanic map[string]int
	//userlistPanic["login"] = 12343
	//userlistPanic, _ := getUserList()

	//------------------------------------------------
	//тут явно инициализируем пустуйю мапу и всё ок
	//var userlistOk = map[string]int{}
	//userlistOk["login"] = 12343
	//
	////можно и так
	//userlist := make(map[string]int)
	//userIds := map[int]string{
	//	1: "andrey@gmail.com",
	//	2: "ruslan@gmail.com",
	//}
	//fmt.Println(userlist, "\n", userIds)

	// ---> Добавление элемента
	productPrice := make(map[string]int)
	productPrice["apple"] = 11900
	productPrice["banana"] = 6900
	productPrice["apple"] = 1200
	//
	//// получение элемента
	//fmt.Println("цена яблока--->", productPrice["apple"])
	//
	//// если ничего нет - то вернет zero_value
	//fmt.Println("zero_value--->", productPrice["potato"])

	// удаление
	//fmt.Println("исходная мапа--->", productPrice)
	//delete(productPrice, "apple") // Удаление элемента по ключу "apple"\
	//delete(productPrice, "potato") // Можно удалить несуществующий ключ - это безопасно
	//fmt.Println("после удаления--->", productPrice)

	//// проверка на существование ключа
	//price, ok := productPrice["apple"]
	//if ok {
	//	fmt.Println("цена яблока--->", price)
	//} else {
	//	fmt.Println("яблок нет")
	//}

	// проверка на существование ключа без получения значения
	//if v, ok := productPrice["apple"]; ok {
	//	fmt.Println(v)
	//	fmt.Println("яблоки есть")
	//} else {
	//	fmt.Println("яблок нет")
	//}
	//
	// итерация по мапе
	//for key, value := range productPrice {
	//	fmt.Println(key, " - ", value)
	//}

	// итерация по мапе без значений
	//for key := range productPrice {
	//	fmt.Println(key)
	//}

	// итерация по мапе без ключей
	for _, value := range productPrice {
		fmt.Println(value)

	}

	//// итерация по мапе без ключей и значений
	//for range productPrice {
	//	fmt.Println("итерация")
	//}
	//
	// длина мапы
	fmt.Println("длина мапы--->", len(productPrice))

	// мапа в мапе
	var users = map[string]map[string]string{}
	users["login_password"] = map[string]string{
		"user_login_1": "password_1",
		"user_login_2": "password_2",
	}

	users["names"] = map[string]string{
		"user_login_1": "name_1",
		"user_login_2": "name_2",
	}

	fmt.Println(users["login_password"]["user_login_2"])

}

func getUserList() (map[string]int, error) {
	/*
		Запрос куда-то для получения map
		........

		какая та логика
		......
	*/

	return map[string]int{
		"login_1": 123,
		"login_2": 456,
	}, nil
}
