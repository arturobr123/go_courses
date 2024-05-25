package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Imprime "SINGLETON" en la consola
	fmt.Println("SINGLETON")

	// Crea un grupo de espera para sincronizar múltiples goroutines
	var wg sync.WaitGroup

	// Inicia un bucle que se repite 10 veces
	for i := 0; i < 10; i++ {
		// Añade una goroutine al grupo de espera
		wg.Add(1)
		// Inicia una goroutine, que es una función ligera que se ejecuta concurrentemente
		go func() {
			// Asegura que se marque como completada al finalizar la goroutine
			defer wg.Done()
			// Llama a la función que obtiene o crea una instancia de la base de datos
			GetDatabaseInstance()
		}()
	}

	// Espera a que todas las goroutines en el grupo de espera terminen
	wg.Wait()
}

type Database struct {
}

var db *Database
var lock = sync.Mutex{}

func (db *Database) createSingleConnection() {
	fmt.Println("Creating singletion to database")
	time.Sleep(time.Second * 2)
	fmt.Println("Database connected")
}

func GetDatabaseInstance() *Database {
	// Lock the mutex to avoid race condition
	// tiene que esperar a que el thread termine para que los demas accedan
	lock.Lock()

	defer lock.Unlock() // Unlock the mutex after the function returns
	if db == nil {
		fmt.Println("Creating new database instance")
		db = &Database{}
		db.createSingleConnection()
	} else {
		fmt.Println("Database already created")
	}
	return db
}
