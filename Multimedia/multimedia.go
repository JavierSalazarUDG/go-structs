package Multimedia

import "fmt"

type ContenidoWeb struct {
	Multimedia []Multimedia
}

func (mult *ContenidoWeb) Mostrar() {
	for _, m := range mult.Multimedia {
		m.Mostrar()
	}
}

type Multimedia interface {
	Mostrar()
}
type Imagen struct {
	Titulo  string
	Formato string
	Canales string
}

type Audio struct {
	Titulo   string
	Formato  string
	Duracion int
}

type Video struct {
	Titulo  string
	Formato string
	Frames  int
}

func (image *Imagen) Mostrar() {
	fmt.Println("**********Imagen**********")
	fmt.Println("Titulo:", image.Titulo)
	fmt.Println("Formato:", image.Formato)
	fmt.Println("Canales:", image.Canales)
	fmt.Println()
}
func (audio *Audio) Mostrar() {
	fmt.Println("**********Audio**********")
	fmt.Println("Titulo:", audio.Titulo)
	fmt.Println("Formato:", audio.Formato)
	fmt.Println("Duracion:", audio.Duracion)
	fmt.Println()
	fmt.Println()

}
func (video *Video) Mostrar() {
	fmt.Println("**********Video**********")
	fmt.Println("Titulo:", video.Titulo)
	fmt.Println("Formato:", video.Formato)
	fmt.Println("Frames:", video.Frames)
	fmt.Println()
}
