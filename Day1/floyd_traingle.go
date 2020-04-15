package main
import "fmt"
 
func main(){
 //   var rows int
    var temp int = 1
 //   fmt.Print("Enter number of rows : ")
 //   fmt.Scan(&rows)
 
    for i := 1; i <= 5; i++ { 
         
        for k := 1; k <= i; k++ {
 
            fmt.Printf(" %d",temp)              
            temp++
        }
        fmt.Println("")     
    }
 }

// output :
/*
 1
 2 3
 4 5 6
 7 8 9 10
 11 12 13 14 15
*/