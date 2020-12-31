package main
 
import "fmt"
import "strconv"
func multResult(x int) (int,int) {
	return x,x+1
}
func main() {
	x, y := multResult(2)
	var z = x+y
	str := strconv.Itoa(z)// :=声明并赋值
	fmt.Printf(str)
}