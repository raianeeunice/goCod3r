package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"time"
)

func main(){
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo computador

	// forma reduzida de criar uma var
	area := PI * math.Pow(raio,2)
	fmt.Println("A área da circunferência é", area)

	var el, t string = "elefante", "tatu"
	fmt.Println(el, t)

	g, h, i := 2, false, "opa!"
	fmt.Println(g, h, i)

	var i2 rune ='a' // rune representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é: ", reflect.TypeOf(i2))
	fmt.Println(i2)

	s1 := "Olá meu nome é Rai"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da minha string é:", len(s1))

	s2 := `Olá
	meu
	nome é
	Rai`
	fmt.Println("O tamanho de s2 é :", len(s2))

	var a int
	var b float64
	var c bool 
	var d string
	var e *int
	fmt.Printf("%v, %v, %v, %q, %v", a, b, c, d, e) // valores 0 da linguagem Go


	x:= 2.4 //float
	y:= 2 // inteiro
	// não dá pra fazer simplesmente x/y
	fmt.Println(x/float64(y))
	fmt.Println(int(x)/y)

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println("Nota final: ", notaFinal)

	// cuidado... Não seja besta!
	fmt.Println("Nota " + string(rune(97)))

	// int para string
	fmt.Println("Nota " + strconv.Itoa(97))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num)

	p, _ := strconv.ParseBool("true") // true ou "1"
	if p{
		fmt.Println("Verdadeiro!", p)
	}

	// existe um comando go run *.go que serve tipo: tenho uma pasta e lá tenho um arquivo com uns metódos sem a função main e dentro dessa mesma pasta tenho um arquivo com a função main que chama esses pacotes
	
	// outras operações usando math...
	ri := 3.0
	ro := 2.0
	fmt.Println("Maior=>", math.Max(float64(ri), float64(ro)))
	fmt.Println("Menor=>", math.Min(ri, ro))
	fmt.Println("Exponenciação=>", math.Pow(ri, ro))

	// atribuiçao
	tico, teco := 1, 2
	fmt.Println(tico, teco)

	//troca de variaveis
	tico, teco = teco, tico
	fmt.Println(tico,teco)

	// comparando valores - não vê a posição na memória por enquanto
	d1 := time.Unix(0,0)
	d2 := time.Unix(0,0)
	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2))

	// apenas posfixado
	x++
	y --
}
