**Ejercicio Práctico #1**
- Comienza con este código. En vez de usar el identificador blank (underscore), asegúrate de que el código esté chequeando y manejando el error.
    ```
    package main

    import (
        "encoding/json"
        "fmt"
    )

    type persona struct {
        Nombre   string
        Apellido string
        Frases   []string
    }

    func main() {
        p1 := persona{
            Nombre:   "James",
            Apellido: "Bond",
            Frases:   []string{"Shaken, not stirred", "¿Algún último deseo?", "Nunca digas nunca."},
        }

        bs, _ := json.Marshal(p1)
        fmt.Println(string(bs))

    }
    ```
