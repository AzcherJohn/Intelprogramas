package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	err error
	db  *sql.DB
	tab = "catalogo_familia"
)

func main() {
	nuevaConexionDB()
	catalogoID("catalogo_familia", "Computadoras")

	//modificarDatosBD()
	//separador()
	resultadosQuery("catalogo_familia")
	separador()
	agregarDatosBD()
	_ = crearTablaBD()
	eliminarTablaBD("personas")
}

func separador() {
	fmt.Println("------------")
}

func nuevaConexionDB() {
	//usuario:contraseña@tcp(direccion y puerto)/nombre base
	db, err = sql.Open("mysql", string("root:*****@tcp(127.0.0.1:3306)/****"))
	revisarError(err)

	err = db.Ping()
	revisarError(err)
}

func catalogoID(tabla, where string) {
	var familia, serie string

	unicoDato := db.QueryRow("SELECT familia, serie  FROM "+tabla+" WHERE descripcion = ? ", where)

	err = unicoDato.Scan(&familia, &serie)
	revisarError(err)
	fmt.Println(familia, serie)
}

func resultadosQuery(tabla string) {

	query, _ := db.Query("SELECT * FROM " + tabla)

	for query.Next() {
		var id, familia, descripcion, fecha, serie string
		err = query.Scan(&id, &familia, &descripcion, &fecha, &serie)
		revisarError(err)
		fmt.Println(id, familia, descripcion, fecha, serie)
	}
}

func modificarDatosBD() {
	serie := "1"
	modi, err := db.Exec("Update catalogo_familia SET serie= ? ", serie)
	revisarError(err)

	_, err = modi.RowsAffected()
	revisarError(err)
}

func agregarDatosBD() {
	familia, descri, fecha, serie := "lamled", "lamparas led", "2020-07-09", "0"
	queryTexto := "Insert into catalogo_familia (familia, descripcion, ultima_actualizacion, serie)"
	agregar, err := db.Exec(queryTexto+" values (?, ?, ?, ?)", familia, descri, fecha, serie)
	revisarError(err)
	status, err := agregar.LastInsertId()
	revisarError(err)

	resultadosQuery(tab)
	separador()
	eliminarDatosBD(status)
}

func eliminarDatosBD(lastID int64) {
	eliminar, err := db.Exec("delete from catalogo_familia where catalogo_familia_id = ?", lastID)
	revisarError(err)
	status, err := eliminar.RowsAffected()
	revisarError(err)
	fmt.Println(status)

	separador()
	resultadosQuery(tab)
}

func crearTablaBD() string {
	nombreTabla := "personas2"
	crearT, err := db.Exec("create table " + nombreTabla + " (id int, nombre varchar(15));")
	revisarError(err)
	_, err = crearT.RowsAffected()
	revisarError(err)
	return nombreTabla
}

func eliminarTablaBD(nombreTabla string) {
	eliminarT, err := db.Exec("drop table " + nombreTabla)
	revisarError(err)
	_, err = eliminarT.RowsAffected()
	revisarError(err)
}

func revisarError(err error) {
	if err != nil {
		panic(err)
	}
}
