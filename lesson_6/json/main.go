package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	//one()
	//two()
	//tree()
	//four()
	//
	//five()
	//
	six()

	//seven()

}

func one() {
	type myStruct struct {
		Name   string
		Age    int
		Status bool
		Values []int
	}

	s := myStruct{
		Name:   "John Connor",
		Age:    35,
		Status: true,
		Values: []int{15, 11, 37},
	}

	// Функция Marshal принимает аргумент типа interface{} (в нашем случае это структура)
	// и возвращает байтовый срез с данными, кодированными в формат JSON.
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", data) // {"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}
}

func two() {
	type myStruct struct {
		Name   string
		Age    int
		Status bool
		Values []int
	}

	s := myStruct{
		Name:   "John Connor",
		Age:    35,
		Status: true,
		Values: []int{15, 11, 37},
	}

	data, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", data)

	//{
	//	"Name": "John Connor",
	//	"Age": 35,
	//	"Status": true,
	//	"Values": [
	//		15,
	//		11,
	//		37
	//	]
	//}
}

//func tree() {
//	data := []byte(`{"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}`)
//
//	type myStruct struct {
//		Name   string
//		Age    int
//		Status bool
//		Values []int
//	}
//
//	var s myStruct
//
//	if err := json.Unmarshal(data, &s); err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Printf("%+v", s) // {John Connor 35 true [15 11 37]}
//}

//func four() {
//	type user struct {
//		Name     string
//		Email    string
//		Status   bool
//		Language []byte
//	}
//
//	m := user{Name: "John Connor", Email: "test email"}
//
//	data, _ := json.Marshal(m)
//
//	data = bytes.Trim(data, "{") // испортим json удалив '{'
//
//	// функция json.Valid возвращает bool, true - если json правильный
//	if !json.Valid(data) {
//		fmt.Println("invalid json!") // вывод: invalid json!
//	}
//
//	fmt.Printf("%s", data) // вывод: "Name":"John Connor","Email":"test email","Status":false,"Language":null}
//}

func five() {
	// в общем виде аннотация выглядит так: `json:"используемое_имя,опция,опция"`

	type myStruct struct {
		// при кодировании / декодировании будет использовано имя name, а не Name
		Name string `json:"name"`

		// при кодировании / декодировании будет использовано то же имя (Age),
		// но если значение поля равно 0 (пустое значение: false, nil, пустой слайс и пр.),
		// то при кодировании оно будет опущено
		Age int `json:",omitempty"`

		// при кодировании / декодировании поле всегда игнорируется
		Status bool `json:"-"`
	}

	m := myStruct{Name: "John Connor", Age: 0, Status: true}

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", data) // {"name":"John Connor"}
}

func six() {
	type testStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var (
		src = testStruct{Name: "John Connor", Age: 35} // структура с данными
		dst = testStruct{}                             // структура без данных
		buf = new(bytes.Buffer)                        // буфер для чтения и записи
	)

	enc := json.NewEncoder(buf)
	dec := json.NewDecoder(buf)

	enc.Encode(src)
	dec.Decode(&dst)

	fmt.Print(dst) // {John Connor 35}
}

func seven() {
	file, err := os.Open("./lesson_6/json/test.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer file.Close()

	var s fileStruct
	if err := json.NewDecoder(file).Decode(&s); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", s)

	data, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("\nФорматированный вывод \n\n")
	fmt.Printf("%s", data)
}

type fileStruct struct {
	HeaderTitle    string   `json:"header_title"`
	ShareLink      string   `json:"share_link"`
	SelectedGender string   `json:"selected_gender"`
	BrandId        string   `json:"brand_id"`
	BrandGroupId   string   `json:"brand_group_id"`
	GenderMenu     []string `json:"gender_menu"`
	Article        struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Title       string `json:"title"`
		Subtitle    string `json:"subtitle"`
		Link        string `json:"link"`
		ImageMobile string `json:"image_mobile"`
	} `json:"article"`
	HeaderBannerCarousel struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Pages []struct {
			Id          int    `json:"id"`
			Title       string `json:"title"`
			ShowButton  bool   `json:"show_button"`
			ButtonText  string `json:"button_text"`
			Link        string `json:"link"`
			ImageMobile string `json:"image_mobile"`
		} `json:"pages"`
	} `json:"header_banner_carousel"`
}
