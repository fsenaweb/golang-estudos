package main

import (
    "fmt"
    "os"
)

func main() {

    var opcao string
    var valorOrigem float64
    var unidadeOrigem string

    var valorDestino float64
    var unidadeDestino string

    fmt.Println("CONVERSOR DE TEMPERATURA E DISTÂNCIA")
    fmt.Println("O que deseja converter? (T) Temperatura ou (D) Distância?")
    fmt.Scanf("%s\n", &opcao)

    if opcao == "T" {
        fmt.Println("Informe a temperatura:")
        fmt.Scanf("%f\n", &valorOrigem)
        fmt.Println("Converter para unidade: (C) Celsius ou (F) Fahrenheit?")
        fmt.Scanf("%s\n", &unidadeOrigem)
    } else {
        fmt.Println("Informe a distância:")
        fmt.Scanf("%f\n", &valorOrigem)
        fmt.Println("Converter para unidade: (Q) Quilometros ou (M) Milhas?")
        fmt.Scanf("%s\n", &unidadeOrigem)
    }

    if unidadeOrigem == "C" {
        unidadeDestino = "F"
    } else if unidadeOrigem == "F" {
        unidadeDestino = "C"
    } else if unidadeOrigem == "Q" {
        unidadeDestino = "M"
    } else if unidadeOrigem == "M" {
        unidadeDestino = "Q"
    } else {
        fmt.Printf("%s não é uma unidade conhecida!", unidadeDestino)
        os.Exit(1)
    }

    if unidadeOrigem == "C" {
        valorDestino = valorOrigem*1.8 + 32
    } else if unidadeOrigem == "F"{
        valorDestino = (valorOrigem - 32) / 1.8
    } else if unidadeOrigem == "Q" {
        valorDestino = valorOrigem * 0.62137
    } else {
        valorDestino = valorOrigem / 0.62137
    }

    if unidadeOrigem == "C" || unidadeOrigem == "F" {
        fmt.Printf("%.2f °%s = %.2f °%s\n",
            valorOrigem, unidadeOrigem,
            valorDestino, unidadeDestino)
    } else if unidadeOrigem == "Q" {
        fmt.Printf("%.2f km = %.2f mi\n",
            valorOrigem, valorDestino)
    } else if unidadeOrigem == "M" {
        fmt.Printf("%.2f mi = %.2f km\n",
            valorOrigem, valorDestino)
    } else {
        fmt.Printf("%s não é uma unidade conhecida!", unidadeOrigem)
        os.Exit(1)
    }


}