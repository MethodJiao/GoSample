package main
 
import "fmt"

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
 }
func say(){
	fmt.Println("123")
}
func main() {
	var numbers = make([]int,3,5)
	go printSlice(numbers)
	printSlice(numbers)
	var i int
	fmt.Scan(&i)
}