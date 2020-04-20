package main

import "fmt"

func main(){
	var slice =[]string{"This","is","a","Assignment", "For", "Go"}

	defer deferDemo("one")
	defer deferDemo("two")
	fmt.Println("The length of slice is : ", len(slice))

	fmt.Println("\nUpdate test")
	fmt.Println("Before update : ",slice)
	updateSlice(slice, 4, "of")
	fmt.Println("After update : ",slice)

	fmt.Println("\nNew slice update and compare test")
	newSlice := slice
	updateSlice(newSlice, 5, "GO Lang")
	fmt.Println("New slice after update  :", newSlice)
	compareSlices(slice,newSlice)
}


func updateSlice(slice []string, index int, updateWith string){
	slice[index] = updateWith
	//The below doesn't work, probably because append works on value rather than the reference
	//slice=append(slice,"And this is a new element")
}

func compareSlices(slice[] string, updatedSlice[] string){
	for index, value :=range slice{
		if updatedSlice[index] != value {
			fmt.Println("Compared slices aren't the same")
			fmt.Printf("Abberation at index %d, value %s\n", index, value)
			return
		}
	}
	fmt.Println("Compared slices are the same")
}


func deferDemo(toPrint string){
	fmt.Printf("\n%s this is the last statement",toPrint)
}