**Ejercicio Práctico #2**

- Crea un strutc persona
- Crea una función llamada “cambiame” con un *persona como parámetro
    - Cambia el valor almacenado en la dirección de memoria del *persona
    - Impotante
        - Para desreferenciar un struct, usa (*valor).campo 
            - p1.nombre
            - (*p1).nombre
        - “Como una excepción, si el tipo de x es un tipo puntero con nombre y (*x).c es una expresión válida de selección denotando un campo (pero no un método), x.c es una forma abreviada de (*x).c”.
            - https://golang.org/ref/spec#Selectors 
- Crea un valor de tipo persona
    - Imprime el valor
- En func main
    - llama “cambiame”
- En func main
    - Imprime el valor
