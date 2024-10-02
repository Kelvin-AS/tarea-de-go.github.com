package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    
    reader := bufio.NewReader(os.Stdin)
    
    
    fmt.Print("Por favor, ingresa tu nombre: ")
    nombre, _ := reader.ReadString('\n')
    
    
    nombre = strings.TrimSpace(nombre)
    
    
    if strings.Contains(nombre, "a") || strings.Contains(nombre, "A") {
        fmt.Println("¡Tu nombre contiene la letra 'a'!")
    } else {
        fmt.Println("Tu nombre no contiene la letra 'a'.")
    }
}