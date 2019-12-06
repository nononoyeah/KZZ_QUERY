package main

import "fmt"

// Person 人
type Person struct {
	name string
	sex  byte
	age  int
}

// Student 学生
type Student struct {
	Person // 继承至Person
	id     int
	name   string // 优先级高于Person中的name
}

func main() {

	// 结构体赋值
	s1 := Student{Person{"mike", 'm', 20}, 1, "haha"}
	// s1= {{mike 109 20} 1}
	fmt.Println("s1=", s1)
	fmt.Println("s1.Person.name=", s1.Person.name)
	fmt.Println("s1.name=", s1.name)

	s2 := Student{Person: Person{"mike", 'm', 20}, id: 10}
	//s2= {Person:{name:mike sex:109 age:20} id:10}
	fmt.Printf("s2= %+v\n", s2)

}
