**Ejercicio Práctico #1**

- Comenzando con este código, marshal el slice []usuario a JSON. Hay una pequeña curva en este ejercicio - recuerda preguntarte qué necesitas para EXPORTAR un valor de un paquete.

    package main

    type usuario struct {
        nombre string
        edad   int
    }

    func main() {
        u1 := usuario{
            nombre: "Eduar",
            edad:   32,
        }

        u2 := usuario{
            nombre: "Juan",
            edad:   27,
        }

        u3 := usuario{
            nombre: "Che",
            edad:   54,
        }

        usuarios := []usuario{u1, u2, u3}

        //tu código va aquí
    }