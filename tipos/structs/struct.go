package main
import (
	"fmt"
)

type produto struct{
	nome string
	preco float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64{
	return p.preco * (1-p.desconto)
}

func main(){
	produto1 := produto{
		nome: "Lápis",
		preco: 1.50,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Notebook", 1945.99, 0.15}
	fmt.Println(produto2.precoComDesconto())
}