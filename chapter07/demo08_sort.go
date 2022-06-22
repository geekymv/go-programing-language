package main

import (
	"fmt"
	"sort"
)

type Stu struct {
	Name  string
	Score int
}

type StuSort struct {
	data []*Stu
	less func(i, j *Stu) bool
}

func (s StuSort) Len() int {
	return len(s.data)
}

func (s StuSort) Less(i, j int) bool {
	return s.less(s.data[i], s.data[j])
}

func (s StuSort) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

func main() {

	stuList := []*Stu{
		{
			Name:  "tony",
			Score: 90,
		},
		{
			Name:  "jack",
			Score: 70,
		},
		{
			Name:  "tom",
			Score: 90,
		},
	}

	// 先根据成绩排序，成绩相同再根据姓名排序
	sutSort := StuSort{
		data: stuList,
		less: func(i, j *Stu) bool {
			if i.Score != j.Score {
				return i.Score > j.Score
			}
			return i.Name < j.Name
		},
	}

	sort.Sort(sutSort)

	for _, v := range stuList {
		fmt.Println(*v)
	}
}
