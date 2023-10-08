package main

import (
	"fmt"
	"math/rand"
)

func main() {

 play()
}

func play (){
	numAleatorio := rand.Intn(100)
	fmt.Println(numAleatorio)
	var numIngresado int;
	var numIntentos int;
	const MaxIntentos = 10;

	for numIntentos < MaxIntentos{
		numIntentos++;
		fmt.Printf("Ingrese número: ");
		fmt.Scan(&numIngresado);

		if numIngresado == numAleatorio{
			fmt.Println("¡Felicidades has Ganado!");
			break;
		}else{
			fmt.Printf("¡Fallaste! sigue intentando, te quedan %d", MaxIntentos - numIntentos)
		}
	}
}