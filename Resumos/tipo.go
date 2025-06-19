package main

import "fmt"
import "reflect"

func main(){
	id := 8
	var nome = "Carlos"
nota := 8.2
	fmt.Println("Olá, jogador número", id,"!")
fmt.Println("Seu nome é:",  nome)
fmt.Println("Sua nota é:", nota)

porta := nota
tipo := reflect.TypeOf(porta)
fmt.Println("O tipo é:", tipo)}

