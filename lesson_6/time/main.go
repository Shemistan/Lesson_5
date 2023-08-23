package main

import (
	"fmt"
	"time"
)

func main() {
	one()
	two()
	tree()
	four()

}

func one() {
	// func Now() Time
	// возвращает текущую дату и время
	now := time.Now()

	// func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	// возвращает дату и время в соответствии с заданными параметрами: годом, месяцем, днем, временем и пр.
	currentTime := time.Date(
		2020,     // год
		time.May, // месяц
		15,       // день
		10,       // часы
		13,       // минуты
		12,       // секунды
		45,       // наносекунды
		time.UTC, // временная зона
	)

	// func Unix(sec int64, nsec int64) Time
	// возвращает дату и время в соответствии с заданными параметрами: секундами и наносекундами, прошедшими с начала эпохи Unix — 01.01.1970 г.
	// https://ru.wikipedia.org/wiki/Unix-%D0%B2%D1%80%D0%B5%D0%BC%D1%8F
	unixTime := time.Unix(
		150000, // секунды
		1,      // наносекунды
	)

	fmt.Println(now.Format("02-01-2006 15:04:05"))         // 15-05-2020 09:58:16
	fmt.Println(currentTime.Format("02-01-2006 15:04:05")) // 15-05-2020 10:13:12
	fmt.Println(unixTime.Format("02-01-2006 15:04:05"))    // 02-01-1970 22:40:00
}

func two() {
	// func Parse(layout, value string) (Time, error)
	// парсит дату и время в строковом представлении
	firstTime, err := time.Parse("2006/01/02 15-04", "2020/05/15 17-45")
	if err != nil {
		fmt.Println(err)
		return
	}

	// LoadLocation находит временную зону в справочнике IANA
	// https://www.iana.org/time-zones
	loc, err := time.LoadLocation("Asia/Yekaterinburg")
	if err != nil {
		fmt.Println(err)
		return
	}

	// func ParseInLocation(layout, value string, loc *Location) (Time, error)
	// парсит дату и время в строковом представлении с отдельным указанием временной зоны
	secondTime, err := time.ParseInLocation("Jan 2 06 03:04:05pm", "May 15 20 05:45:10pm", loc)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(firstTime.Format("02-01-2006 15:04:05"))  // 15-05-2020 17:45:00
	fmt.Println(secondTime.Format("02-01-2006 15:04:05")) // 15-05-2020 17:45:10
}

func tree() {
	current := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)

	// func (t Time) Date() (year int, month Month, day int)
	fmt.Println(current.Date()) // 2020 May 15

	// func (t Time) Year() int
	fmt.Println(current.Year()) // 2020

	// func (t Time) Month() Month
	fmt.Println(current.Month()) // May

	// func (t Time) Day() int
	fmt.Println(current.Day()) // 15

	// func (t Time) Clock() (hour, min, sec int)
	fmt.Println(current.Clock()) // 17 45 12

	// func (t Time) Hour() int
	fmt.Println(current.Hour()) //17

	// func (t Time) Minute() int
	fmt.Println(current.Minute()) // 45

	// func (t Time) Second() int
	fmt.Println(current.Second()) // 12

	// func (t Time) Unix() int64
	fmt.Println(current.Unix()) // 1589546712

	// func (t Time) Weekday() Weekday
	fmt.Println(current.Weekday()) // Friday

	// func (t Time) YearDay() int
	fmt.Println(current.YearDay()) // 136

	//------------------------------------------------------------

	//Конвертирование структуры Time в строку
	// func (t Time) Format(layout string) string
	current2 := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)
	fmt.Println(current2.Format("02-01-2006 15:04:05")) // 15-05-2020 17:45:12
	//------------------------------------------------------------
	//Сравнение структур Time
	firstTime := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)
	secondTime := time.Date(2020, time.May, 15, 16, 45, 12, 0, time.Local)

	// func (t Time) After(u Time) bool
	// true если позже
	fmt.Println(firstTime.After(secondTime)) // true

	// func (t Time) Before(u Time) bool
	// true если раньше
	fmt.Println(firstTime.Before(secondTime)) // false

	// func (t Time) Equal(u Time) bool
	// true если равны
	fmt.Println(firstTime.Equal(secondTime)) // false

	//------------------------------------------------------------

	//Методы, изменяющие структуру Time
	now := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)

	// func (t Time) Add(d Duration) Time
	// изменяет дату в соответствии с параметром - "продолжительностью"
	future := now.Add(time.Hour * 12) // перемещаемся на 12 часов вперед

	// func (t Time) AddDate(years int, months int, days int) Time
	// изменяет дату в соответствии с параметрами - количеством лет, месяцев и дней
	past := now.AddDate(-1, -2, -3) // перемещаемся на 1 год, два месяца и 3 дня назад

	// func (t Time) Sub(u Time) Duration
	// вычисляет время, прошедшее между двумя датами
	fmt.Println(future.Sub(past)) // 10332h0m0s
}

func four() {
	now := time.Now()
	now.Add(10 * time.Hour)
	past := now.AddDate(0, 0, -1)
	future := now.AddDate(0, 0, 1)

	// func Since(t Time) Duration
	// вычисляет период между текущим моментом и заданным временем в прошлом
	fmt.Println(time.Since(past).Round(time.Second)) // 24h0m0s

	// func Until(t Time) Duration
	// вычисляет период между текущим моментом и заданным временем в будущем
	fmt.Println(time.Until(future).Round(time.Second)) // 24h0m0s

	// func ParseDuration(s string) (Duration, error)
	// преобразует строку в Duration с использованием аннотаций:
	// "ns" - наносекунды,
	// "us" - микросекунды,
	// "ms" - миллисекунды,
	// "s" - секунды,
	// "m" - минуты,
	// "h" - часы.
	dur, err := time.ParseDuration("1h12m3s")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dur.Round(time.Hour).Hours()) // 1
}
