package main

import (
	"fmt"
	"io/ioutil"
)

type contact struct {
	Id int
    email string

}

type Student struct {
	firstName, lastName string
	age                 int
	contact 
}

func (s Student) print(){
    fmt.Printf("%v\n", s)
}

func saveToFile(str string) {
	ioutil.WriteFile("student.txt", []byte(str), 0644)
}

func toString(s Student) string  {
	return fmt.Sprintf("%+v\n",s)
}

func (sptr *Student) UpdateName(Newfname string){
    sptr.firstName = Newfname
}

func main() {
     stud := Student{
	firstName: "John",
	lastName :  "Snow",
	age:       45,
	contact: contact{1, "abc@gmail.com"},
	}			

    stud2 := Student{"karan", "kundra", 25, contact{2, "pqr@gmail.com"},}
    fmt.Println("Before update name : ")
    stud.print()
    stud2.print()
    sptr := &stud
    sptr.UpdateName("Jonny")
    fmt.Println("============================================")
    fmt.Println("After update name : ")
    sptr.print()
    str:=toString(stud) + toString(stud2)
    saveToFile(str)

}

