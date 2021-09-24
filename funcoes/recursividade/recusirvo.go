package main
import "fmt"

func fatorial(n int) int{
	if n>0{
		return n * fatorial(n-1)
	}else{
		return 1
	}
}

func main(){
	resultado := fatorial(-5)
	fmt.Println(resultado)
}