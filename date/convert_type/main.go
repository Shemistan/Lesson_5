package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Преобразование данных
	//
	////Приведение целочисленных типов
	//var a int8 = 10
	//var b int16 = 20
	//var c int16 = 127
	//var d int16 = 232
	//
	//aInt32 := int32(a) //Приведение типа int8 к типу int32
	//bInt32 := int32(b) //Приведение типа int16 к типу int32
	//cInt8 := int8(c)   //Приведение типа int16 к типу int8
	//dInt8 := int8(d)   //Приведение типа int16 к типу int8, при этом теряется часть данных (232 -> -24)
	//fmt.Printf("%T %v\n", aInt32, aInt32)
	//fmt.Printf("%T %v\n", bInt32, bInt32)
	//fmt.Printf("%T %v\n", cInt8, cInt8)
	//fmt.Printf("%T %v\n", dInt8, dInt8)

	//-----------------------------------------------
	//
	////Приведение целых чисел и чисел с плавающей точкой
	//var a int64 = 1234
	//var b = float64(a)
	//fmt.Println("b -->", b)
	//
	////Преобразование чисел с плавающей точкой в целые числа (с потерей данных)
	//var c float64 = 123.456
	//var d = int32(c)
	//fmt.Println("d -->", d)
	//
	////Числа, конвертируемые с помощью деления с остатком
	//var e int32 = 1234
	//var f int16 = 4321
	//var g = int16(e % int32(f))
	//var h = 11 / 3
	//fmt.Println("g -->", g)
	//fmt.Println("h -->", h)

	//-----------------------------------------------
	////Конвертация строк в байты/rune и обратно
	//a := "текст"
	//b := []byte(a)
	//c := string(b)
	//fmt.Println("a -->", a) // текст
	//fmt.Println("b -->", b) // [209 130 208 189 209 130 208 181]
	//fmt.Println("c -->", c) // текст
	//
	////Тоже самое работает и со срезами типа rune:
	//d := []rune(a)
	//e := string(d)
	//fmt.Println("d -->", d) //[115 116 114 105 110 103]
	//fmt.Println("e -->", e) //текст

	//-----------------------------------------------
	////Конвертация в строки
	//a := "string"
	//b := 1234
	////fmt.Println(a + "text" + b)  - ошибка
	//fmt.Println(a + "text" + fmt.Sprint(b)) //stringtext1234  можно так
	//// А есть библиотека strconv, которая позволяет конвертировать в строку любой тип данных
	//fmt.Println(a + "text" + strconv.Itoa(b)) //stringtext1234  можно так
	////Интересное дополнение, метод Itoa это всего-лишь обертка для FormatInt: (кусок исходного кода пакета strconv)
	///*
	//		func Itoa(i int) string {
	//			return FormatInt(int64(i), 10)
	//		}
	//
	//	То-есть вызывая метод Itoa мы по сути вызываем FormatInt который принимает систему счисления в качестве 2 аргумента,
	//	но туда сразу передается - десятичная система счисления.
	//
	//	Но никто нам не мешает напрямую вызыватьFormatInt, полезно если работаем с разными системами счисления:
	//*/
	//
	//// приставка '0x' означает что число в шестнадцатеричной системе счисления
	//var c int64 = 0xB                     // 'B' в шестнадцатеричной это 11 в десятичной системе счисления
	//fmt.Println(strconv.FormatInt(c, 2))  //1011 - двоичная система счисления
	//fmt.Println(strconv.FormatInt(c, 10)) //11 - десятичная система счисления
	//fmt.Println(strconv.FormatInt(c, 16)) //b - шестнадцатеричная система счисления

	////-----------------------------------------------
	//
	////Конвертация целых беззнаковых чисел в строку
	//var a uint64 = 1234
	//resUint64 := strconv.FormatUint(a, 10)
	//fmt.Println(resUint64) //1234 - строка
	//fmt.Printf("%T\n", resUint64)
	//
	////Конвертация чисел с плавающей запятой в строку
	//var b float64 = 123.45678
	//resFloat64 := strconv.FormatFloat(b, 'f', 2, 64) // 'f' - формат, 2 - кол-во знаков после запятой, 64(32) - тип float64(32)
	//fmt.Println(resFloat64)                          //123.46 - строка
	//fmt.Printf("%T\n", resFloat64)                   //string
	//
	//// если мы хотим учесть все цифры после запятой, то можем в prec передать -1
	//resFloat64All := strconv.FormatFloat(b, 'f', -1, 64)
	//fmt.Println(resFloat64All) //123.45678 - строка
	//
	////Конвертация bool в string
	//var c bool = true
	//resBool := strconv.FormatBool(c)
	//fmt.Println(resBool) //true - строка
	//fmt.Printf("%T\n", resBool)

	////-----------------------------------------------
	// Конвертация строк в другие типы
	//Конвертация строк в целые числа
	var a string = "1234"
	resA, _ := strconv.ParseInt(a, 10, 64) // 10 - система счисления, 64 - тип int64
	fmt.Println(resA)                      //1234 - int64
	fmt.Printf("%T\n", resA)

	c := "1234"
	d := int64(5678)
	//fmt.Println(c + d)  - ошибка
	convertC, _ := strconv.Atoi(c)
	summ := d + int64(convertC)
	fmt.Println(summ) //6912

	//Строка это не число
	var b string = "123.456Е"
	convertB, err := strconv.ParseInt(b, 10, 64)
	if err != nil {
		fmt.Println(err.Error()) // strconv.ParseInt: parsing "123.456Е": invalid syntax
	}
	fmt.Println(convertB) //0

	// Конвертация string в float с помощью метода ParseFloat:
	var e string = "123.456"
	resE, _ := strconv.ParseFloat(e, 32) // 32 - тип float32
	fmt.Println(resE)                    //123.456
	fmt.Printf("%T\n", resE)             //float64

	//Конвертация string в bool
	var f string = "true"
	resF, _ := strconv.ParseBool(f)
	fmt.Println(resF)        //true
	fmt.Printf("%T\n", resF) //bool

}
