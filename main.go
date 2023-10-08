package main

import (
	"fmt"
	"math/rand"
)

func main() {

	play()
}

func play() {
	numAleatorio := rand.Intn(100)
	var numIngresado int
	var numIntentos int
	const MaxIntentos = 10

	for numIntentos < MaxIntentos {
		numIntentos++
		fmt.Printf("\nIngrese número: ")
		fmt.Scan(&numIngresado)

		if numIngresado == numAleatorio {
			fmt.Println("¡Felicidades has Ganado!, el numero sorpresa es: ", numAleatorio)
			break
		} else if numIngresado < numAleatorio {
			fmt.Println("El número es mayor, ¡Súbele un poco! \n el número de intentos restantes es: ", MaxIntentos-numIntentos)
		} else if numIngresado > numAleatorio {
			fmt.Println("El número es menor, ¡Bájale un poco! \n el número de intentos restantes es: ", MaxIntentos-numIntentos)
		}
	}

	fmt.Println("¡Has perdido!, Mejor suerte a la próxima, el numero sorpresa era: ", numAleatorio)
	playAgain()
}

func playAgain() {
	var choice string

	fmt.Println("¿Desea jugar de nuevo? (s/n)")
	fmt.Scanln(&choice)

	switch choice {
	case "s":
		play()
	case "n":
		fmt.Println("Es una pena ='(, esperamos que regreses pronto")
	
	default:
		fmt.Println("Por favor, ingresa una opción válida, ¡Hasta la próxima!")

	}

}
