package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

type User struct {
	Codigo int
	User   string
	Name   string
	Ape    string
	Edad   string
}

var users map[int]*User
var counter int
var templates = template.Must(template.ParseGlob("templates/*"))

func MainHandler(res http.ResponseWriter, req *http.Request) {

	templates.ExecuteTemplate(res, "base", nil)
}

func Form(res http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(res, "form", users)
}

func MethodsUser(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		templates.ExecuteTemplate(res, "users", users)
	}

	if req.Method == "POST" {
		codigo := counter
		usuario := req.FormValue("usuario")
		nombre := req.FormValue("nombre")
		apellido := req.FormValue("ape")
		edad := req.FormValue("edad")
		user := &User{codigo, usuario, nombre, apellido, edad}
		users[counter] = user
		counter++
		err := templates.ExecuteTemplate(res, "form", users)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func UserActions(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	if req.Method == "GET" {
		templates.ExecuteTemplate(res, "search_user", users[id])
	}

}

func NotFound(res http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(res, "404", nil)
}

func main() {
	// Inicializamos en 0 el MAP users
	users = make(map[int]*User, 0)
	// Declaramos EndPoints
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// http.HandleFunc("/form", Form)
	// http.HandleFunc("/user", MethodsUser)
	// http.HandleFunc("/", MainHandler)

	r := mux.NewRouter()
	// r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.HandleFunc("/", MainHandler)
	r.HandleFunc("/form", Form)
	r.HandleFunc("/user", MethodsUser)
	r.HandleFunc("/user/{id}", UserActions)
	r.NotFoundHandler = http.HandlerFunc(NotFound)
	//El servidor estara escuchando en el port 5000 puedes hacer que ejecute un metodo en vez de el nil
	http.ListenAndServe(":5000", r)

}
