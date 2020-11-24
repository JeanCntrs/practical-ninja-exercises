**Ejercicio Práctico #3**

- Comenzando con este código, extrae los valores del canal usando un ciclo for range

    package main

    import (
        "fmt"
    )

    func main() {
        c := gen()
        recibir(c)

        fmt.Println("A punto de finalizar.")
    }

    func gen() <-chan int {
        c := make(chan int)

        for i := 0; i < 100; i++ {
            c <- i
        }

        return c
    }