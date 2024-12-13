package main

import (
	"fmt"
	"time"
)

func main() {
	var canal01 chan string = make(chan string)
	var canal02 chan string = make(chan string)

	go ping(canal01, canal02)
	go lerMensagem(canal01, canal02)
	var escreva string
	fmt.Scanln(&escreva)
	// var pontoEbulicaoAgua int = 100
	// var graulCelsius float64 = 273.15
	// var resultadoConversorCelsiusKelvin = conversorCelsiusKelvin(pontoEbulicaoAgua, graulCelsius)
	// var resultadoConversorKelvinCelsius float64 = conversorKelvinCelsius(resultadoConversorCelsiusKelvin, graulCelsius)

	// fmt.Printf("O ponto de ebulição da água é %.2f°C\n", resultadoConversorKelvinCelsius)

	// const contador int = 100
	// pinPan(contador)

}

func conversorCelsiusKelvin(pontoEbulicaoAgua int, graulCelsius float64) float64 {
	resultado := float64(pontoEbulicaoAgua) + graulCelsius
	return resultado
}

func conversorKelvinCelsius(pontoEbulicaoAgua float64, graulCelsius float64) float64 {
	resultado := float64(pontoEbulicaoAgua) - graulCelsius
	println(resultado)
	return resultado
}

func pinPan(valor int) {
	for i := 1; i <= valor; i++ {
		if i%3 == 0 {
			fmt.Println("Pin")
		} else {
			fmt.Println(i)
		}
		if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}

func ping(canal01 chan string, canal02 chan string) {
	const contador = 10
	for i := 0; i < contador; i++ {
		canal01 <- "Ping"
		canal02 <- "Pong"
	}

}

func lerMensagem(canal01 chan string, canal02 chan string) {
	for {
		mensagem := <-canal01
		// mesangemCanal2 := <-canal02
		fmt.Println(mensagem)
		time.Sleep(time.Second * 1)
		mensagemCanal2 := <-canal02
		fmt.Println(mensagemCanal2)
		time.Sleep(time.Second * 1)
	}
}
