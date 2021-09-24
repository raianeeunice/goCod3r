package main

import "fmt"

func main(){
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabigol Silveira" : 12345.6,
			"Gunga Delgado" : 23104.6,
		},

		"J" : {
			"José Maria" : 12349.4,
			"Juvents Junior" : 32455.6,
		},

		"P" : {
			"Pedro Pereira" : 98564.8,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra{
		fmt.Println(letra, funcs)
	}
}