package main

import (
	"fmt"
)

func main(){
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678] = "Maria"
	aprovados[23456789] = "Pedro"
	aprovados[14567899] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados{
		fmt.Printf("%s (CPF: %d) \n", nome, cpf)
	}
	fmt.Println(aprovados[23456789])
	delete(aprovados, 23456789)
	fmt.Println(aprovados[23456789])
}