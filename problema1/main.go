package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Lanzar varias goroutines que imprimen mensajes y esperar a que todas terminen.
// TODO: Completa los pasos marcados con TODO para entender goroutines y WaitGroup.

func worker(id int, veces int, wg *sync.WaitGroup) {
	// TODO: asegurar que al finalizar la función se haga wg.Done() trabajo prueba 2
	defer wg.Done() // esta linea es para avisarle a waitGroup cuando el trabajador termina su trabajo

	for i := 1; i <= veces; i++ {
		fmt.Printf("[worker %d] hola %d\n", id, i)
		// TODO: dormir un poco para simular trabajo (p. ej. 100–300 ms)
		time.Sleep(400 * time.Millisecond)// este es un contador, da tiempo para simular que el trabajador esta realizando una tarea y que no se complete de forma isntantanea, si no que se espere un moneto para ver el avance"
	}
}

func main() {
	var wg sync.WaitGroup

	// TODO: cambiar estos parámetros y observar el intercalado de salidas
	numGoroutines := 5 // esta es para indicar la cantidad de trabajadores que van a estar realizando la tarea asignada
	veces := 8 // esta es la cantidad de veces que cada trabajador va a realizar la tarea que se asigno

	// TODO: lanzar varias goroutines, sumar al WG y esperar con wg.Wait()
	for id := 1; id <= numGoroutines; id++ {
		wg.Add(1) // esta funcion es para indicar que se va a agregar un trbajador
		go worker(id, veces, &wg) // esta funcion es para lanzar les  goroutines y que cada trabajador realice su tarea
	}

	// Esperar a que todas las goroutines terminen
	wg.Wait()

	fmt.Println("Listo: todas las goroutines terminaron.")
}
