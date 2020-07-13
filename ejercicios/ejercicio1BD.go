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

type Frutas struct {
	Descripcion, Precio string
}

type CatalogoBD struct {
	Id, Descripcion, Precio string
}

const tablaFruta = "catalogo_frutas"

func main() {
	var fruta []Frutas
	fruta = append(fruta, Frutas{"manzana", "15"})
	fruta = append(fruta, Frutas{"pera", "22"})
	fruta = append(fruta, Frutas{"sandia", "25"})

	conexionBD()
	tablasBase := mostrarTablasBD()

	crearTablaBD(tablasBase)
	for _, tabla := range mostrarTablasBD() {
		fmt.Println(tabla)
	}
	for _, datos := range fruta {
		fmt.Println(insertarDatosCatalgoBD(tablaFruta, "descrip, precio", datos.Descripcion, datos.Precio))
	}
	fmt.Println(mostrarTodosDatosBD(tablaFruta))
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

func insertarDatosCatalgoBD(nombreTabla, columnas, valor1, valor2 string) int64 {
	insertar, err := db.Exec("INSERT INTO "+nombreTabla+" ("+columnas+") VALUES (?, ?)", valor1, valor2)
	comprobarError(err)
	status, err := insertar.LastInsertId()
	return status
}

func mostrarTodosDatosBD(nombreTabla string) (fruta []CatalogoBD) {
	mostrar, err := db.Query("SELECT * FROM " + nombreTabla)
	comprobarError(err)
	for mostrar.Next() {
		var id, desc, precio string
		err = mostrar.Scan(&id, &desc, &precio)
		comprobarError(err)

		fruta = append(fruta, CatalogoBD{id, desc, precio})
	}
	return
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
