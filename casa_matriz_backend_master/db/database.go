package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

// Conexión de la base de datos
var databaseInstance *sql.DB
var lock = &sync.Mutex{}

// Singleton para la base de datos & abrir conexión
func GetDatabase() *sql.DB {
	if databaseInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if databaseInstance == nil {
			databaseInstance = dbConnect()
		} else {
			return databaseInstance
		}
	} else {
		return databaseInstance
	}
	return databaseInstance
}

// Devolver instancia a cualquier parte del backend
func dbConnect() *sql.DB {
	var db *sql.DB
	var error error
	db, error = sql.Open("mysql", "kabra:Belleza!09@tcp(25.4.182.181:3306)/casamatriz")
	if error != nil {
		log.Fatal(error)

	}
	return db
}

func TestDB() {
	// Hacer una consulta a la base de datos para obtener la versión
	var version string
	var db = GetDatabase()
	err := db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("La versión de MariaDB es:", version)
}
