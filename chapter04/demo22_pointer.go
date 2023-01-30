package main

import "fmt"

type Sheet struct {
	TInfos []TInfo `json:"tInfo"`
}

type TInfo struct {
	Name string `json:"name"`
}

// 结构体指针
func ChangeName(s *Sheet) {
	for _, t := range s.TInfos {
		t.Name = "选择题"
	}
}

func main() {
	s := &Sheet{
		TInfos: []TInfo{
			{
				Name: "填空题",
			},
		},
	}
	ChangeName(s)

	for _, t := range s.TInfos {
		fmt.Println(t.Name)
	}
}
