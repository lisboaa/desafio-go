package main

import "fmt"

func main() {
	var pontoEbulicaoAgua int = 100
	var graulCelsius float64 = 273.15
	var resultadoConversorCelsiusKelvin = conversorCelsiusKelvin(pontoEbulicaoAgua, graulCelsius)
	var resultadoConversorKelvinCelsius float64 = conversorKelvinCelsius(resultadoConversorCelsiusKelvin, graulCelsius)

	fmt.Printf("O ponto de ebulição da água é %.2f°C\n", resultadoConversorKelvinCelsius)
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
