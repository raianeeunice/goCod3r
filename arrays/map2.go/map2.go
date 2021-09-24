package main

import "fmt"

func main(){
	funcsESalarios := map[string]float64{
		"José João" : 11325.45,
		"Silva Silveira" : 12346.89,
		"Pedro Pedreira" : 23412.46,
	}

	funcsESalarios["Rafael Neto"] = 12345.62
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios{
		fmt.Println(nome, salario)
	}
}