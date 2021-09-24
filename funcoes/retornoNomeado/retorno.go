package main
import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int){
	segundo = p2
	primeiro = p1
	return // retorno limpo
}

func main(){
	x,y := trocar(2,3)
	fmt.Println(x,y)

	segundo, primeiro := trocar(8,4)
	fmt.Println(segundo, primeiro)
}