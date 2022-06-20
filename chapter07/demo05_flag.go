package main

import (
	"flag"
	"fmt"
)

// 2.5类型声明

type Celsius float64
type Fahrenheit float64

// FtoC 华氏温度转摄氏温度
func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g℃", c)
}

// 实现了 flag.Value 接口
type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "℃", "C":
		f.Celsius = Celsius(value)
		return nil
	case "℉", "F":
		f.Celsius = FtoC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {

	//c := FtoC(212.0)
	//fmt.Println(c)
	flag.Parse()
	fmt.Println(*temp)

}
