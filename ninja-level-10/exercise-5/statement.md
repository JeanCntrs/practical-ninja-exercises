**Ejercicio Práctico #5**

- Demuestra el uso del idioma coma ok con este código

    package main

    import (
        "fmt"
    )

    func main() {
        c := make(chan int)

        v, ok :=
        fmt.Println(v, ok)

        close(c)
        
        v, ok =
        fmt.Println(v, ok)
    }
