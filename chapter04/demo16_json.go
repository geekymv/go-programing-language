package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string   `json:"title" gorm:"title"`
	Year   int      `json:"released"` // tag
	Color  bool     `json:"color,omitempty"`
	Actors []string `json:"actors"`
}

func main() {

	movies := []Movie{
		{
			Title:  "大话西游",
			Year:   1996,
			Color:  false,
			Actors: []string{"周星驰", "朱茵", "吴孟达"},
		},
		{
			Title:  "大话西游2",
			Year:   1997,
			Color:  true,
			Actors: []string{"周星驰", "朱茵"},
		},
	}

	// 编码 将Go结构体转成JSON字符串
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("json.Marshal err:%v", err)
	}

	fmt.Println(string(data))

	// 解码 将JSON字符串转成Go结构体
	var m2 []Movie
	err = json.Unmarshal(data, &m2)
	if err != nil {
		log.Fatalf("json.Unmarshal err:%v", err)
	}
	fmt.Printf("%#v\n", m2)
}
