package main
import "fmt"

func digitFaktorial(n int)int{
	var i, keluaran int
	i = n-1
	for i > 0{
		n = n*i
		i--
	}
	for n>0{
		n = n/10
		keluaran++
	}
	return keluaran
}
func main(){
	var num int
	fmt.Scan(&num)
	fmt.Print(digitFaktorial(num))
}