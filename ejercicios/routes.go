package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var paginas = populateStaticPages()

func main() {
	//GorillaRoute es el nombre de la variable que su función es la de permitir navegar mediantes URLs en el servidor
	//StaticSlash Nos permite tener ramificaciones en las URLs
	gorillaRoute := mux.NewRouter().StrictSlash(true)
	//La página por defecto del serevidor se encuentra en el URL '/' y lo que se mostrará esta en la función ContenidoServer
	gorillaRoute.HandleFunc("/", contenidoServer)
	//Las laves '{}' indican que ahí va una variable y que no se trata de una palabra literal
	gorillaRoute.HandleFunc("/{aliasPagina}", contenidoServer)
	//Aquí se busca en las carpetas los recursos que el servidor necesita, en caso que los necesite, para encontrar los archivos se usa la func RecursosServidor

	//Se menciona la primer ruta que se utiliza en el servidor
	http.Handle("/", gorillaRoute)

	/*
		router := NewRouter()

		//El dominio o lugar en donde se encuentra el servidor
		server := http.ListenAndServe(":9090", router)
		log.Fatal(server)
	*/

	http.ListenAndServe(":9090", nil)
}

func populateStaticPages() *template.Template {
	//Aquí se buscan todos los templates html principales para el servidor.

	TemplatePaths := new([]string)

	basePath := "html"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()
	templatePathsRaw, _ := templateFolder.Readdir(-1)

	for _, pathInfo := range templatePathsRaw {
		log.Println(pathInfo.Name())
		*TemplatePaths = append(*TemplatePaths, basePath+"/"+pathInfo.Name())
	}

	/*funciones := template.FuncMap{
		"MenuDesplegable": MenuDesplegable,
		"CambioLista":     CambioLista,
	}*/

	//resultado, _ := template.New("InicioSesion.html").Funcs(funciones).ParseFiles(*TemplatePaths...)
	resultado, _ := template.New("InicioSesion.html").ParseFiles(*TemplatePaths...)
	return resultado
}

func contenidoServer(w http.ResponseWriter, r *http.Request) {
	var pagina *template.Template
	//Permite leer el URL que se ha introducido en el navegador
	parametroURL := mux.Vars(r)
	//Aqui se guarda la URL introducida
	aliasPagina := parametroURL["aliasPagina"]

	pagina = paginas.Lookup(aliasPagina + ".html")
	if pagina == nil {
		pagina = paginas.Lookup("NoEncontrado.html")
	}
	switch aliasPagina {

	//case "InicioSesion":
	//pagina.Execute(w, ListaPrecios(w, r /*, SucursalId*/))
	default:
		pagina.Execute(w, nil)
	}
}
