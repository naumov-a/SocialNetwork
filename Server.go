package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}

func register(w http.ResponseWriter, r *http.Request)  {
	session, _ := store.Get(r, "cookie-name")

	//FIXME
	if auth, ok := session.Values["authenticated"].(bool); ok || auth {
		http.Error(w, "Already registered", http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/registration.html"))
	tmpl.Execute(w,nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf("User %s", id)
	fmt.Fprint(w, response)
}


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user/{id:[0-9]+}", userHandler)
	http.Handle("/",router)

	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/register", register)


	http.ListenAndServe(":8080", nil)
}