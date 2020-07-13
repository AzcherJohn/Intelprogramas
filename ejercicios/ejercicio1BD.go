package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	err error
	db  *sql.DB
)

type Tablas struct {
	Nombre, Tipo string
}

func main() {
	conexionBD()
	tablasBase := mostrarTablasBD()

	crearTablaBD(tablasBase)
	for _, tabla := range mostrarTablasBD() {
		fmt.Println(tabla)
	}

	cerrarConexionBD()
}

func mostrarTablasBD() (tabla []Tablas) {
	mostrar, err := db.Query("SHOW FULL TABLES FROM CompuSi")
	comprobarError(err)

	for mostrar.Next() {
		var nombre, tipo string
		err = mostrar.Scan(&nombre, &tipo)
		comprobarError(err)

		tabla = append(tabla, Tablas{nombre, tipo})
	}

	return
}

func crearTablaBD(tabla []Tablas) {
	const nombreTabla = "catalogo_frutas"

	var tablaExist = false

	for i := 0; i < len(tabla); i++ {
		if tabla[i].Nombre == nombreTabla {
			tablaExist = true
		}
	}

	if tablaExist == false {
		crear, err := db.Exec("CREATE TABLE " + nombreTabla + "(id int primary key auto_increment, descrip varchar(50) not null, precio double not null);")
		comprobarError(err)
		_, err = crear.RowsAffected()
		comprobarError(err)
	}

}

func conexionBD() {
	db, err = sql.Open("mysql", string("root:Int0603239I7!@tcp(127.0.0.1:3306)/CompuSi"))
	comprobarError(err)

	err = db.Ping()
	comprobarError(err)
}

func comprobarError(err error) {
	if err != nil {
		panic(err)
	}
}

func cerrarConexionBD() {
	defer db.Close()
}
