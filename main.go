package main

import (
	"bufio"
	"fmt"
	"os"

	"./Multimedia"
)

func clear() {
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	println()
	println()
	println()
	println()
}
func main() {
	mul := Multimedia.ContenidoWeb{Multimedia: []Multimedia.Multimedia{}}
	for {
		fmt.Println("Estructuras en golang")
		fmt.Println("Selecciona una opcion")
		fmt.Println("1.- Registrar una nueva Imagen")
		fmt.Println("2.- Registrar un nuevo Audio")
		fmt.Println("3.- Registrar un nuevo Video")
		fmt.Println("4.- Mostrar")
		var opc int
		fmt.Scanln(&opc)
		switch opc {
		case 1:
			var titulo string
			var formato string
			var canales string
			fmt.Print("Ingresa un Titulo: ")
			fmt.Scanln(&titulo)
			fmt.Print("Ingresa un Formato: ")
			fmt.Scanln(&formato)
			fmt.Print("Ingresa un Canales: ")
			fmt.Scanln(&canales)
			mul.Multimedia = append(mul.Multimedia, &Multimedia.Imagen{titulo, formato, canales})
			fmt.Println("Muy bien, Imagen almacenada correctamente")
			clear()
		case 2:
			var titulo string
			var formato string
			var duracion int
			fmt.Print("Ingresa un Titulo: ")
			fmt.Scanln(&titulo)
			fmt.Print("Ingresa un Formato: ")
			fmt.Scanln(&formato)
			fmt.Print("Ingresa un Duracion: ")
			fmt.Scanln(&duracion)
			mul.Multimedia = append(mul.Multimedia, &Multimedia.Audio{titulo, formato, duracion})
			fmt.Println("Muy bien, Audio almacenada correctamente")
			clear()
		case 3:
			var titulo string
			var formato string
			var frames int
			fmt.Print("Ingresa un Titulo: ")
			fmt.Scanln(&titulo)
			fmt.Print("Ingresa un Formato: ")
			fmt.Scanln(&formato)
			fmt.Print("Ingresa un Duracion: ")
			fmt.Scanln(&frames)
			mul.Multimedia = append(mul.Multimedia, &Multimedia.Audio{titulo, formato, frames})
			fmt.Println("Muy bien, Video almacenada correctamente")
			clear()
		case 4:
			mul.Mostrar()
			clear()
		default:
			fmt.Println("La opcion seleccionada es invalida")
			clear()
		}
	}
}
