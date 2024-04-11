package main
import "fmt"

type Service interface {
	fun(int, string) bool
	
}

type Service2 interface {
	fun(int) int
}

type K1 struct{

}
type K2 struct{

}
type K struct{
	K1
	K2
}

func (k K)fun(int) int {
	
	return 20
}  

func (k K)fun(int, string) bool {
	return true
} 



func main(){
	k := K{}

	fmt.Println(k.fun(20))
	fmt.Println(k.fun(30, "oekxl"))
}