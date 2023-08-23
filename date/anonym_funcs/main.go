package main

import "fmt"

/*
Анонимная функция в Go - это функция, которая определена без имени.
Она позволяет вам создавать функции внутри других функций или

	передавать их как аргументы другим функциям.

Анонимные функции в Go также называются "замыканиями".
*/
func main() {
	//a := 10
	//b := 5
	//// Область видимости анонимной функции
	//// Внутри анонимной функции можно обращаться к переменным, которые были объявлены вне анонимной функции.
	//
	//// Вызываем анонимную функцию сразу после ее объявления.
	//
	//func(x, y int) {
	//	a++
	//	println(x + y)
	//}(a, b)
	//
	//fmt.Println("----> a", a)
	//
	//// Присваиваем анонимную функцию переменной anonymousFunc
	//anonymousFunc := func(x, y int) int {
	//	println(x + y)
	//	return x + y
	//}
	//
	//res := anonymousFunc(17, 22)
	//fmt.Println(res)
	//
	//fmt.Println(anonymousFunc(a, b))

	//// ---------------------------------------------------------
	// Замыкание анонимной функции
	//Когда анонимная функция использует переменные, объявленные за ее рамками, ее называют замыканием.

	fn := func() func(int) int {
		count := 0
		return func(i int) int {
			count++
			return count * i
		}
	}()

	for i := 1; i <= 5; i++ {
		fmt.Println(fn(i)) // 1, 2, 3, 4, 5
	}

	for i := 10; i <= 15; i++ {
		fmt.Println(fn(i)) // 6, 7, 8, 9, 10
	}

}

// ---------------------Пример использования---------------------
//// Функция calculate принимает на вход функцию сигнатуры func(int, int) int
//// и выполняет переданную функцию с аргументами x и y.
//func calculate(x, y int, operation func(int, int) int) int {
//	result := operation(x, y)
//	return result
//}
//
//func main() {
//	// Передаем анонимную функцию в calculate, которая выполняет сложение.
//	sum := calculate(10, 5, func(a, b int) int {
//		return a + b
//	})
//
//	// Передаем другую анонимную функцию, которая выполняет вычитание.
//	difference := calculate(10, 5, func(a, b int) int {
//		return a - b
//	})
//
//	fmt.Println("Sum:", sum)
//	fmt.Println("Difference:", difference)
//}

// ---------------------в defer---------------------

//func main() {
//
//
//	// ---------------------------------------------------------
//	//Отложенный вызов анонимной функции
//	conn := &Conncet{
//		Host: "localhost",
//		Port: 5432,
//		TTL:  30,
//	}
//
//	// устанавливаем соединение с БД
//	// и сразу же откладываем закрытие соединения
//	defer func() {
//		err := conn.CloseConnectDB()
//		if err != nil {
//			fmt.Println("failed to close connection to database")
//		}
//	}()
//
//	/*
//		Какая-то логика связанная с базой данных
//
//		происходит соединение с БД
//
//		происходит отправка запроса и необходимо закрыть созданное соединение с БД
//
//	*/
//}
//
//type Conncet struct {
//	Host string
//	Port int
//	TTL  int
//}
//
//func (c *Conncet) CloseConnectDB() error {
//	return nil
//}
