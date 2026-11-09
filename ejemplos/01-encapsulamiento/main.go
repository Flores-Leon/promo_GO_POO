package main

import (
	"fmt"
	"promo-go-poo/juego"
)

func main() {
	heroe := juego.NewPersonaje("Hechicero", 100)
	fmt.Println(heroe.Nombre(), "arranca con", heroe.Vida(), "de vida")

	heroe.RecibirDano(30)
	fmt.Println("Recibe 30 de daño:", heroe.Vida())
	heroe.RecibirDano(-50) //como es negativo el objeto lo ignora
	fmt.Println("Daño negativo ignorado:", heroe.Vida())

	heroe.RecibirDano(500)
	fmt.Println("Recibe 500 de daño", heroe.Vida(), "| está vivo?", heroe.EstaVivo())
}
