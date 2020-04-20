package main

import "fmt"

func main(){
	studentMap := make(map[int]string)
	studentMap[1] = "Rohan"
	studentMap[2] = "Mayuri"
	studentMap[3] = "Mayur"
	studentMap[4] = "Yogesh"
	studentMap[5] = "Dev"
	studentMap[6] = "Raj"

	for key,value := range studentMap{
		fmt.Printf("Roll number %d : %s\n",key,value)
	}

	updateStudentName(studentMap,"Raju", 6)
	fmt.Println("\nUpdated map : ",studentMap)

	removeFailures(studentMap, 4)
	fmt.Println("\nAfter removing failure : ",studentMap)
}

func updateStudentName(students map[int]string, updatedName string, rollNumber int){
	for key, _ := range students {
		if key == rollNumber{
			students[key]=updatedName
		}
	}
}

func removeFailures(students map[int]string, rollNumber int){
	delete(students, rollNumber)
}