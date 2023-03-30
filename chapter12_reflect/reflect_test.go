package kindpkg

import (
	"fmt"
	"reflect"
	"testing"
)

// https://medium.com/capital-one-tech/learning-to-use-go-reflection-822a0aed74b7

type Foo struct {
	Name string
	Age  int `json:"age" gorm:"age"`
}

// 测试指定方法 go test -v reflect_test.go -test.run TestStruct
func TestStruct(t *testing.T) {

	varType := reflect.TypeOf(Foo{})
	fmt.Println(varType)        // 包名.Foo
	fmt.Println(varType.Name()) // Foo
	fmt.Println(varType.Kind()) // struct
}

func TestField(t *testing.T) {

	foo := Foo{}
	varType := reflect.TypeOf(foo)
	if varType.Kind() == reflect.Struct {
		// 获取结构体属性数量
		fmt.Println(varType.NumField())
		for i := 0; i < varType.NumField(); i++ {
			f := varType.Field(i)
			fmt.Printf("name:%v, type:%v, kind:%v\n", f.Name, f.Type.Name(), f.Type.Kind())
			if f.Tag != "" {
				fmt.Printf("name:%v tags: json:%v, gorm:%v\n", f.Name, f.Tag.Get("json"), f.Tag.Get("gorm"))
			}
		}
	}

}

// slice
func TestSlice(t *testing.T) {

	foos := make([]Foo, 0)

	varType := reflect.TypeOf(foos)
	fmt.Println(varType) // []包名.Foo
	// Name() return the name of the type.
	// Some types, like a slice or a pointer, don’t have names and this method returns an empty string.
	fmt.Println(varType.Name()) // 空字符串
	fmt.Println(varType.Kind()) // slice

	switch varType.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		// slice 包含的元素类型
		eleType := varType.Elem()
		fmt.Println(eleType)        // 包名.Foo
		fmt.Println(eleType.Name()) // Foo
		fmt.Println(eleType.Kind()) // struct
	}

}

func TestChangeVarValue(t *testing.T) {
	greeting := "hello"
	gVal := reflect.ValueOf(greeting)
	fmt.Println(gVal.Interface())

	// 修改变量值
	gpVal := reflect.ValueOf(&greeting)
	gpVal.Elem().SetString("goodbye")
	fmt.Println(greeting)
}

func TestChangeStructValue(t *testing.T) {
	f := Foo{
		Name: "tom",
		Age:  18,
	}
	fVal := reflect.ValueOf(&f)
	// 修改结构体属性值
	fVal.Elem().Field(0).SetString("jack")
	//fVal.Elem().Field(0).Set(reflect.ValueOf("jack2"))
	fVal.Elem().Field(1).SetInt(20)

	fmt.Printf("%+v\n", f)
	fmt.Printf("%s\n", f.Name)
	fmt.Printf("%d\n", f.Age)
}

// 新创建一个结构体
func TestNewStruct(t *testing.T) {
	f := Foo{
		Name: "tom",
		Age:  18,
	}

	fType := reflect.TypeOf(f)
	// 创建一个新的 Value
	fVal := reflect.New(fType)

	fVal.Elem().Field(0).SetString("jack")
	fVal.Elem().Field(1).SetInt(20)

	f2 := fVal.Elem().Interface().(Foo)
	fmt.Printf("%+v\n", f2)
	fmt.Printf("%s\n", f2.Name)
	fmt.Printf("%d\n", f2.Age)
}
