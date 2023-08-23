package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//IO()
	BUFIO()
}

//func IO() {
//	//b := bytes.NewReader([]byte("Данные в объекте io.Reader"))
//	//
//	//// 'ReadAll' is deprecated
//	//dataIoutil, err := ioutil.ReadAll(b)
//	//
//	//if err != nil {
//	//	fmt.Println(err)
//	//}
//	////os
//	//fmt.Printf("%s\n", dataIoutil) // Данные в объекте io.Reader
//
//	// --------------------
//	//// создаем файл, если файл существоет - то он перезатрется
//	//f, _ := os.Create("text2.txt")
//	//
//	////До 1.21 требовалось закрытие файла
//	//_ = f.Name()
//	//
//	//// переименовываем файл
//	//os.Rename("text2.txt", "new_text.txt")
//	//
//	//// удаляем файл. os.Remove - до 1.21
//	//os.RemoveAll("new_text.txt")
//
//	// Путь начинается с корневой директории проекта, а не с того места где находиться исполняемый файл
//	// Просто читаем файл
//	data, err := os.ReadFile("./lesson_6/io/test.txt")
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(string(data))
//	// Открыть файл, при этом доступно много методов для работы с файлом, в доке можно почитать
//	file, err := os.OpenFile("./lesson_6/io/test.txt", os.O_RDWR|os.O_CREATE, 0755)
//	//defer func() {
//	//	errClose := file.Close()
//	//	if errClose != nil {
//	//		fmt.Println(errClose)
//	//	}
//	//}()
//
//	//Записали с конца (если использовать другие методы, то курсор может съехать)
//	res, _ := file.WriteString("\nSTRING")
//	data, _ = os.ReadFile("./lesson_6/io/test.txt")
//	fmt.Println(string(data))
//
//	fmt.Println(file.Name())
//
//	res, _ = file.Write([]byte("Добавили Это"))
//	_ = res
//
//	//data, _ = os.ReadFile("./lesson_6/io/test.txt")
//	//fmt.Println(string(data))
//	//
//	res, _ = file.WriteAt([]byte("Указали с какого байта"), 3)
//	data, _ = os.ReadFile("./lesson_6/io/test.txt")
//	fmt.Println(string(data))
//
//}

func BUFIO() {
	// Пакет bufio предоставляет буферизованные реализации для чтения и записи данных.
	// Он реализует интерфейсы io.Reader, io.Writer и io.ReaderWriter.
	// Буферизованные реализации предоставляют множество методов, не предоставляемых исходными интерфейсами.
	// Например, методы ReadString и WriteString.
	// Буферизованные реализации также предоставляют некоторые методы, которые не являются частью интерфейсов,
	// например, методы Peek, ReadBytes и ReadSlice.
	// Буферизованные реализации обычно более эффективны, чем исходные интерфейсы.
	// Буферизованные реализации могут быть полезны, даже если вам не нужна их эффективность,
	// потому что они предоставляют методы, которые не предоставляют исходные интерфейсы,
	// и потому что их ReadByte и UnreadByte методы работают правильно с неограниченными источниками данных.
	// Буферизованные реализации не предназначены для использования в нескольких горутинах одновременно.
	//
	//// bufio.Reader
	//file, err := os.Open("test.txt")
	//if err != nil {
	//	_ = err
	//}
	////defer file.Close()
	//
	//rd := bufio.NewReader(file)
	//
	//buf := make([]byte, 10)
	//n, err := rd.Read(buf) // читаем в buf 10 байт из ранее открытого файла
	//if err != nil && err != io.EOF {
	//	// io.EOF не совсем ошибка - это состояние, указывающее, что файл прочитан до конца
	//	println(err)
	//}
	//fmt.Printf("прочитано %d байт: %s\n", n, buf) // прочитано 10 байт: bufio ...
	//
	//fmt.Println("-----> читаем данные до разрыва абзаца '\\n'")
	//s, err := rd.ReadString('\n') // читаем данные до разрыва абзаца ('\n')
	//fmt.Printf("%s\n", s)
	//
	//// Читаем всё
	//for io.EOF != nil {
	//	s, err := rd.ReadString('\n')
	//	if err == io.EOF {
	//		break
	//	}
	//	fmt.Printf("%s\n", s)
	//}
	//---------------------------------------------

	//////bufio.Scanner
	//file, err := os.Open("test.txt")
	//if err != nil {
	//	println(err)
	//}
	////defer file.Close()
	//
	//s := bufio.NewScanner(file)
	//
	//for s.Scan() { // возвращает true, пока файл не будет прочитан до конца
	//	fmt.Printf("%s\n", s.Text()) // s.Text() содержит данные, считанные на данной итерации
	//}

	//---------------------------------------------
	// bufio.Writer
	file, err := os.Create("test_3.txt")
	if err != nil {
		println(err)
	}
	//defer file.Close()

	w := bufio.NewWriter(file)
	n, err := w.WriteString("Запишем строку")
	if err != nil {
		println(err)
	}
	fmt.Printf("Записано %d байт\n", n) // Записано 27 байт

	// bufio.Writer имеет собственный буфер, чтобы быть уверенным, что данные точно записаны,
	// вызываем метод Flush()
	w.Flush()
}
