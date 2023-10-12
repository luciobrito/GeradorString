package main

import (
    "fmt"
    "unicode/utf8"
    "math/rand"
)
        
func main() {
    var tamanhoString int
    fmt.Println("Digite o tamanho da string aleatória:")
    fmt.Scan(&tamanhoString)
    fmt.Println(stringAleatoria(tamanhoString));
}
func stringAleatoria(tamanho int) string{
    letras := "abcdefghijklmnopqrstuvwxyz"
    var resultado string;
    //Loop que vai se repetir conforme o usuário especificar
    for i := 0; i < tamanho; i++{
        //Contar o tamanho da string em UTF8
        randomico := numAleatorio(0, utf8.RuneCountInString(letras))
        //                     Transformar de bit para string
        resultado = resultado + string(letras[randomico])
    }
    return resultado;
}

//Função para gerar numéro aleatório em um intervalo em GO  
func numAleatorio(minimo int, maximo int) int{
    return minimo + rand.Intn(maximo-minimo);
} 