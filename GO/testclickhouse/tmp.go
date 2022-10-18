package main

import "fmt"

func main() {
	d := new(Demo)
	d.Age = 18
	d.Name = "xxml"
	fmt.Println(*d)

	q := Demo{
		Name: "qqq",
		Age:  18,
	}
	fmt.Println(&q)
	fmt.Println("-----------")
	var v interface{} = uint64(89)
	n := 2
	switch v.(type) {
	case int:
		v = v.(int) + int(n)
	case int8:
		v = v.(int8) + int8(n)
	case int16:
		v = v.(int16) + int16(n)
	case int32:
		v = v.(int32) + int32(n)
	case int64:
		v = v.(int64) + int64(n)
	case uint:
		v = v.(uint) + uint(n)
	case uintptr:
		v = v.(uintptr) + uintptr(n)
	case uint8:
		v = v.(uint8) + uint8(n)
	case uint16:
		v = v.(uint16) + uint16(n)
	case uint32:
		v = v.(uint32) + uint32(n)
	case uint64:
		v = v.(uint64) + uint64(n)
	case float32:
		v = v.(float32) + float32(n)
	case float64:
		v = v.(float64) + float64(n)
	default:
		fmt.Errorf("The value for %s is not an integer", v)
	}
	fmt.Printf("v = %s, type of v = %T\n", v, v)

	fmt.Println("-------------")

	sch := &School{Students: nil, Address: "shanghai"}
	stu1 := &Student{School: sch, Name: "nana", Age: 18}
	stu2 := &Student{School: sch, Name: "lili", Age: 18}
	//sch2 := &School{Students: []*Student{stu1, stu2}, Address: "北京"}
	sch.Students = []*Student{stu1, stu2}
	fmt.Println(sch) // &{[0xc0000503e0 0xc000050400] shanghai }
	//fmt.Println(sch2) // &{[0xc0000503e0 0xc000050400] 北京 }

}

type Demo struct {
	Name string
	Age  int32
}

type School struct {
	Students   []*Student
	Address    string
	SchoolName string
}

type Student struct {
	School *School
	Name   string
	Age    int32
}
